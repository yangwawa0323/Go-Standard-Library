package reflect_test

import (
	"reflect"
	"testing"
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
