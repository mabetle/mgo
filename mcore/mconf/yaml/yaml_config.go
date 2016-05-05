package yaml

import (
	conf "github.com/mabetle/mgo/mcore/mconf"
	"github.com/mabetle/mgo/mlog"
)

var(
	logger = mlog.GetLoggerWithCatalog("github.com/mabetle/mgo/mcore/mconf/yaml")
)

type YamlConfig struct{
}


func GetConfig(file string) conf.Config {
	logger.Debug("build Config, using yaml")
	return GetYamlConfig(file)
}

func GetYamlConfig(file string)*YamlConfig{
	return &YamlConfig{}
}

// implements for Config

func (c YamlConfig) GetString(key string)string{
	return ""
}

func (c YamlConfig)GetInt(key string)int{
	return 0
}

func (c YamlConfig)GetBool(key string)bool{
	return false
}

func (c YamlConfig)IsContain(key string)bool{
	return false
}

