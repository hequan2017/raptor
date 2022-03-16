package v1

import (
	"raptor/server/api/v1/autocode"
	"raptor/server/api/v1/example"
	"raptor/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
