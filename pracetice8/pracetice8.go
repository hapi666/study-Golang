package main

import (
	"fmt"
	//"github.com/satori/go.uuid"
	"time"
)

func main() {
	var er error
	er = fmt.Errorf("fail")
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	//if time.Now().Second().(string)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("18:04:01"))
	//tuuid, err := uuid.NewV4()
	// token := tuuid.String()
	// if err != nil {
	// 	fmt.Println("fail")
	// }
	// Token, err := uuid.FromString(token)
	// TrueToken := Token.String()
	// fmt.Println(token)
	// fmt.Println(Token)
	// fmt.Println(TrueToken)
	// fmt.Println([]byte("123abc"))
}

// aes.go
// package main

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"encoding/hex"
// 	"errors"
// 	"fmt"
// 	"io"
// )

// func main() {
// 	se, err := AesEncrypt("aes-20170416-30-1000")
// 	fmt.Println(se, err)
// 	sd, err := AesDecrypt(se)
// 	fmt.Println(sd, err)
// }

// var (
// 	commonkey = []byte("nanjishidu170416")
// 	//syncMutex sync.Mutex
// )

// // func SetAesKey(key string) {
// // 	syncMutex.Lock()
// // 	defer syncMutex.Unlock()
// // 	commonkey = []byte(key)
// // }
// func AesEncrypt(plaintext string) (string, error) {
// 	block, err := aes.NewCipher(commonkey)
// 	if err != nil {
// 		return "", err
// 	}
// 	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return "", err
// 	}
// 	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
// 		[]byte(plaintext))
// 	return hex.EncodeToString(ciphertext), nil

// }
// func AesDecrypt(d string) (string, error) {
// 	ciphertext, err := hex.DecodeString(d)
// 	if err != nil {
// 		return "", err
// 	}
// 	block, err := aes.NewCipher(commonkey)
// 	if err != nil {
// 		return "", err
// 	}
// 	if len(ciphertext) < aes.BlockSize {
// 		return "", errors.New("ciphertext too short")
// 	}
// 	iv := ciphertext[:aes.BlockSize]
// 	ciphertext = ciphertext[aes.BlockSize:]
// 	fmt.Println(len(ciphertext), len(iv))
// 	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
// 	return string(ciphertext), nil
// }
