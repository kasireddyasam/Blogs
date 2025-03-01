package utils
import (
	"crypto/sha256"
	"crypto/rand"
	"encoding/hex"
	"log";"os";
	"github.com/joho/godotenv"
)

func LoadEnv(){
	err:=godotenv.Load()
	if err!=nil {
		log.Fatal("Error loading in .env file")
	}
}

// generate salt string in random
func GenerateSalt() string{
	salt:=make([]byte,16) // 16 bytes of salt
	_,err:=rand.Read(salt)
	if err!=nil {
		log.Fatal("Failed to prepare salt")
	}
	return hex.EncodeToString(salt)

}
func HashPassword(password ,salt string ) string {
	LoadEnv()
	pepper:=os.Getenv("PIPPER_KEY")
	if pepper=="" {
		log.Fatal("PIPPER_KEY not set in .env file")
	}
	NewPassword:=password+salt+pepper
	hasher:=sha256.New()
	hasher.Write([]byte(NewPassword))
	return hex.EncodeToString(hasher.Sum(nil))  // what is 
}