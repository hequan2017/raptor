package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type KeySearch struct{
    autocode.Key
    request.PageInfo
}