//jsl created at 2019/12/17

package gutil

import (
	"strconv"
	"testing"
)

func TestListOf(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	if list.shiftValue.Kind().String() != "slice" {
		t.FailNow()
	}
}

func TestListOf_interface(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	var lf ListFunc = list
	t.Logf("%v", lf.ToList())
}

func TestListOf_indefinite(t *testing.T) {
	list := ListOf("a", "b", "c")

	if list.shiftValue.Kind().String() != "slice" {
		t.FailNow()
	}

	if list.shiftValue.Index(0).Kind().String() != "string" {
		t.FailNow()
	}
}

func TestList_ToListArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	list2 := list.ToList().([]int)
	t.Log(list2[0])
}

func TestList_AnyArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Any(func (i interface{}) bool {
		return i.(int) == 3
	})
	t.Log(result)
}

func TestList_AnySlice(t *testing.T) {
	v := []int{1,2,3}
	list := ListOf(v)

	result := list.Any(func (i interface{}) bool {
		return i.(int) == 3
	})
	t.Log(result)
}

func TestList_AnyIndefinite(t *testing.T) {
	list := ListOf("a", "b", "c")
	result := list.Any(func (i interface{}) bool {
		return i.(string) == "b"
	})
	t.Log(result)
}

func TestList_ForEachArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	list.ForEach(func (i interface{}) {
		println(i.(int))
	})
}

func TestList_ForEachSlice(t *testing.T) {
	v := []int{1,2,3}
	list := ListOf(v)

	list.ForEach(func (i interface{}) {
		println(i.(int))
	})
}

func TestList_ForEachIndefinite(t *testing.T) {
	list := ListOf("a", "b", "c")
	list.ForEach(func (i interface{}) {
		println(i.(string))
	})
}

func TestList_ForEachIndexArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	list.ForEachIndex(func (index int, i interface{}) {
		println(index, i.(int))
	})
}

func TestList_ForEachIndexSlice(t *testing.T) {
	v := []int{1,2,3}
	list := ListOf(v)

	list.ForEachIndex(func (index int, i interface{}) {
		println(index, i.(int))
	})
}

func TestList_ForEachIndexIndefinite(t *testing.T) {
	list := ListOf("a", "b", "c")
	list.ForEachIndex(func (index int, i interface{}) {
		println(index, i.(string))
	})
}

func TestList_WhereArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Where(func (i interface{}) bool {
		return i.(int) != 3
	}).ToList().([]int)

	if len(result) != 2 || result[1] != 2 {
		t.FailNow()
	}
}

func TestList_RemoveWhereArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.RemoveWhere(func (i interface{}) bool {
		return i.(int) == 3
	}).ToList().([]int)

	if len(result) != 2 || result[1] != 2 {
		t.FailNow()
	}
}

func TestList_Remove(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Remove(2).ToList().([]int)

	if result[0] != 1 && result[1] != 3 {
		t.FailNow()
	}
}

func TestList_FirstArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.First(func (i interface{}) bool {
		return i.(int)%2 == 1
	}).(int)

	if result != 1 {
		t.FailNow()
	}
}

func TestList_FirstArrayNotFound(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	defer func() {
		err := recover()
		if err == nil {
			t.FailNow()
		}
		t.Log(err)
	}()

	result := list.First(func (i interface{}) bool {
		return i.(int) == 4
	}).(int)

	println(result)
}

func TestList_FirstOrNullArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.FirstOrNull(func (i interface{}) bool {
		return i.(int)%2 == 1
	}).(int)

	if result != 1 {
		t.FailNow()
	}
}

func TestList_FirstOrNullArrayNotFound(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.FirstOrNull(func (i interface{}) bool {
		return i.(int) == 4
	})

	if result != nil {
		t.FailNow()
	}
}

func TestList_LastArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Last(func (i interface{}) bool {
		return i.(int)%2 == 1
	}).(int)

	if result != 3 {
		t.FailNow()
	}
}

func TestList_LastArrayNotFound(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	defer func() {
		err := recover()
		if err == nil {
			t.FailNow()
		}
		t.Log(err)
	}()

	result := list.Last(func (i interface{}) bool {
		return i.(int) == 4
	}).(int)

	println(result)
}

func TestList_LastOrNullArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.LastOrNull(func (i interface{}) bool {
		return i.(int)%2 == 1
	}).(int)

	if result != 3 {
		t.FailNow()
	}
}

