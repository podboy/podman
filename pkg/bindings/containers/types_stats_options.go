// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *StatsOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *StatsOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithStream set field Stream to given value
func (o *StatsOptions) WithStream(value bool) *StatsOptions {
	o.Stream = &value
	return o
}

// GetStream returns value of field Stream
func (o *StatsOptions) GetStream() bool {
	if o.Stream == nil {
		var z bool
		return z
	}
	return *o.Stream
}

// WithInterval set field Interval to given value
func (o *StatsOptions) WithInterval(value int) *StatsOptions {
	o.Interval = &value
	return o
}

// GetInterval returns value of field Interval
func (o *StatsOptions) GetInterval() int {
	if o.Interval == nil {
		var z int
		return z
	}
	return *o.Interval
}