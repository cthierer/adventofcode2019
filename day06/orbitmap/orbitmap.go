package orbitmap

import "fmt"

type node struct {
	parent   *node
	children []*node
}

func (n *node) Depth() int {
	ptr := n.parent
	count := 0
	for ptr != nil {
		count++
		ptr = ptr.parent
	}
	return count
}

// SpaceObject is an object in orbit in space.
type SpaceObject struct {
	node
	Label string
}

// OrbitsAround puts this object in order around the other.
func (o *SpaceObject) OrbitsAround(p *SpaceObject) {
	o.node.parent = &p.node
	p.children = append(p.node.children, &o.node)
}

// NumOrbits is the number of orbits (both direct and indirect).
func (o *SpaceObject) NumOrbits() int {
	return o.node.Depth()
}

func (o *SpaceObject) ancestors() []*node {
	var ancestors []*node
	ptr := o.node.parent

	for ptr != nil {
		ancestors = append(ancestors, ptr)
		ptr = ptr.parent
	}

	return ancestors
}

// OrbitMap maintains a mapping of all objects in orbit.
type OrbitMap struct {
	index map[string]*SpaceObject
}

// New creates a new OrbitMap.
func New() *OrbitMap {
	orbitMap := OrbitMap{
		index: make(map[string]*SpaceObject),
	}
	return &orbitMap
}

// Add a node to the map.
func (m *OrbitMap) Add(obj *SpaceObject) {
	m.index[obj.Label] = obj
}

// Traverse iterates over the map and calls the provided callback.
func (m *OrbitMap) Traverse(cb func(*SpaceObject)) {
	for _, o := range m.index {
		cb(o)
	}
}

// PathBetween calculates the length of the path between two nodes.
func (m *OrbitMap) PathBetween(startLabel, endLabel string) (int, error) {
	startAncestors := m.index[startLabel].ancestors()
	endAncestors := m.index[endLabel].ancestors()

	for startI, startAncestor := range startAncestors {
		for endI, endAncestor := range endAncestors {
			if startAncestor == endAncestor {
				return startI + endI + 2, nil
			}
		}
	}

	return 0, fmt.Errorf("no common ancestor found between %s and %s", startLabel, endLabel)
}

// RecordObject records an object in orbit in the map.
func RecordObject(orbitMap *OrbitMap, childLabel, parentLabel string) (*OrbitMap, error) {
	parentObj := orbitMap.index[parentLabel]
	if parentObj == nil {
		parentObj = &SpaceObject{Label: parentLabel}
		orbitMap.Add(parentObj)
	}
	childObj := orbitMap.index[childLabel]
	if childObj == nil {
		childObj = &SpaceObject{Label: childLabel}
		orbitMap.Add(childObj)
	}
	childObj.OrbitsAround(parentObj)
	return orbitMap, nil
}
