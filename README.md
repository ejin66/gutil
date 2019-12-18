## gutil
支持对array/slice/不定长入参的糖语法API

## example
#### ForEach
```go
s := []int {1, 2, 3, 4}
gutil.ListOf(s).ForEach(func(i interface{}) {
  fmt.Print(i.(int) * 2, " ")
})
//print: 2 4 6 8 
```
#### Reverse
```go
r := gutil.ListOf("a", "b", "c").Reverse().ToList().([]string)
fmt.Printf("%v\n", r)
//print: [c b a]
```
#### 流式api
```go
s := []int {1, 2, 3, 4}
ss := gutil.ListOf(s).Where(func(i interface{}) bool {
  return i.(int) != 2
}).Sort(func(i interface{}, i2 interface{}) int {
  return i2.(int) - i.(int)
}).ToList().([]int)
fmt.Printf("%v", ss)
//print: [4 3 1]
```

## doc
- Any
- ForEach
- ForEachIndex
- Where
- RemoveWhere
- Remove
- First
- FirstOrNull
- Last
- LastOrNull
- Map
- FlatMap
- Add
- AddAll
- Sort
- Reverse
- Reduce
- Contains
- Index
- IndexOf
- ToList
