package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/**
 * md5Tools Config
 */
type Md5ToolsConfig struct {
	SaltModule    SaltModule `json:"saltModule"`
	Title         []string   `json:"title"`
	OutSheetName  string     `json:"outSheetName"`
	InSheetName   string     `json:"inSheetName"`
}

/**
 * 是否加盐值模式  on:开启加盐   off:不开启加盐
 */
type SaltModule struct {
	Activate   string      `json:"activate"`
	SaltKey    string       `json:"saltKey"`
}

/**
 * 读取 json 配置文件
 */
func LoadConfig(configPath string) *Md5ToolsConfig {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}

	md5ToolsConfig := &Md5ToolsConfig{}
	err = json.Unmarshal(buf, md5ToolsConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return md5ToolsConfig
}