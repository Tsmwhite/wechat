package verify_code

import (
	"math/rand"
	"time"
)

func createCode(length int) (code string) {
	var str []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	max := len(str)
	var key int
	for i := 0; i < length; i++ {
		randS := rand.New(rand.NewSource(time.Now().UnixNano()))
		key = randS.Intn(max)
		code += str[key]
	}
	return code
}

func RegisterCode(email string) {

}
