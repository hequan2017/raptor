package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"raptor/server/model/system"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"raptor/server/global"
	model "raptor/server/model/system"

	"strconv"
	"time"
)

var db1 *gorm.DB

func GetToken(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	response, _ := client.Do(resp)
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func PostToken(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	return string(result)
}


// 部门更新
func Department(id int, token string) {
	departmentList := PostToken(fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=%s", token),
		map[string]int{
			"dept_id": id,
		},
		"application/json",
	)
	data := gofasion.NewFasion(departmentList)
	result := make(map[int]string)
	for _, v := range data.Get("result").Array() {
		deptId := gjson.Get(v.Json(), "dept_id").Int()
		result[int(deptId)] = gjson.Get(v.Json(), "name").String()
	}
	var menu []model.SysBaseMenu
	db1.Where("id IN (?)", []int{1, 8}).Find(&menu)
	mysqlCfg := global.GVA_CONFIG.Mysql
	if name, _ := os.Hostname(); name == "raptor" {
		mysqlCfg = global.GVA_CONFIG.MysqlProd
	}
	a, _ := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?", mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Path,
		mysqlCfg.Dbname))
	Enforcer, _ := casbin.NewEnforcer("./resource/rbac_model.conf", a)

	for i, v := range result {
		var authority model.SysAuthority
		idCasbin := strconv.Itoa(i)
		_ = db1.Where("authority_id = ?", i).Where("deleted_at is null").Unscoped().First(&authority)
		if authority.AuthorityId == "" {
			authorityId := model.SysAuthority{
				AuthorityId:   idCasbin,
				AuthorityName: v,
				ParentId:      strconv.Itoa(id),
				UpdatedAt:     time.Now(),
				SysBaseMenus:  menu,
			}
			authorityId.CreatedAt = time.Now()
			db1.Model(&authority).Create(&authorityId)
		} else {
			authorityId := model.SysAuthority{
				AuthorityId:   idCasbin,
				AuthorityName: v,
				ParentId:      strconv.Itoa(id),
				UpdatedAt:     time.Now(),
			}
			db1.Model(&authority).Where("authority_id = ?", i).Where("deleted_at is null").Updates(&authorityId)
		}
		Enforcer.DeletePermission(idCasbin, "/base/login", "POST")
		Enforcer.DeletePermission(idCasbin, "/user/changePassword", "POST")
		Enforcer.DeletePermission(idCasbin, "/user/setUserInfo", "PUT")
		Enforcer.DeletePermission(idCasbin, "/menu/getMenu", "POST")
		Enforcer.DeletePermission(idCasbin, "/menu/getMenuList", "POST")
		Enforcer.DeletePermission(idCasbin, "/menu/getBaseMenuTree", "POST")
		Enforcer.DeletePermission(idCasbin, "/menu/getBaseMenuById", "POST")
		Enforcer.DeletePermission(idCasbin, "/jwt/jsonInBlacklist", "POST")
		Enforcer.DeletePermission(idCasbin, "/user/getUserList", "POST")
		Enforcer.DeletePermission(idCasbin, "/fileUploadAndDownload/getFileList", "POST")
		Enforcer.AddPolicy(idCasbin, "/base/login", "POST")
		Enforcer.AddPolicy(idCasbin, "/user/changePassword", "POST")
		Enforcer.AddPolicy(idCasbin, "/user/setUserInfo", "PUT")
		Enforcer.AddPolicy(idCasbin, "/menu/getMenu", "POST")
		Enforcer.AddPolicy(idCasbin, "/menu/getMenuList", "POST")
		Enforcer.AddPolicy(idCasbin, "/menu/getBaseMenuTree", "POST")
		Enforcer.AddPolicy(idCasbin, "/user/getUserList", "POST")
		Enforcer.AddPolicy(idCasbin, "/menu/getBaseMenuById", "POST")
		Enforcer.AddPolicy(idCasbin, "/jwt/jsonInBlacklist", "POST")
		Enforcer.AddPolicy(idCasbin, "/fileUploadAndDownload/getFileList", "POST")
		Department(i, token)
	}
	return
}

// 用户更新
func User(dept_id int, token string) {
	user := PostToken(fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/list?access_token=%s", token),
		map[string]int{
			"dept_id": dept_id,
			"cursor":  0,
			"size":    100,
		},
		"application/json",
	)
	data := gofasion.NewFasion(user)
	fmt.Println(data)
	for _, v := range data.Get("result").Get("list").Array() {
		var sysUser model.SysUser
		_ = db1.Where("username = ?", gjson.Get(v.Json(), "name").String()).Where("deleted_at is null").Unscoped().First(&sysUser)
		var sysAuthority []model.SysAuthority
		_ = db1.Where("authority_id = ?", strconv.Itoa(dept_id)).Where("deleted_at is null").Unscoped().Find(&sysAuthority)

		if sysUser.Username == "" {
			uid := uuid.NewV4()
			u := model.SysUser{
				Username:      gjson.Get(v.Json(), "name").String(),
				Password:      "e10adc3949ba59abbe56e057f20f883e",
				NickName:      gjson.Get(v.Json(), "name").String(),
				HeaderImg:     gjson.Get(v.Json(), "avatar").String(),
				Unionid:       gjson.Get(v.Json(), "unionid").String(),
				Authorities: sysAuthority,
				UUID:          uid,
			}
			db1.Create(&u)
		} else {
			u := model.SysUser{
				Unionid:       gjson.Get(v.Json(), "unionid").String(),
				NickName:      gjson.Get(v.Json(), "name").String(),
				HeaderImg:     gjson.Get(v.Json(), "avatar").String(),
				Authorities: sysAuthority,
			}
			db1.Where("username = ?", gjson.Get(v.Json(), "name").String()).Where("deleted_at is null").Model(&sysUser).Updates(&u)
		}
	}

	return
}

func DingUpdate() {
	fmt.Println("钉钉用户更新")
	mysqlCfg := global.GVA_CONFIG.Mysql
	if name, _ := os.Hostname(); name == "raptor" {
		mysqlCfg = global.GVA_CONFIG.MysqlProd
		fmt.Println("线上环境", mysqlCfg.Path)
	} else {
		fmt.Println("测试环境", mysqlCfg.Path)
	}
	db1, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Path,
		mysqlCfg.Dbname, mysqlCfg.Config))
	dingCfg := global.GVA_CONFIG.Ding
	getToken := GetToken(fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", dingCfg.Key, dingCfg.Secret))

	token := gjson.Get(getToken, "access_token").String()

	var authority system.SysAuthority
	db := global.GVA_DB.Model(&system.SysAuthority{})
	_ = db.Where("authority_id = ?", "1").First(&authority).Error

	if authority.AuthorityId != "1" {
		var authorityOne model.SysAuthority
		authorityId := model.SysAuthority{
			AuthorityId:   "1",
			AuthorityName: "公司",
			ParentId: "0",
		}
		db1.Model(&authorityOne).Create(&authorityId)
	}

	Department(1, token)

	var authorityS []model.SysAuthority
	_ = db1.Where("deleted_at is null").Find(&authorityS)
	for i, v := range authorityS {
		if i >= 0 {
			id, _ := strconv.Atoi(v.AuthorityId)
			User(id, token)
		}
	}

}
