### JSW
```
session加密，将其放到客户端。本地只保留密钥，验证只需要找到对应的密钥，再解密session获得用户信息


import (
    "github.com/dgrijalva/jwt-go"
)

func CreateToken() string{
    token := jwt.New(jwt.SigningMethodHS256)
    claims := jwt.MapClaims{
        id : "id",
        name: "name",
        iat: "签发时间"
        //iss: 签发者
        //sub: 面向的用户
        //aud: 接收方
        //exp: 过期时间
        //nbf: 生效时间
        //jti: 唯一身份标识
    }
    token.Claims = claims
    tokenString, _ := token.SignedString([]byte(secret))
    return tokenString
}

func CheckToken(tokenString string, secret string) bool{
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error){
        return []byte(secret),nil
    }) 

    if err != nil || !token.Valid {
        return false
    }

    claims, err := token.Claims.(jwt.MapClaims)
    //claims['id'].(int)
    //cliams['name'].(string)
    return true

}
```