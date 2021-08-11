package config

import (
	"github.com/spf13/viper"
)

type Interface interface {
	Init()
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	AllSettings() map[string]interface{}
	IsSet(key string) bool
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
	GetStringSlice(key string) []string
	UnmarshalKey(key string, rawVal interface{}, decoder ...viper.DecoderConfigOption) error

	BackendConfig() (map[string]string, error)
}
