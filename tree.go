package cuddly_broccoli

import (
	"sort"
	"strings"
)

// branch is an object that contains data used by the converter to convert objects in the JSON data tree
type branch struct {
	toArray  bool      // indicates whether to convert an object to array or not
	objName  string    // the object name in the tree
	elemName string    // an optional name for an element
	parent   *branch   // the link to the parent element
	children []*branch // the slice of child elements
}

// newTree accepts a slice of strings that contains paths to data to convert and returns a root branch of the tree
func newTree(paths []string) *branch {
	// Sort strings to make it easier for the parser to insert data
	sort.Strings(paths)
	// Initialise the tree
	tree := branch{
		children: make([]*branch, 0, len(paths)),
	}
	var params []string
	// Iterate over the paths slice to fill up the tree
	for i := range paths {
		// Split the whole path to parts and feed them to the insert method
		if params = strings.Split(paths[i], "."); len(params) > 0 {
			tree.insert(params)
		}
	}
	return &tree
}

// newBranch generates a branch out of a path portion
func newBranch(s string) *branch {
	d := branch{
		toArray: true,
	}
	// Get the node name and other attributes
	switch j := strings.IndexAny(s, "#*"); j {
	case -1:
		// If not found, then it's a simple path without array iteration
		d.objName, d.toArray = s, false
	default:
		// We check whether the node name should be included or not
		d.objName = s[:j]
		if len(s[j:]) == 1 {
			if s[j] == '#' {
				d.elemName = "number"
			}
			if s[j] == '*' {
				d.elemName = "name"
			}
		}
	}
	return &d
}

// getByName searches the branch child branches for a branch with a particular name
func (b *branch) getByName(name string) *branch {
	for i := range b.children {
		if b.children[i].objName == name {
			return b.children[i]
		}
	}
	return nil
}

// insert accepts a slice of path parts
func (b *branch) insert(parts []string) {
	if len(parts) < 1 {
		return
	}
	// Create a node
	nb := newBranch(parts[0])

	// If we find the param in the list of children, at this level,
	// then we append the rest to the children list
	if temp := b.getByName(nb.objName); temp != nil {
		nb.parent = temp
		if len(parts) > 1 {
			temp.insert(parts[1:])
		}
		return
	}

	// If we are here, then we did not find the branch in the tree
	// and should therefore add a fully grown child to the branch
	b.children = append(b.children, nb)
	nb.parent = b
	if len(parts) > 1 {
		nb.insert(parts[1:])
	}
}

// validate returns an error if a branch is invalid
// TODO: come up with some sort of validation for the branch object
func (b *branch) validate() error {
	return nil
}
