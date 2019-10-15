package common

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateTmpDirName() string {
	val, err := randomHex(10)
	if err != nil {
		log.Panicf("! Could not generate tmp dir name")
	}

	return val
}
