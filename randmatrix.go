package randmatrix

import (
	"errors"
	"math/rand"
)

var (
	ErrNotEnoughRandomNumbers = errors.New("not enough random numbers")
)

func RandMatrix(cols int, rows int, randCap int) ([][]int, error) {
	totalCount := cols * rows

	if randCap < totalCount {
		return nil, ErrNotEnoughRandomNumbers
	}

	mat := newMatrix(cols, rows)
	got := make(map[int]struct{}, totalCount)
	gotCount := 0
	var random int

	for gotCount < totalCount {
		random = rand.Intn(randCap)

		got[random] = struct{}{}

		if gotCount == len(got) {
			continue
		}

		gotCount++

		mat.push(random)
	}

	return mat.matrix(), nil
}

type matrix struct {
	mat  [][]int
	x    int // column index
	y    int // row index
	cols int
	rows int
}

func newMatrix(cols int, rows int) *matrix {
	mat := make([][]int, cols)

	for i := 0; i < cols; i++ {
		mat[i] = make([]int, rows)
	}

	return &matrix{
		mat:  mat,
		cols: cols,
		rows: rows,
	}
}

func (m *matrix) matrix() [][]int {
	return m.mat
}

func (m *matrix) push(num int) {
	m.mat[m.x][m.y] = num

	m.x++
	m.x %= m.cols

	if m.x == 0 {
		m.y++
		m.y %= m.rows
	}
}
