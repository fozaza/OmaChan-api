package module

import (
	"errors"

	"github.com/go-crypt/crypt"
	"github.com/go-crypt/crypt/algorithm/argon2"
	"github.com/gofiber/fiber/v2"
)

// for check password is sefty?
func ch_vf_pw(password string) ErrorOmaChan {
	if len(password) < 6 {
		return ErrorOmaChan{errors.New("OmaChan >>> error bad password recomment >= 6"), fiber.StatusBadRequest}
	}
	return ErrorOmaChan{nil, 0}
}

// function encode password
func Cr_pw(password string) (string, ErrorOmaChan) {
	if err := ch_vf_pw(password); err.Err != nil {
		return "", err
	}

	// create argon2
	hasher, err := argon2.New(argon2.WithProfileRFC9106LowMemory())
	if err != nil {
		return "", ErrorOmaChan{err, fiber.StatusInternalServerError}
	}

	// hash password
	digest, err := hasher.Hash(password)
	if err != nil {
		return "", ErrorOmaChan{err, fiber.StatusInternalServerError}
	}

	// check password
	if err := Ch_pw(password, digest.Encode()); err != nil {
		return "", ErrorOmaChan{errors.New("OmaChan >>> error bad Decode password"), fiber.StatusInternalServerError}
	}
	return digest.String(), ErrorOmaChan{nil, 0}
}

// function check password
func Ch_pw(password string, digest string) error {
	if _, err := crypt.CheckPassword(password, digest); err != nil {
		return errors.New("OmaChan >>> password check not match")
	}
	return nil
}
