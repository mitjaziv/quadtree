package quadtree

type Point struct {
	x    float32
	y    float32
	data interface{}
}

// NewPoint generates a new *Point struct.
func NewPoint(x, y float32, d interface{}) *Point {
	return &Point{
		x:    x,
		y:    y,
		data: d,
	}
}
