package pattern

import "fmt"



//_______________________________________________________________________
// Интерфейс посетителя, с методами "посещения" для каждого класса элемента, который существует в программе

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForrectangle(*Rectangle)
}

//_______________________________________________________________________
// интерфейс элементов

type Shape interface {
	GetType() string
	Accept(Visitor)
}

//_______________________________________________________________________
// AreaCalculator

type AreaCalculator struct {
	Area int
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) VisitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

//_______________________________________________________________________
// MiddleCoordinates 

type MiddleCoordinates struct {
	X int
	Y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) VisitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

//_______________________________________________________________________
// структура элемента квадрат

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

//_______________________________________________________________________
// структура элемента треугольник

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForrectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}

//_______________________________________________________________________
// структура элемента круг

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

