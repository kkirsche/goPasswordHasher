package passwordhasher

// CreateHash represents the information needed to fill in the create password
// hash website
type CreateHash struct {
	Title       string
	Description string
}

// HashedPassword represents the information needed to fill in the hashed
// password portion of the website
type HashedPassword struct {
	Title       string
	Description string

	SHA3ShakeSum256 string
	SHA3ShakeSum128 string
	SHA512          string
	SHA256          string
	NTLM            string
	Bcrypt          string
	APR1            string
	MD5             string
}
