package cmd

import (
	"encoding/hex"

	"github.com/kless/osutil/user/crypt/apr1_crypt"
	"github.com/kless/osutil/user/crypt/md5_crypt"
	"github.com/kless/osutil/user/crypt/sha256_crypt"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// GenerateSHA3ShakeSum256FromString creates a SHA3 SHAKE-256 hash from a
// password strings.
func GenerateSHA3ShakeSum256FromString(password string) string {
	passwordByteStream := []byte(password)
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	passwordHash := make([]byte, 64)
	sha3.ShakeSum256(passwordHash, passwordByteStream)

	return hex.EncodeToString(passwordHash)
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

// GenerateMD4WindowsNTLMFromString creates a MD4 based password hash from a
// password string which is compatible with Linux / BSD operating systems.
func GenerateMD4WindowsNTLMFromString(password string) string {
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	hasher := md4.New()
	t := transform.NewWriter(hasher, enc)
	t.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateBcryptFromString creates a Bcrypt password hash from a password
// string which is compatible with Linux / BSD operating systems.
func GenerateBcryptFromString(password string) (string, error) {
	passwordByteStream := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordByteStream, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
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
