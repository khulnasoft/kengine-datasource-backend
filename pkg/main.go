package main

import (
	"os"

	"github.com/khulnasoft/kengine-plugin-sdk-go/backend/datasource"
	"github.com/khulnasoft/kengine-plugin-sdk-go/backend/log"
	"github.com/khulnasoft/kengine-starter-datasource-backend/pkg/plugin"
)

func main() {
	// Start listening to requests sent from Kengine. This call is blocking so
	// it won't finish until Kengine shuts down the process or the plugin choose
	// to exit by itself using os.Exit. Manage automatically manages life cycle
	// of datasource instances. It accepts datasource instance factory as first
	// argument. This factory will be automatically called on incoming request
	// from Kengine to create different instances of SampleDatasource (per datasource
	// ID). When datasource configuration changed Dispose method will be called and
	// new datasource instance created using NewSampleDatasource factory.
	if err := datasource.Manage("myorgid-simple-backend-datasource", plugin.NewSampleDatasource, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
