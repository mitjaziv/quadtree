package quadtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Rectangle
func Test_Rectangle(t *testing.T) {
	// mock
	var x1 float32 = 15.0
	var x2 float32 = 16.0
	var y1 float32 = 12.0
	var y2 float32 = 13.0

	id := "test_id"

	// new rectangle
	r := NewRectangle(x1, x2, y1, y2)
	assert.NotNil(t, r, "expected rectangle not to be nil")

	// contains point
	p1 := NewPoint(1.0, 1.0, id)
	b1 := r.FallsIn(p1)
	assert.Equal(t, false, b1, "expected rectangle to not contain point")

	p2 := NewPoint(15.5, 12.5, id)
	b2 := r.FallsIn(p2)
	assert.Equal(t, true, b2, "expected rectangle to contain point")

	// intersect
	b3 := r.Intersect(NewRectangle(14.5, 15.5, 11.5, 12.5))
	assert.Equal(t, true, b3, "expected rectangles to intersect")

	b4 := r.Intersect(NewRectangle(4.5, 5.5, 1.5, 2.5))
	assert.Equal(t, false, b4, "expected rectangles not to intersect")

	// contains
	b5 := r.Intersect(NewRectangle(15.2, 15.8, 12.2, 12.8))
	assert.Equal(t, true, b5, "expected rectangles to intersect/contain")
}
