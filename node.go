package quadtree

const (
	MinLongitude float32 = -180.0
	MaxLongitude float32 = 180.0
	MinLatitude  float32 = -90.0
	MaxLatitude  float32 = 90.0

	Capacity = 8
	MaxDepth = 8
)

type (
	Node struct {
		bounds *Rectangle
		depth  int
		points []*Point
		parent *Node
		childs [4]*Node
	}
)

// NewQuadTree creates QuadTree. With bounding box for whole world and maximum depth 8.
// In case you need special settings for QuadTree you can use NewQuadTreeNode directly.
func NewQuadTree() *Node {
	return NewNode(
		NewRectangle(MinLongitude, MaxLongitude, MinLatitude, MaxLatitude),
		0,
		nil,
	)
}

// NewQuadTreeNode creates QuadTree node.
func NewNode(rectangle *Rectangle, depth int, parent *Node) *Node {
	return &Node{
		bounds: rectangle,
		depth:  depth,
		parent: parent,
	}
}

// split creates 4 quadrants in current node and moves point to it's children.
func (n *Node) split() {
	if n.childs[0] != nil {
		return
	}

	// Crete top left quadrant (north-west).
	boxNorthWest := NewRectangle(
		n.bounds.minimumX,
		(n.bounds.minimumX+n.bounds.minimumX)/2,
		(n.bounds.minimumY+n.bounds.maximumY)/2,
		n.bounds.maximumY,
	)
	n.childs[0] = NewNode(boxNorthWest, n.depth+1, n)

	// Create top right quadrant (north-east).
	boxNorthEast := NewRectangle(
		(n.bounds.minimumX+n.bounds.minimumX)/2,
		n.bounds.maximumX,
		(n.bounds.minimumY+n.bounds.maximumY)/2,
		n.bounds.maximumY,
	)
	n.childs[1] = NewNode(boxNorthEast, n.depth+1, n)

	// Create bottom left quadrant (south-west).
	boxSouthWest := NewRectangle(
		n.bounds.minimumX,
		(n.bounds.minimumX+n.bounds.minimumX)/2,
		n.bounds.minimumY,
		(n.bounds.minimumY+n.bounds.maximumY)/2,
	)
	n.childs[2] = NewNode(boxSouthWest, n.depth+1, n)

	// Create bottom right quadrant (south-east).
	boxSouthEast := NewRectangle(
		(n.bounds.minimumX+n.bounds.minimumX)/2,
		n.bounds.maximumX,
		n.bounds.minimumY,
		(n.bounds.minimumY+n.bounds.maximumY)/2,
	)
	n.childs[3] = NewNode(boxSouthEast, n.depth+1, n)

	// Reinsert points to child nodes.
	for _, p := range n.points {
		for _, child := range n.childs {
			if child.Insert(p) {
				break
			}
		}
	}
	n.points = nil
}

// Search will return all the points within bounding box definition.
func (n *Node) Search(a *Rectangle) []*Point {
	if !n.bounds.Intersect(a) {
		return nil
	}

	// Add points in current node.
	var results []*Point
	for _, p := range n.points {
		if a.FallsIn(p) {
			results = append(results, p)
		}
	}

	// Exit if child node does not exists.
	if n.childs[0] == nil {
		return results
	}

	// Range over child nodes recursively and search for points.
	for _, child := range n.childs {
		results = append(results, child.Search(a)...)
	}

	return results
}

// Insert will try to insert point into the QuadTree.
func (n *Node) Insert(p *Point) bool {
	// Check if point falls in bound of current node.
	if !n.bounds.FallsIn(p) {
		return false
	}

	// We insert point in current node, if there is space available or if we are on MaxDepth.
	if n.childs[0] == nil {
		if len(n.points) < Capacity || n.depth == MaxDepth {
			n.points = append(n.points, p)
			return true
		}

		// Split full node in 4 quadrants.
		n.split()
	}

	// Range over child nodes recursively and try to insert point.
	for _, child := range n.childs {
		if child.Insert(p) {
			return true
		}
	}

	return false
}

// Update will update the location of a point within the tree.
func (n *Node) Update(p *Point, np *Point) bool {
	// Check if point falls in bound of current node.
	if !n.bounds.FallsIn(p) {
		return false
	}

	// Update point in current node, if there are no child nodes.
	if n.childs[0] == nil {
		for i, val := range n.points {
			if val != p {
				continue
			}

			// Update coordinates to new point.
			p.x = np.x
			p.y = np.y

			// Check if new position still falls in current node.
			if n.bounds.FallsIn(np) {
				return true
			}

			// Remove and reinsert point to matching node.
			n.points = append(n.points[:i], n.points[i+1:]...)
			return n.reInsert(p)
		}
		return false
	}

	// Range over child nodes recursively and try to update point.
	for _, child := range n.childs {
		if child.Update(p, np) {
			return true
		}
	}

	return false
}

// reInsert inserts point in to matching node.
func (n *Node) reInsert(p *Point) bool {
	// Try to insert point in child nodes.
	if n.Insert(p) {
		return true
	}

	// We climbed to root node.
	if n.parent == nil {
		return false
	}

	// Try to insert point in parent node.
	return n.parent.reInsert(p)
}

// Remove recursively over nodes and tries to remove a point from the QuadTree.
func (n *Node) Remove(p *Point) bool {
	// Check if point falls in bound of current node.
	if !n.bounds.FallsIn(p) {
		return false
	}

	// Remove point in current node, if there are no child nodes.
	if n.childs[0] == nil {
		for i, val := range n.points {
			if val != p {
				continue
			}

			// Remove point from points array.
			n.points = append(n.points[:i], n.points[i+1:]...)

			return true
		}
		return false
	}

	// Range over child nodes recursively and try to delete point.
	for _, child := range n.childs {
		if child.Remove(p) {
			return true
		}
	}

	return false
}
