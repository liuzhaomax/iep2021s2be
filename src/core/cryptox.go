/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/2 12:05
 * @version     v1.0
 * @filename    cryptox.go
 * @description
 ***************************************************************************/
package core

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"os"
	"time"
)

func MD5(byt []byte) string {
	hash := md5.New()
	_, _ = hash.Write(byt)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func MD5Str(str string) string {
	return MD5([]byte(str))
}

func SHA1(byt []byte) string {
	hash := sha1.New()
	_, _ = hash.Write(byt)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SHA1Str(str string) string {
	return SHA1([]byte(str))
}

func SHA1MD5Str(str string) string {
	return SHA1Str(MD5Str(str))
}

func BASE64Encode(byt []byte) string {
	encoded := base64.StdEncoding.EncodeToString(byt)
	return encoded
}

func BASE64EncodeStr(str string) string {
	encoded := BASE64Encode([]byte(str))
	return encoded
}

func BASE64Decode(str string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	return decoded, err
}

func BASE64DecodeStr(str string) (string, error) {
	decoded, err := BASE64Decode(str)
	if err == nil {
		decodedStr := string(decoded)
		return decodedStr, nil
	}
	return "", err
}

func GenRsaKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	projPath := GetProjectPath()
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	// public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err := os.Create(projPath + "/bin/public.pem")
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, err
}

func PublicKeyToString() (string, error) {
	projPath := GetProjectPath()
	file, err := os.Open(projPath + "/bin/public.pem")
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(file)
	var str string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		str = str + line
	}
	err = file.Close()
	if err != nil {
		return "", err
	}
	//str = strings.Replace(str, "\n", "", -1)
	return str, err
}

func RSADecrypt(privateKey *rsa.PrivateKey, encryptedStr string) (string, error) {
	cipherTextB64, _ := base64.StdEncoding.DecodeString(encryptedStr)
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTextB64)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

func RSAEncrypt(publicKey *rsa.PublicKey, str string) (string, error) {
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(str))
	if err != nil {
		return "", err
	}
	encryptedStr := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedStr, nil
}

func GenerateToken(text string, duration time.Duration) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": text,
		"exp": time.Now().Add(duration).Unix(),
	})
	token, err := at.SignedString([]byte(ctx.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(ctx.JWTSecret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
