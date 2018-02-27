package quadtree

type Rectangle struct {
	minimumX float32
	maximumX float32
	minimumY float32
	maximumY float32
}

func NewRectangle(minimumX, maximumX, minimumY, maximumY float32) *Rectangle {
	return &Rectangle{
		minimumX: minimumX,
		maximumX: maximumX,
		minimumY: minimumY,
		maximumY: maximumY,
	}
}

// FallsIn checks whether the point provided resides within the axis aligned bounding box.
func (r *Rectangle) FallsIn(p *Point) bool {
	if p.x >= r.minimumX &&
		p.x <= r.maximumX &&
		p.y >= r.minimumY &&
		p.y <= r.maximumY {

		return true
	}
	return false
}

// Intersect checks whether two axis aligned bounding boxes overlap.
func (r *Rectangle) Intersect(bounds *Rectangle) bool {
	if r.minimumX > bounds.maximumX ||
		r.maximumX < bounds.minimumX ||
		r.minimumY > bounds.maximumY ||
		r.maximumY < bounds.minimumY {

		return false
	}
	return true
}
