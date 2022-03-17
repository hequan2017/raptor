package config

type Ding struct {
	Key    string `mapstructure:"appkey" json:"appkey" yaml:"appkey"`
	Secret string `mapstructure:"appsecret" json:"appsecret" yaml:"appsecret"`
}
