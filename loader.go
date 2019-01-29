package oci

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

var schemaFS = FS(false)

func loadSchema(file string) *gojsonschema.Schema {
	path := fmt.Sprintf("file:///%s", file)
	loader := gojsonschema.NewReferenceLoaderFileSystem(path, schemaFS)
	schema, _ := gojsonschema.NewSchema(loader)
	return schema
}

// ImageIndexSchema provides a JSONLoader that can be used to validate an image
// index.
var ImageIndexSchema = func() *gojsonschema.Schema {
	return loadSchema("image-index.schema.json")
}

// ImageManifestSchema provides a JSONLoader that can be used to validate an
// image manifest.
var ImageManifestSchema = func() *gojsonschema.Schema {
	return loadSchema("image-manifest.schema.json")
}

// ImageConfigSchema provides a JSONLoader that can be used to validate an image
// config.
var ImageConfigSchema = func() *gojsonschema.Schema {
	return loadSchema("image-config.schema.json")
}

// TagListSchema provides a JSONLoader that can be used to validate an tag list.
var TagListSchema = func() *gojsonschema.Schema {
	return loadSchema("tag-list.schema.json")
}

// CatalogSchema provides a JSONLoader that can be used to validate a catalog of
// repositories.
var CatalogSchema = func() *gojsonschema.Schema {
	return loadSchema("catalog.schema.json")
}

// ErrorSchema provides a JSONLoader that can be used to validate an
// OCI-compliant error response.
var ErrorSchema = func() *gojsonschema.Schema {
	return loadSchema("errors.schema.json")
}

// contentDescriptorSchema provides a JSONLoader that can be used to validate a
// content descriptor.
var contentDescriptorSchema = func() *gojsonschema.Schema {
	return loadSchema("content-descriptor.schema.json")
}
