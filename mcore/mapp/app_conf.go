package mapp

import (
	"fmt"
	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mcore/mconf/ini"
	"os"
)

const (
	ConfDir  = "/rundata/mapps"
	ConfFile = "app.conf"
)

type AppConf struct {
	confDir     string
	confFile    string
	AppName     string
	runMode     string
	performLoad bool
	mcore.StringKeyValueMap
}

func NewAppConf(dir, file, appName string) *AppConf {
	t := &AppConf{confDir: dir, confFile: file, AppName: appName}
	//t.runMode = RunMode
	t.performLoad = true
	return t
}

func NewDefaultAppConf(appName string) *AppConf {
	return NewAppConf(ConfDir, ConfFile, appName)
}

func (c *AppConf) SetConfDir(dir string) *AppConf {
	c.confDir = dir
	return c
}

func (c *AppConf) Init() {
	if c.IsExist() {
		return
	}
	sb := mcore.NewStringBuffer()
	sb.AppendLine("# Auto Generate")
	sb.AppendLine(fmt.Sprintf("%s=%s", KEY_VENDOR_NAME, V_VENDOR_MABETLE))
	logger.Tracef("Init config. AppName:%s Location:%s", c.AppName, c.Location())
	_, err := mcore.WriteFile(c.Location(), sb.String())
	if err != nil {
		logger.Error(err)
	}
}

func (c *AppConf) Location() string {
	return fmt.Sprintf("%s/%s/%s", c.confDir, c.AppName, c.confFile)
}

func (c *AppConf) IsExist() bool {
	return mcore.IsFileExist(c.Location())
}

func (c *AppConf) Load() *AppConf {
	if c.IsExist() {
		logger.Tracef("Load config. AppName:%s Location:%s", c.AppName, c.Location())
		c.StringKeyValueMap = ini.NewIniConfig(c.Location()).LoadKeyValue()
		return c
	}
	logger.Tracef("Location not exists, Create Default StringKeyValueMap. AppName:%s", c.AppName)
	c.StringKeyValueMap = mcore.NewStringKeyValueMap()
	c.performLoad = false
	return c
}

func (c *AppConf) GetString(key string) string {
	if c.performLoad {
		logger.Tracef("Need reload config. AppName:%s", c.AppName)
		c.Load()
	}
	return c.StringKeyValueMap.GetString(key)
}

func (c *AppConf) GetVendorName() string {
	return c.GetString(KEY_VENDOR_NAME)
}

func (c *AppConf) RunMode() string {
	if c.runMode != "" {
		return c.runMode
	}

	//1. check flag
	if mcore.IsHasDevArg(){
		c.runMode = MODE_DEV
	}else if mcore.IsHasProdArg(){
		c.runMode = "prod"
	}else if mcore.IsHasTestArg(){
		c.runMode="test"
	}

	if c.runMode !=""{
		return c.runMode
	}

	// 2.check env
	if v := os.Getenv("RUN_MODE");v!=""{
		c.runMode = v
		return c.runMode
	}

	// after all try, give default RunMode.
	if c.runMode == "" {
		c.runMode = MODE_DEV
	}
	return c.runMode
}
