package conf

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

// Config 配置
var Config = struct {
	PushAddress map[string]interface{}
	Port        string
	Branch      string
}{}

var k = koanf.New(".")

func init() {
	if err := k.Load(file.Provider("conf.json"), json.Parser()); err != nil {
		panic(err)
	}
	Config.PushAddress = k.Get("puhAddress").(map[string]interface{})
	Config.Port = k.String("port")
	Config.Branch = k.String("branch")
}
