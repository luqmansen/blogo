// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
