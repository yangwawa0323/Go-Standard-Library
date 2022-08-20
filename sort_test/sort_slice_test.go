package sort_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person base on
// the Age feild.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

/////////////////////////////////////////////////////////////////

var people []Person = []Person{
	{"Bob", 31},
	{"John", 42},
	{"Michael", 17},
	{"Jenny", 26},
}

func Test_Sort_ByAge(t *testing.T) {
	t.Log(people)

	sort.Sort(ByAge(people))
	t.Log(people)
}

func Test_Sort_Slice(t *testing.T) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	t.Log(people)

	// whether a slice is sorted is depend on the Less func of SliceIsSorted
	// by default Float64sAreSorted, IsSorted return true only in ascending order.
	t.Logf("%v is sorted float64: %v \n", people,
		sort.SliceIsSorted(people, func(i, j int) bool {
			return people[i].Age > people[j].Age
		}))
}

func Test_Sort_Float64(t *testing.T) {
	s1 := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s1)
	t.Log(s1)

	s2 := []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s2)
	t.Log(s2)

	t.Logf("%v is sorted float64: %v \n", s1, sort.Float64sAreSorted(s1))
	t.Logf("%v is sorted float64: %v \n", s2, sort.Float64sAreSorted(s2))

}

func Test_Sort_Search(t *testing.T) {
	s1 := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	// sort.Float64s(s1)

	fdIdx := sort.Search(len(s1), func(i int) bool {
		return s1[i] > 0
	})
	t.Logf("Found the first element greate than zero is: %f", s1[fdIdx])
}
