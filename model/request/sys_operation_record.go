package request

import "gin-learn/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
