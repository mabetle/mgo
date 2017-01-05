package ini

import (
	"fmt"

	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mcore/mconf"
	"github.com/robfig/config"
)

// IniConfig
// use robfig config as default
// IniConfig implements KeyValueLoader mconf.Config interface.
type IniConfig struct {
	// extends config.Config
	*config.Config
	Locations []string
}

// NewIniConfig supports more than one location.
// returns IniConfig, implements mcore.Config
func NewIniConfig(locations ...string) *IniConfig {
	c := &IniConfig{Locations: locations}
	conf := config.NewDefault()
	for _, location := range locations {
		confItem, err := config.ReadDefault(location)
		// read config error
		if err != nil {
			continue
		}
		conf.Merge(confItem)
	}
	c.Config = conf
	return c
}

// NewConfig make sure IniConfig implements mcore.Config
func NewConfig(locations ...string) mconf.Config {
	iniConf := NewIniConfig(locations...)
	return mconf.NewConfig(iniConf)
}

func (c IniConfig) LoadKeyValue() mcore.StringKeyValueMap {
	logger.Debugf("load config from %v.", c.Locations)
	skv := mcore.NewStringKeyValueMap()
	for _, section := range c.Sections() {
		options, _ := c.Config.Options(section)
		for _, option := range options {
			k := ""
			if section == "DEFAULT" {
				k = fmt.Sprintf("%s", option)
			} else {
				k = fmt.Sprintf("%s@%s", option, section)
			}
			v, _ := c.Config.String(section, option)
			logger.Tracef("load key:%s value:%s", k, v)
			skv.Put(k, v)
		}
	}
	return skv
}
