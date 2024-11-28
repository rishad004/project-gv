package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"math/rand"

	"golang.org/x/crypto/argon2"
)

func HashPassword(Password string) (string, string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(Password), salt, 1, 64*1024, 4, 32)

	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)

	return hashEncoded, saltEncoded, err
}

func VerifyPassword(Password, Hashed, Salted string) error {
	hash, err := base64.RawStdEncoding.DecodeString(Hashed)
	if err != nil {
		return err
	}

	salt, err := base64.RawStdEncoding.DecodeString(Salted)
	if err != nil {
		return err
	}

	computedHash := argon2.IDKey([]byte(Password), salt, 1, 64*1024, 4, 32)

	if subtle.ConstantTimeCompare(hash, computedHash) == 1 {
		return nil
	}

	return errors.New("wrong credentials")
}