package config

import "github.com/spf13/viper"

type Config interface {
	Init()
	SetString(key string, value string)
	GetString(key string) string
	SetInt(key string, value int)
	GetInt(key string) int
	SetBool(key string, value bool)
	GetBool(key string) bool
}

type viperConfig struct {
}

func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}

func (v *viperConfig) Init() {

	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) SetString(key string, value string) {
	viper.Set(key, value)
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) SetInt(key string, value int) {
	viper.Set(key, value)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) SetBool(key string, value bool) {
	viper.Set(key, value)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}
