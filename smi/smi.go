package smi

import "math"

//Rectangle ...
//the
//
// noun
//
// a plane figure with four straight sides and four right angles,
// especially one with unequal adjacent sides,
// in contrast to a square.
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle ...
//
//kolo or okrag bcoz there is no difference in english
//
type Circle struct {
	Radius float64
}

//Triangle ...
//
//like... three people doing ... stuff
type Triangle struct {
	Base, Height float64
}

//Shape is interface to find Area64 :)
type Shape interface {
	Area() float64
}

//Perimeter ...
//the continuous line forming the boundary
//of a closed geometrical figure.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

//Area ...
// the extent or measurement of a surface or piece of land.
//
//here: of a rectangje
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Area is
// the extent or measurement of a surface or piece of land.
//
//here: of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area is
// the same as everything else
func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

//Area ...
//
//noun
//
// the extent or measurement of a surface or piece of land.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
