package util

import (
	"github.com/dgrijalva/jwt-go"
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

func CreateToken(userId,appid ,secret string) string {
	mySigningKey := []byte(secret)
	claims:=jwt.MapClaims{
		"userid":userId,//签发人 appid
		"iss":  appid,//签发人 appid
		"iat":  time.Now().Unix(),//签发时间
		"nbf": time.Now().Unix(),//生效时间
		"exp": 60*60,// 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken,_ := token.SignedString(mySigningKey)
	return accessToken
}
func VerifyToken(userId,appid ,secret string) string {
	mySigningKey := []byte(secret)
	claims:=jwt.MapClaims{
		"userid":userId,//签发人 appid
		"iss":  appid,//签发人 appid
		"iat":  time.Now().Unix(),//签发时间
		"nbf": time.Now().Unix(),//生效时间
		"exp": 60*60,// 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken,_ := token.SignedString(mySigningKey)
	return accessToken
}

func Sum(numbers []int) int {
	sum := 0
	// 这个 bug 是故意的
	for _, n := range numbers {
		sum += n
	}
	return sum
}