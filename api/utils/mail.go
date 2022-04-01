package utils

import (
	"time"
)

// The states a user account is in at any given time.
const (
	StateCreated  = iota // User account has been created but not yet verified.
	StateVerified        // User account has been verified and can be used.
)

type NewUser struct {
	id             int
	state          int
	email          string
	verificationID string
}

func (nu NewUser) SetID(id int) {
	nu.id = id
}

func (nu NewUser) GetID() int {
	return nu.id
}

// We need to be able to copy user IDs.

func (nu NewUser) SetState(state int) {
	nu.state = state
}

// At any time, a user is in exactly one state: StateCreated, StateVerified,
// or StateExpired.
func (nu NewUser) GetState() int {
	return nu.state
}

func (nu NewUser) SetEmail(email string) {
	nu.email = email
}

// The users email address. This package will always change email addresses to
// lowercase before setting or comparing them.
func (nu NewUser) GetEmail() string {
	return nu.email
}

func (nu NewUser) SetVerificationID(verificationID string, created time.Time) {
	nu.verificationID = verificationID
}

// Verification IDs (a 22 character long string and its creation time) are
// used to verify new (or changed) user accounts.
func (nu NewUser) GetVerificationID() (string, time.Time) {
	return nu.verificationID, time.Now()
}
