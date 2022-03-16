package request

import (
	"raptor/server/model/common/request"
	"raptor/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
