package system

import (
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

const (
	// DEFAULT_CONFIG_VERSION 默认版本
	DEFAULT_CONFIG_VERSION = "v1"
)

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Version: DEFAULT_CONFIG_VERSION,
		Email:   mail.NewDefaultConfig(),
		SMS:     sms.NewDefaultConfig(),
	}
}

// Config 系统配置
type Config struct {
	Version string       `bson:"_id" json:"version"`
	Email   *mail.Config `bson:"email" json:"email"`
	SMS     *sms.Config  `bson:"sms" json:"sms"`
}

// Desensitize 脱敏
func (c *Config) Desensitize() {
	c.Email.Desensitize()
	c.SMS.Desensitize()
}
