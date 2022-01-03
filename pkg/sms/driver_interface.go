package sms

type Driver interface {
	// 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
