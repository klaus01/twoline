package test

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	_ "github.com/klaus01/twoline/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	currpath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	apppath := filepath.Dir(currpath)
	beego.TestBeegoInit(apppath)

	// beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("testsessionproviderconfig")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("testmysqlconn"))

	// 创建测试数据库
	o := orm.NewOrm()
	dbName := beego.AppConfig.String("testmysqldatabase")
	if _, err := o.Raw("DROP DATABASE IF EXISTS " + dbName).Exec(); err != nil {
		beego.Error(err)
		return
	}
	if _, err := o.Raw("CREATE DATABASE " + dbName).Exec(); err != nil {
		beego.Error(err)
		return
	}
	if _, err := o.Raw("USE " + dbName).Exec(); err != nil {
		beego.Error(err)
		return
	}
	sqlFilePath := filepath.Join(currpath, "twoline.sql")
	fileByte, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		beego.Error(err)
		return
	}
	fileString := string(fileByte)
	index := strings.Index(fileString, ";")
	for index >= 0 {
		sql := fileString[:index]
		if _, err := o.Raw(sql).Exec(); err != nil {
			beego.Error(err)
			return
		}
		fileString = fileString[index+1:]
		index = strings.Index(fileString, ";")
	}
	if len(strings.TrimSpace(fileString)) > 0 {
		if _, err := o.Raw(fileString).Exec(); err != nil {
			beego.Error(err)
			return
		}
	}
}
