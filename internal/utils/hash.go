package utils
import (
	"crypto/sha1"
	"encoding/hex"
)
func HashPassword(password string ) string {
	hasher:=sha1.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))  // what is 
}