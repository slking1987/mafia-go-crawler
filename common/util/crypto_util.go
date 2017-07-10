package util

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"crypto/sha256"
	"crypto/sha512"
)

func Sha1(s string) string {
	return sha(s, sha1.New())
}

func Sha256(s string) string {
	return sha(s, sha256.New())
}

func Sha512(s string) string {
	return sha(s, sha512.New())
}

func sha(s string, hash hash.Hash) string {
	hash.Write([]byte(s))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}