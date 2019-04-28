package conf

import "time"

//完整时间格式
const SystimeFormat = "2006-01-02 15:04:05"

//短时间格式
const SystimeFormatShort = "2006-01-02"

//设置时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

//签名密钥
var SignSecret = []byte("qwerqwer123")

//cookie中的加密验证密钥
var CookieSecret = "asdfasdf123"
