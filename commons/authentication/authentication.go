package authentication

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

/***********************************************************************/
/* 関数名：Generate_JWT
/* 引数　：Usercd usermstから取得したユーザCD
/* 戻り値：tokenString　生成したJWT
/*        error JWT生成時、エラーが発生した場合にその内容を返す。JWTは返さない。
/* 内容　：取得したユーザ情報から、JWTを生成する。
/* 作成者：Norifumi Sato
/* 作成日：2023/6/22
/* 備考　：
/***********************************************************************/
func Generate_JWT(Usercd string) (string, error) {

	if Usercd == "" {
		return "", errors.New("usercd is empty")
	}
	// Tokenの生成
	token := jwt.New(jwt.SigningMethodHS256)

	// Tokenに対して、情報を追加する。
	claims := token.Claims.(jwt.MapClaims)
	claims["usercd"] = Usercd
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	secret := []byte("secret")

	// 署名
	tokenString, error := token.SignedString(secret)

	if error != nil {
		return "", error
	}

	return tokenString, nil
}
