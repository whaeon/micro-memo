package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Profile struct {
	path     string
	name     string
	confType string
	v        *viper.Viper
}

func NewProfile(path string, name string, confType string) *Profile {
	return &Profile{
		path:     path,
		name:     name,
		confType: confType,
		v:        viper.New(),
	}
}

func (p *Profile) GetValue(key string) string {
	p.v.AddConfigPath(p.path)
	p.v.SetConfigName(p.name)
	p.v.SetConfigType(p.confType)
	err := p.v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found!")
		} else {
			fmt.Println("fatal error: ", err)
		}
	}
	return p.v.GetString(key)
}
