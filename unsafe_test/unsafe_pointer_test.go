package unsafe_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_Pointer(t *testing.T) {
	// An array of contiguous uint32 values store in memory.
	arr := []uint32{1, 2, 3}

	// The number of bytes each uint32 occupies: 4
	const size = unsafe.Sizeof(uint32(0))

	// Take the initial memory address of the array and begin iteration.
	p := uintptr(unsafe.Pointer(&arr[0]))
	for i := 0; i < len(arr); i++ {
		// Print the integer that resides at the current address and then
		// increment the pointer to the next value in the array.
		fmt.Printf("access the pointer %#v value %d\n", p, (*(*uint32)(unsafe.Pointer(p))))
		p += size
	}
}

type Student struct {
	Name string
	Age  int
}

func Test_Struct_Pointer(t *testing.T) {
	s := Student{}
	s.Name = "Peter"
	s.Age = 33

	pStudent := unsafe.Pointer(&s)

	// The entire object is converted to a pointer, and the
	// default is to get the first property
	name := (*string)(unsafe.Pointer(pStudent))

	t.Logf("Name Pointer: %p value: %s ", name, *name)

	// Using Offsetof to get the offset of age attribute
	// to get the attribute

	age := (*int)(unsafe.Pointer(uintptr(pStudent) + unsafe.Offsetof(s.Age)))
	t.Log("Offsetof: ", unsafe.Offsetof(s.Age))
	t.Logf("Age pointer: %p value: %d \n", age, *age)

	// Modify the value of the pointer
	*name = "Mary"
	*age = 20

	t.Log("Student: ", s)

}

type Teacher struct {
	name string
	age  int
}

func Test_Get_Private_Attribute_By_Offsetof(t *testing.T) {
	teacher_y := Teacher{"yangwawa", 47}

	pt := unsafe.Pointer(&teacher_y)

	name := (*string)(unsafe.Pointer(pt))

	t.Log(teacher_y.age)

	t.Log("Private attribute name : ", *name)

	age := (*int)(unsafe.Pointer(uintptr(pt) + unsafe.Offsetof(teacher_y.age)))

	t.Log("Private attribute age : ", *age)
}

func Test_Sizeof(t *testing.T) {
	array := []int{0, 1, -2, 3, 4}

	pointer := &array[0]

	t.Log(*pointer, " ")

	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		t.Log(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
}

func Test_Int_Pointer(t *testing.T) {
	var cookieMaxAge int64 = 3600 * 2
	ptr := uintptr(unsafe.Pointer(&cookieMaxAge))
	t.Logf("pointer ptr: %#v, value: %d ", ptr, *(*int)(unsafe.Pointer(&cookieMaxAge)))

}
