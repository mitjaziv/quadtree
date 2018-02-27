package quadtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Point
func Test_Point(t *testing.T) {
	p := NewPoint(15.0, 16.0, "test_id")
	assert.NotNil(t, p, "expected point not to be nil")
}
