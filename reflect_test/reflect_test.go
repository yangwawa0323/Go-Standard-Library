package reflect_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func Test_Copy(t *testing.T) {

	// the length of destination slice is 2. After copied the source element,
	// the length still be 2. only 2 elements been copied from source slice.
	// destination := reflect.ValueOf([]string{"A", "B"})
	// source := reflect.ValueOf([]string{"D", "E", "F"})

	// this situation is that source provides less elements that destination.
	// the remaining element in destination will be untouched.
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E"})

	n := reflect.Copy(destination, source)
	t.Log(n)
	t.Log("destination: ", destination)
	t.Log("source:", source)
}

type mobile struct {
	Price float64
	Color string
}

func Test_DeepEqual(t *testing.T) {
	// DeepEqual is used to check two slices are equal or not
	s1 := []string{"A", "B", "C", "D", "F"}
	s2 := []string{"D", "E", "F"}
	result := reflect.DeepEqual(s1, s2)
	t.Log("s1 is deep equal s2: ", result)

	// DeepEqual is used to check two arrays are equal or not
	n1 := [5]int{1, 2, 3, 4, 5}
	n2 := [5]int{1, 2, 3, 4, 5}

	result = reflect.DeepEqual(n1, n2)
	t.Log("n1 is deep equal n2: ", result)

	// DeepEqual is used to check two structures are equal or not
	m1 := mobile{500.50, "red"}
	m2 := mobile{400.50, "black"}
	result = reflect.DeepEqual(m1, m2)
	t.Log("m1 is deep equal m2: ", result)
}

func Test_Swapper(t *testing.T) {
	theList := []int{1, 2, 3, 4, 5}
	swap := reflect.Swapper(theList) // return a swap func
	t.Log("original Slice: ", theList)

	// Swapper() function is used to swaps the element of slice
	swap(1, 3)
	t.Log("After swap Slice: ", theList)

	// Reversing a slice using Swapper() function
	for i := 0; i < len(theList)/2; i++ {
		// leftSide => 0 ; rightSide => 5 - 1 - 0 = 4
		// leftSide => 1 ; rightSide => 5 - 1 - 1 = 3
		// leftSide => 2 ; rightSide => 5 - 1 - 2 = 2
		leftSide := i
		rightSide := len(theList) - 1 - i
		swap(leftSide, rightSide)
	}

	t.Log("After Reverse Slice : ", theList)

}

func Test_TypeOf(t *testing.T) {
	v1 := []int{1, 2, 3, 4, 5}
	t.Log(reflect.TypeOf(v1))

	v2 := "Hello World"
	t.Log(reflect.TypeOf(v2))

	v3 := 1000
	t.Log(reflect.TypeOf(v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	t.Log(reflect.TypeOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	t.Log(reflect.TypeOf(v5))

	v6 := true
	t.Log(reflect.TypeOf(v6))

}

func Test_ValueOf(t *testing.T) {
	v1 := []int{1, 2, 3, 4, 5}
	t.Log(reflect.ValueOf(v1))

	v2 := "Hello World"
	t.Log(reflect.ValueOf(v2))

	v3 := 1000
	t.Log(reflect.ValueOf(v3))
	t.Log(reflect.ValueOf(&v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	t.Log(reflect.ValueOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	t.Log(reflect.ValueOf(v5))

	v6 := true
	t.Log(reflect.ValueOf(v6))
}

type T struct {
	A int
	B string
	C float64
	D bool
}

func Test_Type_NumField(t *testing.T) {
	td := T{10, "ABCDEF", 15.30, true}
	typeT := reflect.TypeOf(td)
	t.Log("Total fields :", typeT.NumField())
}

func Test_Field(t *testing.T) {
	td := T{10, "ABCDEF", 15.30, true}
	typeT := reflect.TypeOf(td)
	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		t.Logf("%d: filed is: %s, its value: %s\n", i, field.Name, field.Type)
	}
}

type First struct {
	A int
	B string `json:"FieldBInFirst"`
	C float64
}

type Second struct {
	First
	D bool
}

func Test_FieldByIndex(t *testing.T) {
	s := Second{First: First{10, "ABCDEF", 15.30}, D: true}
	ts := reflect.TypeOf(s)

	// Output : reflect.StructField{Name:"B", PkgPath:"", Type:(*reflect.rtype)(0x885ba0), Tag:"json:\"FieldBInFirst\"", Offset:0x8, Index:[]int{1}, Anonymous:false}
	// the reflect.StructField has a Tag attribute.
	t.Logf("index 0: %#v\ntype is %s \n", ts.FieldByIndex([]int{0}), ts.Field(0).Type)
	t.Logf("index 0, 0 %#v\n: ", ts.FieldByIndex([]int{0, 0}))

	// To get the filed type of embedded struct. you can get each level field type and chainning its inside level
	// For example: `ts.Field(0).Type` returns `reflect_test.First` struct
	// To get the type of 2nd field in `First` , we use chainning expression
	// `ts.Field(0).Type.Field(1).Type` that will be returned result `string`
	t.Logf("index 0, 1: %#v\ntype is %s\n",
		ts.FieldByIndex([]int{0, 1}),
		ts.Field(0).Type.Field(1).Type)
	t.Logf("index 0, 2: %#v\n", ts.FieldByIndex([]int{0, 2}))
	t.Logf("index 1: %#v\n", ts.FieldByIndex([]int{1}))

}

func Test_FieldByName(t *testing.T) {
	s := T{10, "ABCDEF", 15.4, true}
	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("A"))
	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("B"))
	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("C"))

	t.Log(" ============= Now set the field value ===============")

	reflect.ValueOf(&s).Elem().FieldByName("A").SetInt(
		2 * reflect.ValueOf(&s).Elem().FieldByName("A").Int())
	reflect.ValueOf(&s).Elem().FieldByName("B").SetString(
		fmt.Sprintf("Hello world: %s",
			reflect.ValueOf(&s).Elem().FieldByName("B").String()))

	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("A"))
	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("B"))
	t.Log("Field A: ", reflect.ValueOf(&s).Elem().FieldByName("C"))

}

