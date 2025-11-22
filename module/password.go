package module

import (
	"errors"

	"github.com/go-crypt/crypt"
	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
)

// function check password
func Cr_pw(password string) (string, error) {
	// create argon2
	hasher, err := argon2.New(argon2.WithProfileRFC9106LowMemory())
	if err != nil {
		return "", err
	}

	// hash password
	digest, err := hasher.Hash(password)
	if err != nil {
		return "", err
	}

	// check password
	if err := Ch_pw(password, digest); err != nil {
		return "", errors.New("OmaChan >>> error bad Decode password")
	}
	return digest.String(), nil
}

// function check password
func Ch_pw(password string, digest algorithm.Digest) error {
	if _, err := crypt.CheckPassword(password, digest.Encode()); err != nil {
		return errors.New("OmaChan >>> password check not match")
	}
	return nil
}
