package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(date string) string {
	h := md5.New()
	h.Write([]byte(date))
	tt := h.Sum(nil)
	return hex.EncodeToString(tt)
}

func MD5Encode(date string) string {
	return strings.ToUpper(Md5Encode(date))
}

func MakePassword(plainPwd, salt string) string  {
    return Md5Encode(plainPwd + salt)
}

func ValidatePassword(plainPwd, salt, encodedPwd string) bool {
    return Md5Encode(plainPwd + salt) == encodedPwd
}
