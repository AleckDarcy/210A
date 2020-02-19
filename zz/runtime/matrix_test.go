package runtime

import (
	"fmt"
	"testing"
)

func TestData_String(t *testing.T) {
	m1 := New(6)
	m2 := New(8)

	n := 1.0
	for i := 0; i < 6; i++ {
		m1.Set(i, FloatData(n))
		n++
	}
	for i := 0; i < 8; i++ {
		m2.Set(i, FloatData(n))
		n++
	}
	m1.Shape(3, 2)
	m2.Shape(2, 4)

	fmt.Println(m1.MulMatrix(m2))

	fmt.Println(m2.Shape(8))
}

func BenchmarkData_MulMatrix(b *testing.B) {
	m1 := New(6)
	m2 := New(8)

	n := 1.0
	for i := 0; i < 6; i++ {
		m1.Set(i, FloatData(n))
		n++
	}
	for i := 0; i < 8; i++ {
		m2.Set(i, FloatData(n))
		n++
	}
	m1.Shape(3, 2)
	m2.Shape(2, 4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m1.MulMatrix(m2)
	}
}
func BenchmarkShape(b *testing.B) {
	m := New(2, 4, 6)

	for i := 0; i < b.N; i++ {
		m.Shape(6, 4, 2)
		//m.Shape(1, 48)
		//m.Shape(48)
		//m.Shape(48, 1)
	}
}
