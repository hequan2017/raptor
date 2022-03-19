package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type BuildSearch struct{
    autocode.Build
    request.PageInfo
}