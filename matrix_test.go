package matrix

import (
	"math"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m, _ := NewMatrix[float64](2, 2)

	if m.size.Cols != 2 || m.size.Rows != 2 {
		t.Error("invalid rows or cols")
	}
}

func TestMatrixSetGet(t *testing.T) {
	m, values := getMatrix[float64](2, 2)
	_, err := m.Get(-1, 100)
	if err == nil {
		t.Error("error was not occurred")
	}

	k := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if val, _ := m.Get(i, j); val != values[k] {
				t.Errorf("invalid value. expected %f, found %f\n", values[k], val)
			}
			k++
		}
	}
}

func TestMatrixMultiplicate(t *testing.T) {
	m1, _ := NewMatrix[int](2, 2)
	m2, _ := NewMatrix[int](3, 3)

	_, err := m1.Multiplicate(m2)
	if err == nil {
		t.Error("error was not occurred")
	}

	m1, _ = getMatrix[int](2, 2)

	m2, _ = getMatrix[int](2, 3)

	m, err := m1.Multiplicate(m2)
	ans := []int{9, 12, 15, 19, 26, 33}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixScalarMultiplication(t *testing.T) {
	m, _ := getMatrix[int](2, 2)

	m.ScalarMultiplication(5)

	ans := []int{5, 10, 15, 20}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixAdd(t *testing.T) {
	m1, _ := getMatrix[int](2, 2)
	m2, _ := getMatrix[int](3, 3)

	if _, err := m1.Add(m2); err == nil {
		t.Error("error was not occurred")
	}

	m2, _ = getMatrix[int](2, 2)
	m, _ := m1.Add(m2)
	ans := []int{2, 4, 6, 8}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixTranspose(t *testing.T) {
	m, _ := getMatrix[int](2, 2)

	m = m.Transpose()

	ans := []int{1, 3, 2, 4}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixExec(t *testing.T) {
	m, _ := getMatrix[float64](2, 2)
	m.Set(0, 1, 4)
	m.Set(1, 0, 9)
	m.Set(1, 1, 16)

	m, _ = m.Exec(func(x float64) float64 { return math.Sqrt(x) }, -1)

	ans := []float64{1, 2, 3, 4}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %f, found %f\n", ans[k], val)
			}
			k++
		}
	}

	m.Set(0, 1, 4)
	m.Set(1, 1, 16)

	if _, err := m.Exec(func(x float64) float64 { return math.Sqrt(x) }, 100); err == nil {
		t.Error("error was not occurred")
	}

	m, _ = m.Exec(func(x float64) float64 { return math.Sqrt(x) }, 1)

	ans = []float64{1, 2, 3, 4}
	k = 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %f, found %f\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixCol(t *testing.T) {
	m, _ := getMatrix[int](2, 2)

	if _, err := m.GetCol(100); err == nil {
		t.Error("error was not occurred")
	}

	m, _ = m.GetCol(0)

	ans := []int{1, 3}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func TestMatrixRow(t *testing.T) {
	m, _ := getMatrix[int](2, 2)

	if _, err := m.GetRow(100); err == nil {
		t.Error("error was not occurred")
	}

	m, _ = m.GetRow(0)

	ans := []int{1, 2}
	k := 0
	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			if val, _ := m.Get(i, j); val != ans[k] {
				t.Errorf("invalid value. expected %d, found %d\n", ans[k], val)
			}
			k++
		}
	}
}

func getMatrix[T int | float64](rows int, cols int) (*Matrix[T], []T) {
	m, _ := NewMatrix[T](rows, cols)

	var values []T

	for i := 0; i < rows*cols; i++ {
		values = append(values, T(i+1))

	}

	k := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, values[k])
			k++
		}
	}
	return m, values
}
