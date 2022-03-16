package request

import (
	"raptor/server/model/common/request"
	"raptor/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
