package test

import (
	"fmt"
	"ops/pkg/utils"
	"ops/service/userRegister/data"
	"testing"
)

//var walletPk = "0x1F102c275b0eE08187d0b8A3bB3FdE97076089b5"
//var randomCode = "eeAJUL4UFBM6P3cBCQQkefECoxD8RqS6Caz6XFNwyrlMAZdYowZAEvreeWJx8fZo1zwFRmy7mDxcnmehy166zFCucngRA0LjV608Uxc8eOzJCWTUZ8mAv8ldlrUqfkV4"

func TestHash(t *testing.T) {
	t.Run("pipeline", func(t *testing.T) {
		pbk := "04514ba25eff2d513d0619bfd4fe41297c64be25186e417ae5de2699a85b7ea1ef037a56cb479514e6e5b67bb47db140a6854240177f2c42c5fc7cffc5bf198041" //传过来的公钥,转成地址值
		//使用公钥hash扫描redid中是否存在这个随机码
		//message := "HrWVmeCAbOAMPDdOi36NYE4RnxaHnLeo0PnGqvXf9n2ONjub7BOteFXQqOIkzqwubaAE7rsn6qD36eoH8De7VK4oi8IUQgtsxzzNuCVMJLRyRDTPeiD4aBIftE79vHGk"
		//sign := "f5a80cacff9512ba58cbfe087864b140cc50dd9f59869b9fbbbd19b6814157fd53b6b779965a5eb5ffa883041c75ecbf4ef443a8f971ed1564996264761077021b"

		fmt.Println(utils.PbkToAddr(pbk))

		//isVerify := verifyLogin(pbk, message, sign)

		fmt.Println(isVerify)
	})
}

func TestToken(t *testing.T) {
	t.Run("getToken", func(t *testing.T) {
		var userId string
		var userLogin *data.Login
		userId = "TOv3z41RCsloMh01xYCi"
		userLogin = &data.Login{
			Uid:      userId,
			Pubkaddr: "0x4b430a57bd51e0dc6367c8b4641716388dbc2918",
		}
		logResult, _ := utils.GeneratedToken(userLogin)
		if errCode == utils.RECODE_GENERATETOKENERR {
			logger.Error(errCode)
		}
		fmt.Println(logResult)
	})
}
