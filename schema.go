package schema

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

var schemaFS = _escFS(false)

func loadSchema(file string) *gojsonschema.Schema {
	path := fmt.Sprintf("file:///%s", file)
	loader := gojsonschema.NewReferenceLoaderFileSystem(path, schemaFS)
	schema, _ := gojsonschema.NewSchema(loader)
	return schema
}

// ImageIndexSchema provides a gojsonschema.Schema that can be used to validate
// an image index.
func ImageIndexSchema() *gojsonschema.Schema {
	return loadSchema("image-index.schema.json")
}

// ImageManifestSchema provides a gojsonschema.Schema that can be used to
// validate an image manifest.
func ImageManifestSchema() *gojsonschema.Schema {
	return loadSchema("image-manifest.schema.json")
}

// ImageConfigSchema provides a gojsonschema.Schema that can be used to validate
// an image config.
func ImageConfigSchema() *gojsonschema.Schema {
	return loadSchema("image-config.schema.json")
}

// TagListSchema provides a gojsonschema.Schema that can be used to validate an
// tag list.
func TagListSchema() *gojsonschema.Schema {
	return loadSchema("tag-list.schema.json")
}

// CatalogSchema provides a gojsonschema.Schema that can be used to validate
// a catalog of repositories.
func CatalogSchema() *gojsonschema.Schema {
	return loadSchema("catalog.schema.json")
}

// ErrorSchema provides a gojsonschema.Schema that can be used to validate an
// OCI-compliant error response.
func ErrorSchema() *gojsonschema.Schema {
	return loadSchema("errors.schema.json")
}

// contentDescriptorSchema provides a gojsonschema.Schema that can be used to
// validate a content descriptor.
func contentDescriptorSchema() *gojsonschema.Schema {
	return loadSchema("content-descriptor.schema.json")
}
