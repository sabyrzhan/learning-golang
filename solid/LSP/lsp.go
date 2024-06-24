package LSP

// LSP states the child classes must be substitutable without changing the parent's behavior
// Popular example of the violation of the LSP is Rectangle-Square problem
// When we have rectangle and square which extends it, square changes rectangle's behavior by mutating
// both width and height fields when one of the field is changed. In the end this produces different results
// when other common methods outcomes depend on these fields.
// The correct approach would be to separate the fields into their own struct, so they affect their own behavior not
// common the ones.

type RectangleShaped interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
	GetArea() int
}

type AreShaped interface {
	GetArea() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) GetArea() int {
	return r.Width * r.Height
}

type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height int) {
	s.Height = height
	s.Width = height
}

type ShapeType int

const (
	RectangleType ShapeType = iota
	SquareType
)

func CreateRectangle(width int, height int, shapeType ShapeType) RectangleShaped {
	switch shapeType {
	case RectangleType:
		return &Rectangle{width, height}
	case SquareType:
		s := &Square{}
		s.SetWidth(width)
		return s
	}

	return nil
}

func CreateRectangle2(width int, height int, shapeType ShapeType) AreShaped {
	switch shapeType {
	case RectangleType:
		return &Rectangle{width, height}
	case SquareType:
		s := &Square{}
		s.SetWidth(width)
		return s
	}

	return nil
}
