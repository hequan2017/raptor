package core

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"raptor/server/task"
	"time"

	"github.com/gogf/gf/os/gcron"
	"go.uber.org/zap"
	"raptor/server/global"
	"raptor/server/initialize"
	"raptor/server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 raptor/server
	当前版本:V0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	定时任务在 core/server.go 里面设置 ，会定时更新 钉钉部门用户信息，同步阿里云资产
`, address)

	_, err := gcron.Add("0  10  16 * * *", func() { task.DingUpdate() })
	_, err = gcron.Add("0  10  16 * * *", func() { task.AssetUpdate() })
	if err != nil {
		fmt.Println(err.Error())
	}
	g.Dump(gcron.Entries())

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}
