package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"
)

func run() {
	lpn := "豫E-MJ893"
	now := time.Now()
	timeNow := fmt.Sprint(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	timeStamp := now.UnixNano()
	fmt.Println(strconv.Itoa(int(timeStamp)))
	fmt.Println(lpn)
	fmt.Println(timeStamp)
	a := lpn + ":" + timeNow
	fmt.Println(a)
	h := sha256.New()
	h.Write([]byte(a))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

func Ecdsa() {
	// 生成公钥和私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	// 公钥是存在在私钥中的，从私钥中读取公钥
	publicKey := &privateKey.PublicKey
	message := []byte("hello,dsa签名")

	// 进入签名操作
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)
	fmt.Println(r, s)
	// 进入验证
	flag := ecdsa.Verify(publicKey, message, r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
	flag = ecdsa.Verify(publicKey, []byte("hello"), r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
}

func AcationParse(level int) (int, int, int, int, error) {
	if level > 15 {
		return 0, 0, 0, 0, errors.New("level is not legal")
	}
	a := level / 8
	level %= 8
	b := level / 4
	level %= 4
	c := level / 2
	d := level % 2
	return a, b, c, d, nil
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncoded), string(pemEncodedPub)
}
func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericpublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericpublicKey.(*ecdsa.PublicKey)
	return privateKey, publicKey
}

func test() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey
	encPriv, encpub := encode(privateKey, publicKey)
	fmt.Println(encPriv)
	fmt.Println(encpub)
	ioutil.WriteFile("./user2sk.txt", []byte(encPriv), 0666)
	ioutil.WriteFile("./user2pk.txt", []byte(encpub), 0666)
	priv2, pub2 := decode(encPriv, encpub)
	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}
}

func main() {
	Ecdsa()
}
