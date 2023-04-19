package config

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
)

var VersionInfo *versionInfo

// ReadConfig 每天更新一次配置表
func ReadConfig() {
	file, _ := os.Open("./version_info_config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&VersionInfo)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func InitCron() {
	cron2 := cron.New() //创建一个cron实例
	//执行定时任务（一个半小时更新一次）
	spec := "@every 1h30m"
	_, err := cron2.AddFunc(spec, ReadConfig)
	if err != nil {
		fmt.Println(err)
	}
	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

type versionInfo struct {
	VersionName string `json:"version_name"`
	VersionCode int    `json:"version_code"`
	ForceUpdate bool   `json:"force_update"`
	UpdateDesc  string `json:"update_desc"`
	ApkUrl      string `json:"apk_url"`
	//强制更新最低版本号
	ForceUpdateMinVersion string `json:"force_update_min_version"`
	//强制更新最低code
	ForceUpdateMinCode int `json:"force_update_min_code"`
}
