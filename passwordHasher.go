package passwordhasher

import "github.com/kless/osutil/user/crypt/sha512_crypt"

// GenerateSHA512 creates a SHA-512 password hash which is compatible with Linux
// operating systems.
func GenerateSHA512(password []byte) (string, error) {
	crypter := sha512_crypt.New()
	passwordByteStream := password
	passwordHash, err := crypter.Generate(passwordByteStream, []byte{})
	if err != nil {
		return "", err
	}

	return passwordHash, nil
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
