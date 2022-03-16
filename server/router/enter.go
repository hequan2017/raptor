package router

import (
	"raptor/server/router/autocode"
	"raptor/server/router/example"
	"raptor/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
