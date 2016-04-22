package main

import (
	"github.com/kless/osutil/user/crypt/apr1_crypt"
	"github.com/kless/osutil/user/crypt/md5_crypt"
	"github.com/kless/osutil/user/crypt/sha256_crypt"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

// HashPassword is used to generate a password hash of the correct type
func HashPassword(password, hashType string) (string, error) {
	var hash string
	var err error

	switch hashType {
	case "sha512":
		hash, err = GenerateSHA512FromString(password)
	case "sha256":
		hash, err = GenerateSHA256FromString(password)
	case "apr1":
		hash, err = GenerateAPR1FromString(password)
	case "md5":
		hash, err = GenerateMD5FromString(password)
	default:
		hash = "Password cannot be a blank value. Please try again."
	}

	return hash, err
}

// GenerateSHA512FromString creates a SHA-512 password hash from a password
// string which is compatible with Linux operating systems.
func GenerateSHA512FromString(password string) (string, error) {
	crypter := sha512_crypt.New()
	passwordByteStream := []byte(password)
	passwordHash, err := crypter.Generate(passwordByteStream, []byte{})
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}

// GenerateSHA256FromString creates a SHA-256 password hash from a password
// string which is compatible with Linux operating systems.
func GenerateSHA256FromString(password string) (string, error) {
	crypter := sha256_crypt.New()
	passwordByteStream := []byte(password)
	passwordHash, err := crypter.Generate(passwordByteStream, []byte{})
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}

// GenerateAPR1FromString creates a APR1 password hash from a password
// string which is compatible with Linux operating systems.
func GenerateAPR1FromString(password string) (string, error) {
	crypter := apr1_crypt.New()
	passwordByteStream := []byte(password)
	passwordHash, err := crypter.Generate(passwordByteStream, []byte{})
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}

// GenerateMD5FromString creates an MD5 password hash from a password
// string which is compatible with Linux operating systems.
func GenerateMD5FromString(password string) (string, error) {
	crypter := md5_crypt.New()
	passwordByteStream := []byte(password)
	passwordHash, err := crypter.Generate(passwordByteStream, []byte{})
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
