package xml

import (
	conf "github.com/mabetle/mcore/mconf"
	"github.com/mabetle/mlog"
)

var(
	logger = mlog.GetLoggerWithCatalog("github.com/mabetle/mcore/mconf/xml")
)

type XmlConfig struct{
}

func GetConfig(file string) conf.Config {
	logger.Debug("build Config, using xml")
	return GetXmlConfig(file)
}

func GetXmlConfig(file string)*XmlConfig{
	return &XmlConfig{}
}

// implements for Config
func (c XmlConfig) GetString(key string)string{
	return ""
}

func (c XmlConfig)GetInt(key string)int{
	return 0
}

func (c XmlConfig)GetBool(key string)bool{
	return false
}

func (c XmlConfig)IsContain(key string)bool{
	return false
}


