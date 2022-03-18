// 自动生成模板Key
package autocode

import (
	"raptor/server/global"
)

// Key 结构体
// 如果含有time.Time 请自行import time包
type Key struct {
      global.GVA_MODEL
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`
      Region  string `json:"region" form:"region" gorm:"column:region;comment:区域/地址;size:255;"`
      Keyid  string `json:"keyid" form:"keyid" gorm:"column:keyid;comment:钥匙;size:255;"`
      Secret  string `json:"secret" form:"secret" gorm:"column:secret;comment:密钥;size:255;"`
      Other  string `json:"other" form:"other" gorm:"column:other;comment:其他;size:255;"`
}


// TableName Key 表名
func (Key) TableName() string {
  return "key"
}

