package conf

import "time"

//完整时间格式
const SystimeFormat = "2006-01-02 15:04:05"

//短时间格式
const SystimeFormatShort = "2006-01-02"

//每日限制ip抽奖次数
const LimitIpMaxNum = 1000

//每日单个ip的抽奖次数
const IpPrizeMax = 500

//每日单个用户的抽奖次数
const UserPrizeMax = 500

//设置时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

//签名密钥
var SignSecret = []byte("qwerqwer123")

//cookie中的加密验证密钥
var CookieSecret = "asdfasdf123"

const GtypeVirtual = 0   // 虚拟币
const GtypeCodeSame = 1  // 虚拟券，相同的码
const GtypeCodeDiff = 2  // 虚拟券，不同的码
const GtypeGiftSmall = 3 // 实物小奖
const GtypeGiftLarge = 4 // 实物大奖
