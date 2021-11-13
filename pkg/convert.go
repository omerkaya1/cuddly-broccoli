package pkg

import (
	"context"

	"github.com/pkg/errors"
	gj "github.com/tidwall/gjson"
)

// ConvertJSONBytes accepts serialized JSON data and converts its internal objects to arrays
// according to the passed paths strings and returns a resulting byte slice
func ConvertJSONBytes(b []byte, paths ...string) ([]byte, error) {
	return traverse(context.Background(), b, paths...)
}

// ConvertJSONBytesContext accepts serialized JSON data and converts its internal objects to arrays
// according to the passed paths strings and returns a resulting byte slice
// It takes into account the context passed to it, so that any context-related actions are handled during execution
func ConvertJSONBytesContext(ctx context.Context, b []byte, paths ...string) ([]byte, error) {
	return traverse(ctx, b, paths...)
}

// ConvertJSONString accepts serialized stringified JSON data and converts its internal objects to arrays
// according to the passed paths strings and returns a resulting string
func ConvertJSONString(str string, paths ...string) (string, error) {
	out, err := traverse(context.Background(), []byte(str), paths...)
	return string(out), err
}

// ConvertJSONStringContext accepts serialized JSON data and converts its internal objects to arrays
// according to the passed paths strings and returns a resulting string
// It takes into account the context passed to it, so that any context-related actions are handled during execution
func ConvertJSONStringContext(ctx context.Context, str string, paths ...string) (string, error) {
	out, err := traverse(ctx, []byte(str), paths...)
	return string(out), err
}

func traverse(ctx context.Context, b []byte, paths ...string) ([]byte, error) {
	// First stage: json validation
	if !gj.ValidBytes(b) {
		return nil, errors.New("provided json data is invalid")
	}
	// Do some incredibly cool stuff
	return nil, nil
}
