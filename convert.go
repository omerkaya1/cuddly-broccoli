package cuddly_broccoli

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// ConvertJSONFromBytes accepts JSON data and converts its internal objects to arrays
func ConvertJSONFromBytes(b []byte, paths []string) ([]byte, error) {
	if !json.Valid(b) {
		return nil, errors.New("provided json data is invalid")
	}
	// First stage: create a tree for traversal
	tree := newTree(paths)
	return nil, tree.validate()
}

// ConvertJSONFromBytesContext accepts JSON data and converts its internal objects to arrays
func ConvertJSONFromBytesContext(ctx context.Context, b []byte, paths []string) ([]byte, error) {
	if !json.Valid(b) {
		return nil, errors.New("provided json data is invalid")
	}
	// First stage: create a tree for traversal
	tree := newTree(paths)
	return nil, tree.validate()
}
