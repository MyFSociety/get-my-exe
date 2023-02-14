package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ValidateFile validates a file against a hash
func ValidateFile(file *os.File, hash string, hashType string) bool {

	var hashTypeFunc = sha1.New()

	switch strings.ToLower(hashType) {
	case "sha256":
		hashTypeFunc = sha256.New()
	case "sha512":
		hashTypeFunc = sha512.New()
	default:
		fmt.Println("Invalid hash type")
		os.Exit(1)
	}

	if _, err := io.Copy(hashTypeFunc, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hashTypeFunc.Sum(nil)) == hash
}
