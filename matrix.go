package matrix

import "errors"

type Size struct {
	Cols int
	Rows int
}

type Matrix[T int | float64] struct {
	size  *Size
	array [][]T
}

type Vector[T int | float64] struct {
	Matrix[T]
	size  *Size
	array [][]T
}

func NewMatrix[T int | float64](rowsCount int, colsCount int) (*Matrix[T], error) {
	if rowsCount == 0 && colsCount == 0 {
		return nil, errors.New("row and cols count cannot be zero")
	}

	var arr [][]T

	for i := 0; i < rowsCount; i++ {
		arr = append(arr, []T{})
		for j := 0; j < colsCount; j++ {
			arr[i] = append(arr[i], 0)
		}
	}

	return &Matrix[T]{
		size:  &Size{Cols: colsCount, Rows: rowsCount},
		array: arr,
	}, nil
}

func NewVector[T int | float64](dimension int) *Vector[T] {
	var v [][]T

	for i := 0; i < dimension; i++ {
		v = append(v, []T{})
		v[i] = append(v[i], 0)
	}

	return &Vector[T]{
		size:  &Size{Cols: 1, Rows: dimension},
		array: v,
	}
}

func (m *Matrix[T]) Rows() int {
	return m.size.Rows
}

func (m *Matrix[T]) Cols() int {
	return m.size.Cols
}

func (m *Matrix[T]) Get(row int, col int) T {
	if !m.validIndex(row, col) {
		panic("index is out of array")
	}

	return m.array[row][col]
}

func (m *Matrix[T]) Set(row int, col int, val T) {
	if !m.validIndex(row, col) {
		panic("index is out of array")
	}
	m.array[row][col] = val
}

func (m *Matrix[T]) validIndex(row int, col int) bool {
	return m.validCol(col) && m.validRow(row)
}

func (m *Matrix[T]) validRow(row int) bool {
	return 0 <= row && row < m.size.Rows
}

func (m *Matrix[T]) validCol(col int) bool {
	return 0 <= col && col < m.size.Cols
}

func (m *Matrix[T]) Multiplicate(other *Matrix[T]) (*Matrix[T], error) {
	if m.size.Cols != other.size.Rows {
		return nil, errors.New("rows and cols count are not equal")
	}

	res, err := NewMatrix[T](m.size.Rows, other.size.Cols)

	if err != nil {
		return nil, err
	}

	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < other.size.Cols; j++ {
			for k := 0; k < m.size.Cols; k++ {
				res.array[i][j] += m.array[i][k] * other.array[k][j]
			}
		}
	}

	return res, nil
}

func (m *Matrix[T]) ScalarMultiplication(val T) *Matrix[T] {
	res := *m

	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			res.array[i][j] *= val
		}
	}

	return &res
}

func (m *Matrix[T]) Dot(other *Matrix[T]) (*Matrix[T], error) {
	if m.size.Rows != other.size.Rows || m.size.Cols != other.size.Cols {
		return nil, errors.New("different sizes")
	}

	res, _ := NewMatrix[T](m.size.Rows, 1)

	for i := 0; i < m.size.Rows; i++ {
		var s T = 0
		for j := 0; j < m.size.Cols; j++ {
			s += m.Get(i, j) * other.Get(i, j)
		}
		res.Set(i, 0, s)
	}

	return res, nil
}

func (v *Vector[T]) Dot(other *Vector[T]) (T, error) {
	if v.size.Rows != other.size.Rows {
		return 0, errors.New("different sizes")
	}

	var s T = 0

	for i := 0; i < v.size.Rows; i++ {
		s += v.Get(i, 0) * other.Get(i, 0)
	}

	return s, nil
}

func (m *Matrix[T]) Add(other *Matrix[T]) (*Matrix[T], error) {
	if m.size.Rows != other.size.Rows || m.size.Cols != other.size.Cols {
		return nil, errors.New("different sizes")
	}

	res := *m

	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			res.array[i][j] += other.array[i][j]
		}
	}

	return &res, nil
}

func (m *Matrix[T]) Subtruct(other *Matrix[T]) (*Matrix[T], error) {
	if m.size.Rows != other.size.Rows || m.size.Cols != other.size.Cols {
		return nil, errors.New("different sizes")
	}

	res := *m

	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			res.array[i][j] -= other.array[i][j]
		}
	}

	return &res, nil
}

func (m *Matrix[T]) Transpose() *Matrix[T] {
	res, _ := NewMatrix[T](m.size.Cols, m.size.Rows)

	for i := 0; i < m.size.Rows; i++ {
		for j := 0; j < m.size.Cols; j++ {
			res.array[j][i] = m.array[i][j]
		}
	}

	return res
}

func (m *Matrix[T]) Exec(f func(x T) T, col int) (*Matrix[T], error) {
	res := *m

	if col == -1 {
		for i := 0; i < m.size.Rows; i++ {
			for j := 0; j < m.size.Cols; j++ {
				res.array[i][j] = f(res.array[i][j])
			}
		}
	} else {
		if !m.validCol(col) {
			return nil, errors.New("col is out of array")
		}
		for i := 0; i < m.size.Rows; i++ {
			res.array[i][col] = f(res.array[i][col])
		}
	}

	return &res, nil
}

func (m *Matrix[T]) GetRow(row int) *Matrix[T] {
	if !m.validRow(row) {
		panic("row is out of array")
	}

	res, _ := NewMatrix[T](1, m.size.Cols)
	for i := 0; i < m.size.Cols; i++ {
		res.array[0][i] = m.array[row][i]
	}

	return res
}

func (m *Matrix[T]) GetCol(col int) *Matrix[T] {
	if !m.validCol(col) {
		panic("col is out of array")
	}

	res, _ := NewMatrix[T](m.size.Rows, 1)
	for i := 0; i < m.size.Rows; i++ {
		res.array[i][0] = m.array[i][col]
	}

	return res
}
