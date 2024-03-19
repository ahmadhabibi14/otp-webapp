package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	crand "crypto/rand"
)

func GenerateOTP() string {
	_1 := rand.Intn(10)
	_2 := rand.Intn(10)
	_3 := rand.Intn(10)
	_4 := rand.Intn(10)

	otp := fmt.Sprintf("%d%d%d%d", _1, _2, _3, _4)

	return otp
}

func GenerateRandomString(size int) string {
	b := make([]byte, size)
	_, _ = crand.Read(b)
	// Encode the random number to a base64 string
	encode := base64.StdEncoding.EncodeToString(b)
	replacer := strings.NewReplacer(
		"&", "",
		"-", "",
		"+", "",
		"=", "",
		"!", "",
		"/", "",
		`\`, "",
		"#", "",
		"*", "",
		"%", "",
	)
	str := replacer.Replace(encode)
	return str
}
