package plugin_test

import (
	"context"
	"testing"

	"github.com/khulnasoft/kengine-plugin-sdk-go/backend"
	"github.com/khulnasoft/kengine-starter-datasource-backend/pkg/plugin"
)

// This is where the tests for the datasource backend live.
func TestQueryData(t *testing.T) {
	ds := plugin.SampleDatasource{}

	resp, err := ds.QueryData(
		context.Background(),
		&backend.QueryDataRequest{
			Queries: []backend.DataQuery{
				{RefID: "A"},
			},
		},
	)
	if err != nil {
		t.Error(err)
	}

	if len(resp.Responses) != 1 {
		t.Fatal("QueryData must return a response")
	}
}
