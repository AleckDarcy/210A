package runtime

import (
	"fmt"
	"reflect"
)

type Datar interface {
	Get(i int) Datar
	Set(i int, data Datar)

	Sizes() []int

	toString(indent string) string
}

type Data struct {
	sizes     []int
	numSizes  []int
	numSize   int
	subMSizes []int

	data  []FloatData
	subMs []*Data

	array     []FloatData
	subMArray []Data
}

type FloatData float64

func (d FloatData) Get(i int) Datar {
	panic("cannot call index on FloatData")
}

func (d FloatData) Set(i int, data Datar) {
	panic("cannot set value on FloatData")
}

func (d FloatData) Sizes() []int {
	return []int{}
}

func (d FloatData) toString(indent string) string {
	return fmt.Sprintf("%s%f", indent, d)
}

func (d FloatData) AddFloat(data FloatData) FloatData {
	return d + data
}

func (d FloatData) MulFloat(data FloatData) FloatData {
	return d * data
}

func (d FloatData) ToDatar() Datar {
	return d
}

func (d FloatData) ToFloatData() FloatData {
	return d
}

func New(sizes ...int) *Data {
	if len(sizes) == 0 {
		return nil
	}

	m := &Data{}
	m.Shape(sizes...)

	return m
}

func (m *Data) Shape(sizes ...int) *Data {
	if len(sizes) == 0 {
		panic("sizes are empty")
	}

	lenSizes := len(sizes)
	//if lenSizes == len(m.sizes) {
	//	equal := true
	//	for i := 0; i < lenSizes; i++ {
	//		if equal = sizes[i] == m.sizes[i]; !equal {
	//			break
	//		}
	//	}
	//
	//	if equal {
	//		return m
	//	}
	//}

	numSize, numSizes := 1, make([]int, lenSizes)
	subMSize, subMSizes := 0, make([]int, lenSizes-1)

	for i := lenSizes - 1; i >= 0; i-- {
		numSize *= sizes[i]
		numSizes[i] = numSize

		if lenSizes > 1 {
			if i == 0 {
				subMSizes[i] = sizes[i]
				subMSize += subMSizes[i]
			} else if i < lenSizes-1 {
				subMSizes[i] = sizes[i] * sizes[i-1]
				subMSize += subMSizes[i]
			}
		}
	}

	if numSize == 0 {
		panic("")
	}

	//fmt.Println("Shape subMSizes:", subMSizes, "subMSize", subMSize, "numSizes", numSizes, "numSize:", numSize)
	if m.numSize == 0 { // init
		array := make([]FloatData, numSize)
		m.numSize = numSize
		m.array = array
	} else {
		if numSize != m.numSize {
			panic("shape not match")
		}
	}

	m.sizes = sizes
	m.numSize, m.numSizes = numSize, numSizes
	m.subMSizes = subMSizes

	//if cap(m.subMArray) >= subMSize {
	//	m.subMArray = m.subMArray[:subMSize]
	//	m.subMs = m.subMs[:subMSize]
	//} else {
	m.subMArray = make([]Data, subMSize)
	m.subMs = make([]*Data, subMSize)
	for i := 0; i < subMSize; i++ {
		m.subMs[i] = &m.subMArray[i]
	}
	//}

	m.shape()
	if lenSizes > 1 {
		m.subMs = m.subMs[:subMSizes[0]]
	}

	return m
}

func (m *Data) shape() {
	//fmt.Println(level, "shape sizes", m.sizes, "subMSizes", m.subMSizes, "numSizes", m.numSizes, "arraySize", len(m.array))
	if len(m.subMSizes) == 0 {
		//fmt.Println(level, "set data")
		m.data = m.array[:m.sizes[0]]
	} else {
		subMOffset := m.subMSizes[0]
		numOffset := 0

		for i := 0; i < m.sizes[0]; i++ {
			//fmt.Printf("%d, %d pick array [%d: %d], subMsSize %d %d\n", level, i, numOffset, numOffset+m.numSizes[1], cap(m.subMs), len(m.subMs))
			subM := m.subMs[i]
			subM.sizes = m.sizes[1:]
			subM.numSizes = m.numSizes[1:]
			subM.subMSizes = m.subMSizes[1:]

			subM.array = m.array[numOffset : numOffset+m.numSizes[1]]
			numOffset += m.numSizes[1]

			if len(m.subMSizes) > 1 {
				//fmt.Printf("%d %d pick subMs [%d: %d]\n", level, i, subMOffset, subMOffset+m.sizes[1])
				subM.subMs = m.subMs[subMOffset : subMOffset+m.sizes[1]]
				subMOffset += m.sizes[1]
			}
			subM.shape()
		}
	}
	//fmt.Println(level, "return shape")
}

