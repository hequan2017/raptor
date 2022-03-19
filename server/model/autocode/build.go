// 自动生成模板Build
package autocode

import (
	"raptor/server/global"
)

// Build 结构体
// 如果含有time.Time 请自行import time包
type Build struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:构建名称;size:255;"`
      Static  string `json:"static" form:"static" gorm:"column:static;comment:构建状态;size:255;"`
      Version  string `json:"version" form:"version" gorm:"column:version;comment:版本号;size:255;"`
      Info  string `json:"info" form:"info" gorm:"column:info;comment:构建信息;"`
      History  string `json:"history" form:"history" gorm:"column:history;comment:执行历史;"`
      Other  string `json:"other" form:"other" gorm:"column:other;comment:其他;"`
}


// TableName Build 表名
func (Build) TableName() string {
  return "build"
}

