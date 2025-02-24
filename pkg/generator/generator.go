package generator

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

type Options struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

func GeneratePassword(opts Options) (string, error) {
	var charset string
	if opts.Uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if opts.Lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if opts.Numbers {
		charset += "0123456789"
	}
	if opts.Symbols {
		charset += "!@#$%^&*()_+-=[]{}|;:,.<>?"
	}

	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	}

	bytes := make([]byte, opts.Length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	result := strings.Builder{}
	for _, b := range bytes {
		result.WriteByte(charset[int(b)%len(charset)])
	}
	return result.String(), nil
}

func GenerateHex(length int) (string, error) {
	bytes := make([]byte, (length+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

func GenerateBase64(length int) (string, error) {
	bytes := make([]byte, (length*3)/4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes)[:length], nil
}
