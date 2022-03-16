package request

import (
	"raptor/server/model/common/request"
	"raptor/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