func Test_MakeSlice(t *testing.T) {
	var str []string
	var strType reflect.Value = reflect.ValueOf(str)
	newSlice := reflect.MakeSlice(reflect.Indirect(strType).Type(), 10, 15)
	newSlice.Index(0).SetString("Hello")
	// call of reflect.Value.Field on slice Value
	// newSlice.Field(1).SetString("World")
	newSlice.Index(1).SetString("World")
	// cap := newSlice.Cap()

	// Can not access the range out of slice length.
	var i = 2

	defer func() {

		// catching the panic
		if recover() != nil {
			t.Log("out of range")
			// Save the orginal elements is the newSlice
			oldSlice := newSlice

			// re-defined the slice capacity and length
			newSlice = reflect.MakeSlice(reflect.Indirect(strType).Type(), 20, 25)
			// copy the orginal elements to newSlice
			reflect.Copy(newSlice, oldSlice)
		}

		// Access the newSlice elements, first of all get the interface, and use assert to the corresponding type
		s := newSlice.Interface().([]string)
		s[11] = "Better way"
		t.Logf("newSlice is: %q, type: %#v", s, newSlice.Type().String())
		t.Log("Kind: ", newSlice.Kind())
		t.Log("Length: ", newSlice.Len())
		t.Log("Capacity: ", newSlice.Cap())
	}()

	// panic out of slice range
	for ; i < newSlice.Len()+1; i++ {
		newSlice.Index(i).SetString(fmt.Sprintf("%d", i))
	}

}

func Test_MakeMap(t *testing.T) {
	var strmap map[string]string

	var strmapType = reflect.ValueOf(strmap)
	newMap := reflect.MakeMap(reflect.Indirect(strmapType).Type())
	_newmap := newMap.Interface().(map[string]string)
	_newmap["host"] = "localhost"
	t.Logf("newMap: %#v", newMap)
}

func Test_MakeChan(t *testing.T) {
	var str []chan string = make([]chan string, 10)
	var strType = reflect.ValueOf(&str)

	strChanSliceType := reflect.MakeSlice(reflect.Indirect(strType).Type(), 0, 0)

	newStrChanSlice := strChanSliceType.Interface().([]chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	var ready = make(chan bool, 1)

	go func() {
		if <-ready {
			for i := 0; i < 10; i++ {
				t.Log("Read message from channel :", <-newStrChanSlice[i])
			}
		}
		wg.Done()
	}()

	go func() {

		time.Sleep(3 * time.Second)
		var strChan = make(chan string, 512)
		for i := 0; i < 10; i++ {
			strChan <- fmt.Sprintf("Channel %d", i)
			newStrChanSlice = append(newStrChanSlice, strChan)
		}
		ready <- true
		wg.Done()

	}()

	wg.Wait()
}

type Sum func(int64, int64) int64

func Test_MakeFunc(t *testing.T) {
	funcType := reflect.TypeOf(Sum(nil))
	mul := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		return []reflect.Value{reflect.ValueOf(a + b)}
	})

	fn, ok := mul.Interface().(Sum)
	if !ok {
		return
	}
	t.Log(fn(5, 6))
}
