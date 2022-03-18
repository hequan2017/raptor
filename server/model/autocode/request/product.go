package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type ProductSearch struct{
    autocode.Product
    request.PageInfo
}