package runtime

import (
	"fmt"
	"testing"
)

func TestData_String(t *testing.T) {
	m0 := New(0)
	fmt.Println(m0)
	m1 := New(6)
	m2 := New(8)
	fmt.Println(m1)
	n := 1.0
	for i := 0; i < 6; i++ {
		m1.Set(i, FloatData(n))
		n++
	}
	for i := 0; i < 8; i++ {
		m2.Set(i, FloatData(n))
		n++
	}
	fmt.Println(m1.Transpose())
	m1.Shape(3, 2)
	m2.Shape(2, 4)
	fmt.Println(m1.Transpose())
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

func TestData_Transpose1(t *testing.T) {

	m := New(3, 2)
	n := 1
	for i := 0; i < 3; i = i + 1 {
		_ = i
		for j := 0; j < 2; j = j + 1 {
			_ = j
			m.Get(i).Set(j, FloatData(n))
			n = n + 1
		}
	}

	fmt.Println(m.MulMatrix(m.Transpose()))
	fmt.Println(m.AddMatrix(m))
	fmt.Println(m.AddFloat(3.2))
	fmt.Println(m.MulFloat(3.2))
	fmt.Println(m.Sizes())
}

func TestFloatData(t *testing.T) {
	a := FloatData(2.2)
	fmt.Println(a.Sizes())
	fmt.Println(a.toString("Test toString: "))
	fmt.Println(a.AddFloat(FloatData(1.1)))
	fmt.Println(a.MulFloat(FloatData(1.1)))
	fmt.Println(a.ToDatar())
	fmt.Println(a.ToFloatData())

	test1 := func() {
		defer func() {
			if err := recover(); err == nil {
				panic("fail")
			} else {
				t.Log(err)
			}
		}()
		a.Get(1)
	}
	test1()
	test2 := func() {
		defer func() {
			if err := recover(); err == nil {
				panic("fail")
			} else {
				t.Log(err)
			}
		}()
		a.Set(1, FloatData(1.1))
	}
	test2()

}

//func TestData_Transpose2(t *testing.T) {
//
//	m := New(3)
//	n := 1
//	for i := 0; i < 3; i = i + 1 {
//		m.Set(i,FloatData(n))
//		n = n + 1
//		}
//	}
//	m1 := m.MulMatrix(m.Transpose())
//	fmt.Println(m1)
//}
