package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, cipherext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, cipherext, nil)
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	ciphertext := encrypt([]byte("Hello World"), "password")
	fmt.Println(string(ciphertext))

	plaintext := decrypt(ciphertext, "password")
	fmt.Println(string(plaintext))

	fileName := "example.txt"
	encryptFile(fileName, []byte("Hello world"), "password")
	plaintext = decryptFile(fileName, "password")
	fmt.Println(string(plaintext))
}

/*

�s˂А�y!$�L�
           �0�8�����i���        ��4���a
Hello World
Hello world
*/

// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
