// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteSegmentURL generates an URL for the delete segment operation
type DeleteSegmentURL struct {
	FlagID    int64
	SegmentID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteSegmentURL) WithBasePath(bp string) *DeleteSegmentURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteSegmentURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteSegmentURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/flags/{flagID}/segments/{segmentID}"

	flagID := swag.FormatInt64(o.FlagID)
	if flagID != "" {
		_path = strings.Replace(_path, "{flagID}", flagID, -1)
	} else {
		return nil, errors.New("FlagID is required on DeleteSegmentURL")
	}
	segmentID := swag.FormatInt64(o.SegmentID)
	if segmentID != "" {
		_path = strings.Replace(_path, "{segmentID}", segmentID, -1)
	} else {
		return nil, errors.New("SegmentID is required on DeleteSegmentURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteSegmentURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteSegmentURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteSegmentURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteSegmentURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteSegmentURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteSegmentURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
