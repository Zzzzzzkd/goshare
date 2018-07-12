package sum

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestSum(t *testing.T) {
	s := [][]int{
		[]int{1, 2, 3},
		[]int{2, 4, 6},
		[]int{3, 6, 9},
	}

	for _, v := range s {
		if v[2] != Sum(v[0], v[1]) {
			t.Logf("Sum(%v, %v) expected %v,but return %v", v[1], v[0], v[2], Sum(v[0], v[1]))
		}
	}

}

type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Age  int    `json:"age"`
}

var s = Student{
	Name: "zkd",
	Id:   0,
	Age:  1111,
}

func Benchmark(b *testing.B) {
	bb, _ := json.Marshal(s)
	var ss Student
	var err error
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(bb, &ss)
	}
	_ = err
}
func BenchmarkSlowSum(b *testing.B) {
	bb, _ := json.Marshal(s)
	var ss Student
	var err error
	for i := 0; i < b.N; i++ {
		err = jsoniter.Unmarshal(bb, &ss)
	}
	_ = err
}
