package hasherModels

// HashedPassword represents the information needed to fill in the hashed
// password portion of the website
type HashedPassword struct {
	Title        string
	PasswordHash string
}
