package main

// CreateHash represents the information needed to fill in the create password
// hash website
type CreateHash struct {
	Title string
}

// HashedPassword represents the information needed to fill in the hashed
// password portion of the website
type HashedPassword struct {
	Title        string
	PasswordHash string
}
