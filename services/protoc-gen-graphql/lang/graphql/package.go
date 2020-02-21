package pgsgql

import (
	"regexp"

	pgs "github.com/lyft/protoc-gen-star"
)

var nonAlphaNumPattern = regexp.MustCompile("[^a-zA-Z0-9]")

// ImportPath generates a filepath for importing an entity
// For example, a graphql file in "./foo/example.graphql" that wants to import a file in "./bar/example.graphql" needs
// to use the filepath "../bar/example.graphql" and not "bar/example.graphql"
func (c context) ImportPath(e pgs.Entity) pgs.FilePath {
	path := c.p.Str("import_prefix") + e.Name().String()
	return pgs.FilePath(path)
}

func (c context) OutputPath(e pgs.Entity) pgs.FilePath {
	return e.File().InputPath().SetExt(".graphql")
}
