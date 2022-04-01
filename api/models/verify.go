package models

import "time"

type VerificationDataType int

const (
	MailConfirmation VerificationDataType = iota + 1
)

type VerificationData struct {
	Email     string               `json:"email"`
	Code      string               `json:"code"`
	ExpiresAt time.Time            `json:"expiresat"`
	Type      VerificationDataType `json:"type"`
}

const verificationSchema = `
		create table if not exists verifications (
			email 		Varchar(100) not null,
			code  		Varchar(10) not null,
			expiresat 	Timestamp not null,
			type        Varchar(10) not null,
			Primary Key (email),
			Constraint fk_user_email Foreign Key(email) References users(email)
				On Delete Cascade On Update Cascade
		)`
