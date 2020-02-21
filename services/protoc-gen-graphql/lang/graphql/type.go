package pgsgql

import pgs "github.com/lyft/protoc-gen-star"

func (c context) IsMap(f pgs.Field) bool {
	return f.Type().IsMap()
}
