package config

import "gohub/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码的长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),

			// 过期时间，单位是分钟
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": 10080,
			// 本地开发环境验证码使用 debug_code
			"debug_code": 123456,

			// 方便本地和 API 自动测试
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
