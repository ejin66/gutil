//jsl created at 2019/12/17

package gutil

import (
	"reflect"
)

type ValueInBoolOut = func(interface{}) bool
type ValueIn = func(interface{})
type IndexValueIn = func(int, interface{})
type ValueInValue2Out = func(interface{}) interface{}
type Value12InValue3Out = func(interface{}, interface{}) interface{}
type ValueSort = func(interface{}, interface{}) int

type ListFunc interface {
	Any(f ValueInBoolOut) bool
	ForEach(f ValueIn)
	ForEachIndex(f IndexValueIn)
	Where(f ValueInBoolOut) *List
	RemoveWhere(f ValueInBoolOut) *List
	Remove(obj interface{}) *List
	First(f ValueInBoolOut) interface{}
	FirstOrNull(f ValueInBoolOut) interface{}
	Last(f ValueInBoolOut) interface{}
	LastOrNull(f ValueInBoolOut) interface{}
	Map(f ValueInValue2Out) *List
	FlatMap(f ValueInValue2Out) *List
	Add(obj interface{}) *List
	AddAll(obj interface{}) *List
	Sort(f ValueSort) *List
	Reverse() *List
	Reduce(f Value12InValue3Out) interface{}
	Contains(obj interface{}) bool
	Index(obj interface{}) int
	IndexOf(index int) interface{}
	ToList() interface{}
}

type List struct {
	//array or slice
	originValue interface{}
	//slice
	shiftValue reflect.Value
}

func ListOf(v ...interface{}) *List {
	if len(v)== 0 {
		panic("arguments can't be empty")
	}

	list := List{
		originValue: v,
	}

	if len(v) > 1 {
		refV := reflect.ValueOf(v[0])
		refKind := refV.Kind().String()
		length := len(v)
		isSameKind := true
		for i := 1; i < length; i++ {
			if refKind != reflect.ValueOf(v[i]).Kind().String() {
				isSameKind = false
				break
			}
		}
		if isSameKind {
			sType := reflect.SliceOf(refV.Type())
			sv := reflect.MakeSlice(sType, 0, length)
			for i := 0; i < length; i++ {
				sv = reflect.Append(sv, reflect.ValueOf(v[i]))
			}
			list.shiftValue = sv
		} else {
			list.shiftValue = reflect.ValueOf(v)
		}
	} else {
		rv := reflect.ValueOf(v[0])
		rvk := rv.Kind()

		if rvk == reflect.Ptr {
			rv = rv.Elem()
			rvk = rv.Kind()
		}

		if rvk == reflect.Slice {
			if rv.Len() == 0 {
				panic("the only slice argument is empty")
			}
			list.shiftValue = rv
		} else if rvk == reflect.Array {
			length := rv.Len()
			if length == 0 {
				panic("the only array argument is empty")
			}
			t := rv.Index(0).Type()
			sliceV := reflect.MakeSlice(reflect.SliceOf(t), 0, length)
			for i := 0; i < length; i++ {
				sliceV = reflect.Append(sliceV, rv.Index(i))
			}
			list.shiftValue = sliceV
		} else {
			panic("the only argument is neither array nor slice")
		}
	}
	return &list
}

func (this *List) Any(f ValueInBoolOut) bool {
	length := this.shiftValue.Len()
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		if f(item) {
			return true
		}
	}
	return false
}

func (this *List) ForEach(f ValueIn) {
	length := this.shiftValue.Len()
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		f(item)
	}
}

func (this *List) ForEachIndex(f IndexValueIn) {
	length := this.shiftValue.Len()
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		f(i, item)
	}
}

func (this *List) ToList() interface{} {
	length := this.shiftValue.Len()
	if length == 0 {
		return make([]interface{}, 0)
	}

	return this.shiftValue.Interface()
}

func (this *List) Where(f ValueInBoolOut) *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}
	nSliceV := reflect.MakeSlice(reflect.SliceOf(this.shiftValue.Index(0).Type()), 0, length)
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		if f(item) {
			nSliceV = reflect.Append(nSliceV, this.shiftValue.Index(i))
		}
	}
	this.shiftValue = nSliceV
	return this
}

func (this *List) RemoveWhere(f ValueInBoolOut) *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}
	nSliceV := reflect.MakeSlice(reflect.SliceOf(this.shiftValue.Index(0).Type()), 0, length)
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		if !f(item) {
			nSliceV = reflect.Append(nSliceV, this.shiftValue.Index(i))
		}
	}
	this.shiftValue = nSliceV
	return this
}

func (this *List) Remove(obj interface{}) *List {
	index := this.Index(obj)
	if index >= 0 {
		this.shiftValue = reflect.AppendSlice(this.shiftValue.Slice(0, index), this.shiftValue.Slice(index + 1, this.shiftValue.Len()))
	}
	return this
}