func TestList_LastOrNullArrayNotFound(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.LastOrNull(func (i interface{}) bool {
		return i.(int) == 4
	})

	if result != nil {
		t.FailNow()
	}
}

func TestList_MapArray(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Map(func (i interface{}) interface{} {
		return strconv.Itoa(i.(int) * 2) + " item"
	}).ToList().([]string)

	if result[0] != "2 item" {
		t.Failed()
	}
}

func TestList_FlatMapArray(t *testing.T) {
	type T struct {
		List []string
	}

	v := []T{
		{List: []string{"a", "b"}},
		{List: []string{"c", "d"}},
		{List: []string{"e", "f"}},
	}
	list := ListOf(v)

	result := list.FlatMap(func (i interface{}) interface{} {
		return i.(T).List
	}).ToList().([]string)

	t.Logf("%v", result)
	if result[3] != "d" {
		t.Failed()
	}
}

func TestList_Add(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.Add(4).ToList().([]int)

	if result[3] != 4 {
		t.FailNow()
	}
}

func TestList_AddAll(t *testing.T) {
	v := [3]int{1,2,3}
	list := ListOf(v)

	result := list.AddAll([]int{4, 5, 6}).ToList().([]int)

	if result[5] != 6 {
		t.FailNow()
	}
}

func TestList_Sort(t *testing.T) {
	v := []int{4, 3, 5, 8, 1, 6}
	list := ListOf(v)

	result := list.Sort(func (i1 interface{}, i2 interface{}) int {
		return i1.(int) - i2.(int)
	}).ToList().([]int)

	t.Logf("%v", result)
	if result[0] != 1 || result[5] != 8 {
		t.FailNow()
	}
}

func TestList_ReverseOdd(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	list := ListOf(v)

	result := list.Reverse().ToList().([]int)

	t.Logf("%v", result)
	if result[0] != 5 || result[4] != 1 {
		t.FailNow()
	}
}

func TestList_ReverseEven(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Reverse().ToList().([]int)

	t.Logf("%v", result)
	if result[0] != 6 || result[5] != 1 {
		t.FailNow()
	}
}

func TestList_ReversePointer(t *testing.T) {
	a,b,c,d := "a", "b", "c", "d"
	list := ListOf(&a, &b, &c, &d)

	result := list.Reverse().ToList().([]*string)

	t.Logf("%v", result)
	if *result[0] != "d" || *result[3] != "a" {
		t.FailNow()
	}
}

func TestList_Reduce(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Reduce(func(i1 interface{}, i2 interface{}) interface{} {
		return i1.(int) + i2.(int)
	}).(int)

	if result != 21 {
		t.FailNow()
	}
}

func TestList_Contains(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Contains(5)

	if !result {
		t.FailNow()
	}
}

func TestList_ContainsNotFound(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Contains(7)

	if result {
		t.FailNow()
	}
}

func TestList_ContainsPointer(t *testing.T) {
	a,b,c,d := "a", "b", "c", "d"
	list := ListOf(&a, &b, &c, &d)

	result := list.Contains(&b)

	if !result {
		t.FailNow()
	}
}

func TestList_ContainsStruct(t *testing.T) {
	type L struct {
		M string
	}
	v := []L {
		{M: "A"},
		{M: "B"},
		{M: "C"},
	}
	list := ListOf(v)

	result := list.Contains(L{M:"B"})

	if !result {
		t.FailNow()
	}
}

func TestList_ContainsStructNotFound(t *testing.T) {
	type L struct {
		M string
	}
	v := []L {
		{M: "A"},
		{M: "B"},
		{M: "C"},
	}
	list := ListOf(v)

	result := list.Contains(L{M:"D"})

	if result {
		t.FailNow()
	}
}

func TestList_Index(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Index(2)

	if result != 1 {
		t.FailNow()
	}
}

func TestList_IndexNotFound(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.Index(0)

	if result != -1 {
		t.FailNow()
	}
}

func TestList_IndexOf(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	result := list.IndexOf(2)

	if result != 3 {
		t.FailNow()
	}
}

func TestList_IndexOfOutOfRange(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	list := ListOf(v)

	defer func() {
		err := recover()
		if err == nil {
			t.FailNow()
		}
		t.Log(err)
	}()

	_ = list.IndexOf(6)
}