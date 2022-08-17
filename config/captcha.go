package config

import "gohub/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码图片高度
			"height": 80,

			// 验证码图片宽度
			"width": 240,

			// 验证码的长度
			"length": 6,

			// 数字的最大倾斜角度
			"maxskew": 0.7,

			// 图片背景里的混淆点数量
			"dotcount": 80,

			// 过期时间，单位是分钟
			"expire_time": 15,

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": 10080,

			// 非 production 环境，使用此 key 可跳过验证，方便测试
			"testing_key": "captcha_skip_test",
		}
	})
}
