package member

import (
	"time"
)

type (
	// EncryptionToken struct {
	// }
	Profile struct {
		Id        string    `json:"_id" bson:"_id,omitempty"`
		FullName  string    `json:"fullName"`
		Username  string    `json:"username" bson:"username"`
		Email     string    `json:"email" bson:"email"`
		CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	}

	CredentialCombindProfile struct {
		*Profile
		Credential *Credential `json:"credential"`
	}

	Credential struct {
		AccessToken string    `json:"accessToken"`
		CreatedAt   time.Time `json:"createdAt"`
	}
)
