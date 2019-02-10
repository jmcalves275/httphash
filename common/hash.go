package common

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(text string) string {
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
