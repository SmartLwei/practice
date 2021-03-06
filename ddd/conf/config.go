package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Setting config info for whole project
type Setting struct {
	Rest *RestSetting `json:"rest" yaml:"rest"`
	Grpc *GrpcSetting `json:"grpc" yaml:"grpc"`
	DB   *DBSetting   `json:"db" yaml:"db"`
}

// RestSetting rest server config
type RestSetting struct {
	Port string `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

func (rs *RestSetting) String() string {
	return fmt.Sprintf("{port: %s, mode: %s}", rs.Port, rs.Mode)
}

// GrpcSetting grpc server config
type GrpcSetting struct {
	Port string `json:"port" yaml:"port"`
}

func (gs *GrpcSetting) String() string {
	return fmt.Sprintf("{port: %s}", gs.Port)
}

// DBSetting database config
type DBSetting struct {
	URI          string `json:"uri" yaml:"uri"`
	LogLevel     int    `json:"log_level" yaml:"log_level"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
}

func (dbs *DBSetting) String() string {
	return fmt.Sprintf("{uri: %s, log_level:%d, max_idle_conns:%d, max_open_conns:%d}", dbs.URI, dbs.LogLevel,
		dbs.MaxIdleConns, dbs.MaxOpenConns)
}

// BuildConfig build config from file
func BuildConfig(fileName string) *Setting {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("read file " + fileName + " failed")
	}
	var cf = Setting{}
	if err := yaml.Unmarshal(content, &cf); err != nil {
		panic("unmarshal setting failed:" + err.Error())
	}
	return &cf
}

//
//func Init(fileName string) {
//	once.Do(func() {
//		content, err := os.ReadFile(fileName)
//		if err != nil {
//			panic("read file " + fileName + " failed")
//		}
//		var cf = new(Setting)
//		if err := yaml.Unmarshal(content, cf); err != nil {
//			panic("unmarshal setting failed:" + err.Error())
//		}
//		setting = cf
//	})
//}
//
//func GetSetting() *Setting {
//	if setting == nil {
//		panic("get setting failed, please init setting first")
//	}
//	return setting
//}
