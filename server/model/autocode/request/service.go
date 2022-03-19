package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type ServiceSearch struct{
    autocode.Service
    request.PageInfo
}