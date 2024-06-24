package LSP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLSP_RectangleStandard(t *testing.T) {
	r := &Rectangle{}
	assertRectangle(t, r)

}

func TestLSP_SquareMustFail(t *testing.T) {
	s := &Square{}
	assertRectangle(t, s)
}

func TestLSP_SquareCorrect(t *testing.T) {
	s := &Square{}
	s.SetWidth(10)
	assertRectangle2(t, s, 100)

	r := &Rectangle{}
	r.SetWidth(10)
	r.SetHeight(5)
	assertRectangle2(t, r, 50)
}

func assertRectangle2(t *testing.T, s AreShaped, expected int) {
	assert.Equal(t, expected, s.GetArea())
}

func assertRectangle(t *testing.T, r RectangleShaped) {
	r.SetWidth(10)
	r.SetHeight(5)
	assert.Equal(t, 50, r.GetArea())
}
