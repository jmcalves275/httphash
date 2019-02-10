package common

import (
	"crypto/md5"
	"encoding/hex"
)

// Function to make an md5 hash of some text
func MD5Hash(text string) string {
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
