package dto

import "time"

type ResponseStatusAndMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserDTO struct {
	UserId    string `json:"userId"`
	NIP       int64  `json:"nip"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type PatientDTO struct {
	IdentityNumber int    `json:"identityNumber"`
	PhoneNumber    string `json:"phoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	CreatedAt      string `json:"createdAt"`
}

type MedicalRecords struct {
	IdentityDetail struct {
		IdentityNumber      int    `json:"identityNumber"`
		PhoneNumber         string `json:"phoneNumber"`
		Name                string `json:"name"`
		BirthDate           string `json:"birthDate"`
		Gender              string `json:"gender"`
		IdentityCardScanImg string `json:"identityCardScanImg"`
	} `json:"identityDetail"`
	Symptoms    string    `json:"symptoms"`
	Medications string    `json:"medications"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   struct {
		Nip    string `json:"nip"`
		Name   string `json:"name"`
		UserId string `json:"userId"`
	} `json:"createdBy"`
}
