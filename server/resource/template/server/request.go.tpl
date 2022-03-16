package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}