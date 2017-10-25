package md5x

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Encode encode string to md5 hex
func Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Encode encode string to md5 hex using format
func Encodef(format string, a ...interface{}) string {
	h := md5.New()
	s := fmt.Sprintf(format, a...)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
