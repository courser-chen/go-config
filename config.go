package cfg

import (
	"errors"
	"github.com/courser-chen/go-config/utils"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"strings"
)

/*
* @Author: jensen.chen
* @Date:   2025/4/28 14:17
 */
type Config struct {
	Data map[string]interface{}
}

var CliConfig Config = Config{}

const KEY_DELIMITER string = "."

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.cron\
 * 如果找不到配置,则返回指定的默认值
 */
func GetConfigValue(key string, defaultValue interface{}) (interface{}, error) {
	val, err := GetConfig(key)
	if err != nil {
		return defaultValue, nil
	}
	if val != nil {
		return val, nil
	} else {
		return defaultValue, nil
	}
}

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.cron
 */
func GetConfig(key string) (interface{}, error) {
	keyAry := strings.Split(key, KEY_DELIMITER)
	var val interface{} = CliConfig.Data
	for _, itemKey := range keyAry {
		val = getMapValue(val.(map[string]interface{}), itemKey)
		if val == nil {
			break
		}

	}
	if val != nil {
		return val, nil
	} else {
		return nil, errors.New("Config Not Exists")
	}
}

func GetKeys(data map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(data))
	for k := range data {
		keys[j] = k
		j++
	}
	return keys
}

func getMapValue(data map[string]interface{}, key string) interface{} {
	rs, _ := data[key]
	return rs
}

/*
*
配置项自动写入struct变量
@key 配置项的ROOT键值
@target Struct指针
*/
func Set(key string, target interface{}) {
	config, _ := GetConfig(key)
	if config != nil && reflect.ValueOf(config).Kind() == reflect.Map {
		data := config.(map[string]interface{})
		utils.MapToStruct(data, target, "config")
	}
}

func Load(path string) {
	log.Println("Read Config", path)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	decode := yaml.NewDecoder(f)
	decode.Decode(&(CliConfig.Data))
	log.Println(CliConfig.Data)
}
