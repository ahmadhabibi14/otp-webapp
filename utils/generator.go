package utils

import (
	"encoding/base64"
	"math/big"
	"strconv"
	"strings"

	crand "crypto/rand"
)

func GenerateOTP(digits int) string {
	otpNumbers := make([]int, 0)

	for i := 0; i < digits; i++ {
		rnum, _ := crand.Int(crand.Reader, big.NewInt(9))
		otpNumbers = append(otpNumbers, int(rnum.Int64()))
	}

	strNumbers := make([]string, len(otpNumbers))
	for i, num := range otpNumbers {
		strNumbers[i] = strconv.Itoa(num)
	}

	otp := strings.Join(strNumbers, "")

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
