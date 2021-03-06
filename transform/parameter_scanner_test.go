package transform

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/morlay/goqcloud/ptr"
)

type Filter struct {
	Name   *string   `json:"Name"`
	Values []*string `json:"Values"`
}

func TestParametersScanner(t *testing.T) {
	tt := assert.New(t)

	req := &struct {
		Filter
		Ptr     *int      `name:"Ptr"`
		PtrName *string   `name:"PtrName"`
		Filters []*Filter `name:"Filters"`
		List    []*string `name:"List"`
	}{
		Filter: Filter{
			Name: ptr.String("name"),
		},
		PtrName: ptr.String("name"),
		Filters: []*Filter{
			{
				Name: ptr.String("key"),
				Values: []*string{
					ptr.String("v1"),
					ptr.String("v2"),
				},
			},
		},
		List: []*string{
			ptr.String("v1"),
			ptr.String("v2"),
		},
	}

	s := NewParameterScanner()
	s.Scan(reflect.ValueOf(req))

	tt.Equal(url.Values{
		"List.0":             {"v1"},
		"List.1":             {"v2"},
		"Filters.0.Name":     {"key"},
		"Filters.0.Values.0": {"v1"},
		"Filters.0.Values.1": {"v2"},
		"Name":               {"name"},
		"PtrName":            {"name"},
	}, s.Params())
}
