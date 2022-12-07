package interfaces

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

const PI = 3.14

func (r *Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r *Rectangle) perimeter() float64 {
	return (2 * (r.length + r.breadth))
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (s *Square) perimeter() float64 {
	return (4 * s.side)
}

func (c *Circle) area() float64 {
	return (PI * c.radius * c.radius)
}

func (c *Circle) perimeter() float64 {
	return (2 * PI * c.radius)
}
