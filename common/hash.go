package common

import (
	"bytes"
	"crypto/aes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值
func MD5String(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

// GetRandomString 生成随机字符串
func GetRandomString(lenS int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenS; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// ShortHash 生成简短hash
func ShortHash(input int64) string {
	binary := int64(62)
	stack := make([]string, 0)
	sign := ""
	if input < 0 {
		sign = "-"
	}

	input = int64(math.Abs(float64(input)))

	table := func(num int) string {
		//打乱顺序
		var t = strings.Split("9132465780zbdcefghjiklnmoqprtsuwvyxaZBCEDFHGIKJMLNPURQSTOVWXYA", "")
		return t[num]
	}
	for input >= binary {
		i := int(input % binary)
		input = int64(math.Floor(float64(input) / float64(binary)))
		stack = append(stack, table(i))
	}
	if input > 0 {
		stack = append(stack, table(int(input)))
	}

	return sign + strings.Join(stack, "")
}

// AES加密
func AesEncrypt(src, key []byte) (out []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		return nil, err
	}
	out = make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AES解密
func AesDecrypt(src, key []byte) (out []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out = make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		return nil, err
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ComputeHmacSha256(message []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}
