package utils

import (
	"encoding/hex"
	"ops/pkg/model/jwt"
	"ops/service/userRegister/data"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"

	"strings"
	"time"
)

func PbkToAddr(pbk string) string {
	pk, _ := hex.DecodeString(pbk)                                       //这个是将公钥转成16进制的一个   byte[]
	hash := sha3.NewLegacyKeccak256()                                    //然后先new 一个sha3 256的变量
	hash.Write(pk[1:])                                                   //用这个变量把刚才生成的地址，截取从第1位开始到最后的
	result := strings.ToLower("0x" + hexutil.Encode(hash.Sum(nil))[26:]) //
	return result                                                        //最终得到公钥地址
}

func GeneratedToken(login *data.Login) (*data.LoginResult, string) {
	j := &jwt.JWT{
		SigningKey: []byte("opsnft_jwt"),
	}
	claims := jwt.CustomClaims{
		login.Uid,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			//ExpiresAt: int64(time.Now().Unix() + 3600*24*10), // 过期时间10天
			ExpiresAt: int64(time.Now().Unix() + 3600*24*62), // 过期时间10天
			Issuer:    "opsnft_jwt",                          //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return &data.LoginResult{}, RECODE_GENERATETOKENERR
	}

	data := &data.LoginResult{
		Token: token,
	}
	return data, RECODE_OK
}
