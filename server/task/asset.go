package task

import (
	"database/sql"
	gorm2 "gorm.io/gorm"
	"raptor/server/global"
	"strings"

	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"os"
	"raptor/server/model/autocode"
	"strconv"
	"time"
)

type DeletedAts sql.NullTime

var db *gorm.DB

func GetAliyunInstances(newInstanceList []string) ([]string, error) {

	var keys []autocode.Key
	_ = db.Find(&keys).Error
	for _, v := range keys {
		client, _ := ecs.NewClientWithAccessKey(v.Region, v.Keyid, v.Secret)
		request := ecs.CreateDescribeInstancesRequest()
		request.PageSize = "100"
		request.PageNumber = "1"
		response, _ := client.DescribeInstances(request)

		r := response.Instances
		for _, i := range r.Instance {
			var a autocode.Asset
			newInstanceList = append(newInstanceList, i.InstanceId)
			_ = db.Where("sn = ?", i.InstanceId).Where("deleted_at is null").Unscoped().First(&a).Error
			if a.Sn != "" {
				//local, _ := time.LoadLocation("Asia/Shanghai")
				ct, _ := time.Parse("2006-01-02T15:04Z", i.CreationTime)
				st, _ := time.Parse("2006-01-02T15:04Z", i.StartTime)
				asset := autocode.Asset{
					Sn:       i.InstanceId,
					Type:     "阿里云主机",
					Brand:    "阿里云",
					Model:    i.InstanceType,
					Hostname: i.HostName,
					System:   i.OSName,
					Intranet: i.NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress,
					Status:   i.Status,
					Public:    i.EipAddress.IpAddress,
					CreatedAt: ct,
					Cpu:       strconv.Itoa(i.Cpu),
					Memory:    strconv.Itoa(i.Memory / 1024),
					Region:    v.Region,
					Uptime:    &st,
				}
				db.Model(&autocode.Asset{}).Where("sn = ?", i.InstanceId).Updates(asset)
			} else {
				ct, _ := time.Parse("2006-01-02T15:04Z", i.CreationTime)
				st, _ := time.Parse("2006-01-02T15:04Z", i.StartTime)
				asset := autocode.Asset{
					Sn:       i.InstanceId,
					Type:     "阿里云主机",
					Brand:    "阿里云",
					Model:    i.InstanceType,
					Hostname: i.HostName,
					System:   i.OSName,
					Intranet: i.NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress,
					Status:   i.Status,
					Public:    i.EipAddress.IpAddress,
					CreatedAt: ct,
					Cpu:       strconv.Itoa(i.Cpu),
					Memory:    strconv.Itoa(i.Memory / 1024),
					Uptime:    &st,
					Region:    v.Region,
				}
				_ = db.Create(&asset).Error
			}
		}

	}
	return newInstanceList, nil
}

func GetAliyunInstancesDisk() {


	var keys []autocode.Key
	_ = db.Find(&keys).Error

	for _, v := range keys {

		var asset []autocode.Asset
		db.Where("deleted_at is  null").Where("region = ?", v).Find(&asset)
		for _, assetSn := range asset {
			client, _ := ecs.NewClientWithAccessKey(v.Region, v.Keyid, v.Secret)
			request := ecs.CreateDescribeDisksRequest()
			request.Scheme = "https"
			request.InstanceId = assetSn.Sn
			response, _ := client.DescribeDisks(request)

			var s []string
			for _, b := range response.Disks.Disk {
				s = append(s, fmt.Sprintf("%s %v G", b.Device, b.Size))
			}
			var assetDisk autocode.Asset
			db.Where("sn = ?", assetSn.Sn).Find(&assetDisk).Update("disk", strings.Join(s, " "))
		}
	}

}

func AssetHostUpdate() {

	var err error
	var newInstanceList []string
	newInstanceList, err = GetAliyunInstances(newInstanceList)
	if err != nil {
		fmt.Println("Got an error retrieving information about your  instances:")
		return
	}
	var oldInstanceList []string
	var oldAsset []autocode.Asset
	err = db.Where("brand = ?", "阿里云").Where("deleted_at is  null").Unscoped().Find(&oldAsset).Error
	for _, v := range oldAsset {
		oldInstanceList = append(oldInstanceList, v.Sn)
	}
	fm := make(map[string]int)
	for i, v := range newInstanceList {
		fm[v] = i
	}

	for _, v := range oldInstanceList {
		if _, ok := fm[v]; ok {
		} else {
			asset := autocode.Asset{
				Sn:        v,
				DeletedAt: gorm2.DeletedAt(sql.NullTime{time.Now(), true}),
			}
			db.Model(&autocode.Asset{}).Where("sn = ?", v).Updates(asset)
		}
	}
	GetAliyunInstancesDisk()
}


func AssetUpdate() {
	fmt.Println("资产更新")
	mysqlCfg := global.GVA_CONFIG.Mysql
	if name, _ := os.Hostname(); name == "raptor" {
		mysqlCfg = global.GVA_CONFIG.MysqlProd
		fmt.Println("线上环境", mysqlCfg.Path)
	} else {
		fmt.Println("测试环境", mysqlCfg.Path)
	}
	db, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Path,
		mysqlCfg.Dbname, mysqlCfg.Config))

	AssetHostUpdate()

	defer db.Close()

}
