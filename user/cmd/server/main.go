package main

import (
	"fmt"
	"user/conf"
)

func main() {
	path := "conf/"
	name := "config"
	confType := "yaml"
	profile := conf.NewProfile(path, name, confType)
	testVal := profile.GetValue("mysql.container.name")
	fmt.Println(testVal)
}
