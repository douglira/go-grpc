package models

type Student struct {
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	IdentityNumber            string `json:"identityNumber"`
	GeneralRegistrationNumber string `json:"generalRegistrationNumber"`
}
