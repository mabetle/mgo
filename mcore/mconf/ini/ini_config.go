package ini

import (
	"fmt"
	"github.com/robfig/config"
	"github.com/mabetle/mgo/mcore"
)

// IniConfig
// use robfig config as default
// IniConfig implements KeyValueLoader mconf.Config interface.
type IniConfig struct {
	Locations []string
	*config.Config
}

// NewIniConfig supports more than one location.
// returns IniConfig, implements mcore.Config
func NewIniConfig(locations ...string) *IniConfig {
	c := &IniConfig{Locations: locations}
	conf := config.NewDefault()
	for _, location := range locations {
		confItem, err := config.ReadDefault(location)
		if logger.CheckError(err) {
			continue
		}
		conf.Merge(confItem)
	}
	c.Config = conf
	return c
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
