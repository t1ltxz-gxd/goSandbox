package tasks

import (
	"fmt"
	"math"
)

// Point - структура для представления точки с координатами x и y
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance - метод для вычисления расстояния между двумя точками
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func TwentyFourth() {
	p1 := NewPoint(2.0, 3.0)
	p2 := NewPoint(5.0, 7.0)

	distance := p1.Distance(p2)
	fmt.Println("Расстояние между точками:", distance)
}