func (m *Data) Get(i int) Datar {
	if 0 <= i && i < m.sizes[0] {
		if len(m.sizes) == 1 {
			return m.data[i]
		}

		return m.subMs[i]
	}

	panic("index out of range")
}

func (m *Data) Set(i int, data Datar) {
	//fmt.Println(i, m.sizes)
	if 0 <= i && i < m.sizes[0] {
		if len(m.sizes) == 1 {
			if fData, ok := data.(FloatData); ok {
				m.data[i] = fData
			} else {
				panic("unreachable code")
			}
		} else if reflect.DeepEqual(m.sizes[1:], data.Sizes()) {
			m.subMs[i] = data.(*Data)
		} else {
			panic("sizes not match")
		}

		return
	}

	panic("index out of range")
}

func (m *Data) Sizes() []int {
	return m.sizes
}

func (m *Data) toString(indent string) string {
	//fmt.Println("sizes", m.sizes)
	if len(m.sizes) == 1 {
		//fmt.Println("wa")
		return fmt.Sprintf("%s%v", indent, m.data)
	}

	childStr := ""
	newIndent := fmt.Sprintf("\t%s", indent)
	for i := 0; i < m.sizes[0]; i++ {
		if i == 0 {
			childStr += fmt.Sprintf("%v", m.subMs[i].toString(newIndent))
		} else {
			childStr += fmt.Sprintf(",\n %v", m.subMs[i].toString(newIndent))
		}
	}

	return fmt.Sprintf(""+
		"%s[\n"+
		"%s\n"+
		"%s]",
		indent, childStr, indent,
	)
}

func (m *Data) String() string {
	return m.toString("")
}

func (m *Data) AddFloat(data FloatData) *Data {
	for i := range m.array {
		m.array[i] = m.array[i].AddFloat(data)
	}

	return m
}

func (m *Data) MulFloat(data FloatData) *Data {
	for i := range m.array {
		m.array[i] = m.array[i].MulFloat(data)
	}

	return m
}

func (m *Data) MulMatrix(data *Data) *Data {
	sizes, dataSizes := m.sizes, data.sizes

	if len(m.sizes) == 2 && len(data.sizes) == 2 && sizes[1] == dataSizes[0] {
		newM := New(sizes[0], dataSizes[1])
		for i := 0; i < sizes[0]; i++ {
			newM_i := newM.Get(i)
			for j := 0; j < dataSizes[1]; j++ {
				temp := FloatData(0)
				for k := 0; k < sizes[1]; k++ {
					numI, numJ := m.array[i*sizes[1]+k], data.array[k*dataSizes[1]+j]
					temp = temp.AddFloat(numI.MulFloat(numJ))
				}
				newM_i.Set(j, temp)
			}
		}

		return newM
	}

	panic("size not match for matrix multiplication")
}

func (m *Data) AddMatrix(data *Data) *Data {
	// todo

	newM := New(m.sizes...)

	for i := range m.array {
		newM.array[i] = m.array[i] + data.array[i]
	}

	return newM
}

func (m *Data) Transpose() *Data {
	if len(m.sizes) == 1 {
		newM := New(m.sizes[0], 1)
		for i := 0; i < m.sizes[0]; i++ {
			newM.Get(i).Set(0, m.Get(i))
		}

		return newM
	}
	if len(m.sizes) == 2 {
		newM := New(m.sizes[1], m.sizes[0])

		for i := 0; i < m.sizes[0]; i++ {
			mI := m.Get(i)
			for j := 0; j < m.sizes[1]; j++ {
				newM.Get(j).Set(i, mI.Get(j))
			}
		}

		return newM
	}

	panic("cannot transpose matrix with dimension greater than 2")
}
