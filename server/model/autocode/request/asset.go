package request

import (
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
)

type AssetSearch struct{
    autocode.Asset
    request.PageInfo
}