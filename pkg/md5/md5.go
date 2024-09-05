package md5pkg

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// CompareMD5 compare md5
func CompareMD5(value, md5Value string) bool {
	return EncodeMD5(value) == md5Value
}