func (this *List) First(f ValueInBoolOut) interface{} {
	result := this.FirstOrNull(f)
	if result == nil {
		panic("none can be found that matched the condition")
	}
	return result
}

func (this *List) FirstOrNull(f ValueInBoolOut) interface{} {
	length := this.shiftValue.Len()
	for i := 0; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		if f(item) {
			return item
		}
	}
	return nil
}

func (this *List) Last(f ValueInBoolOut) interface{} {
	result := this.LastOrNull(f)
	if result == nil {
		panic("none can be found that matched the condition")
	}
	return result
}

func (this *List) LastOrNull(f ValueInBoolOut) interface{} {
	length := this.shiftValue.Len()
	for i := length - 1; i >= 0; i-- {
		item := toObj(this.shiftValue.Index(i))
		if f(item) {
			return item
		}
	}
	return nil
}

func (this *List) Map(f ValueInValue2Out) *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}

	firstV := reflect.ValueOf(f(toObj(this.shiftValue.Index(0))))
	nSliceV := reflect.MakeSlice(reflect.SliceOf(firstV.Type()), 0, length)
	nSliceV = reflect.Append(nSliceV, firstV)
	for i := 1; i < length; i++ {
		item := toObj(this.shiftValue.Index(i))
		nSliceV = reflect.Append(nSliceV, reflect.ValueOf(f(item)))
	}
	this.shiftValue = nSliceV
	return this
}

func (this *List) FlatMap(f ValueInValue2Out) *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}

	nSliceV := reflect.ValueOf(f(toObj(this.shiftValue.Index(0))))
	if nSliceV.Kind() != reflect.Slice {
		panic("the shift value must be slice")
	}
	for i := 1; i < length; i++ {
		itemSliceV := reflect.ValueOf(f(toObj(this.shiftValue.Index(i))))
		nSliceV = reflect.AppendSlice(nSliceV, itemSliceV)
	}
	this.shiftValue = nSliceV
	return this
}

func (this *List) Add(obj interface{}) *List {
	this.shiftValue = reflect.Append(this.shiftValue, reflect.ValueOf(obj))
	return this
}

func (this *List) AddAll(obj interface{}) *List {
	this.shiftValue = reflect.AppendSlice(this.shiftValue, reflect.ValueOf(obj))
	return this
}

func (this *List) Sort(f ValueSort) *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}
	nSliceV := reflect.MakeSlice(reflect.SliceOf(this.shiftValue.Index(0).Type()), 0, length)

	for length > 0 {
		min := this.shiftValue.Index(0)
		minIndex := 0
		for i := 1; i < length; i++ {
			r := f(toObj(min), toObj(this.shiftValue.Index(i)))
			if r > 0 {
				min = this.shiftValue.Index(i)
				minIndex = i
			}
		}
		nSliceV = reflect.Append(nSliceV, min)
		this.shiftValue = reflect.AppendSlice(this.shiftValue.Slice(0, minIndex), this.shiftValue.Slice(minIndex + 1, this.shiftValue.Len()))
		length -= 1
	}

	this.shiftValue = nSliceV
	return this
}

func (this *List) Reverse() *List {
	length := this.shiftValue.Len()
	if length == 0 {
		return this
	}

	tmpV := reflect.New(this.shiftValue.Index(0).Type()).Elem()
	for i := 0; i < length / 2; i++ {
		tmpV.Set(this.shiftValue.Index(i))
		this.shiftValue.Index(i).Set(this.shiftValue.Index(length - 1 -i))
		this.shiftValue.Index(length - 1 -i).Set(tmpV)
	}
	return this
}

func (this *List) Reduce(f Value12InValue3Out) interface{} {
	length := this.shiftValue.Len()
	if length == 0 {
		return nil
	}

	result := this.shiftValue.Index(0).Interface()
	for i := 1; i < length; i++ {
		result = f(result, toObj(this.shiftValue.Index(i)))
	}

	return result
}

func (this *List) Index(obj interface{}) int {
	length := this.shiftValue.Len()
	if length == 0 {
		return -1
	}

	for i := 0; i < length; i++ {
		itemV := this.shiftValue.Index(i)
		if obj == toObj(itemV) {
			return i
		}
	}

	return -1
}

func (this *List) Contains(obj interface{}) bool {
	return this.Index(obj) >= 0
}

func (this *List) IndexOf(index int) interface{} {
	length := this.shiftValue.Len()

	if index >= length {
		panic("index out of range")
	}

	if length == 0 {
		return nil
	}

	return this.shiftValue.Index(index).Interface()
}

func toObj(v reflect.Value) interface{} {
	return v.Interface()
}