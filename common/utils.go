package common

import (
	"crypto/rand"
	"encoding/hex"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// GenerateTmpDirName will generate a 10-character
// random string to be used as a temp dir in the
// workspace directory.
func GenerateTmpDirName() string {
	val, err := randomHex(10)

	if err != nil {
		commonLog.Panicf("! Could not generate tmp dir name")
	}

	return val
}
