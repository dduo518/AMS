package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func FindValIsHasInMap(val string, key string, list []map[string]interface{}) bool {
	var isHas bool = false
	for _, item := range list {
		if _, ok := item[key]; ok {
			if val == item[key] {
				isHas = true
			}
		}
	}
	return isHas
}

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func CreateToken(userId,appid  string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["appid"] = appid
	claims["userId"] = userId
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token.Claims = claims
	accessToken,_ := token.SignedString([]byte(userId))
	return accessToken
}
func VerifyToken(accessToken ,secret string) (bool,error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(secret), nil
	})
	if err!=nil{
		log.Println(err)
		return false,err
	}
	return token.Valid,nil
}

