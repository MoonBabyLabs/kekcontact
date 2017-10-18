package kekcontact


type Company struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	City string `json:"city"`
	Region string `json:"region"`
	PostalCode string `json:"postal_code"`
	CountryCode string `json:"country_code"`
	Id string `json:"id"`
	kekCompanyId string
}

type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Id string `json:"id"`
	Address string `json:"address"`
	City string `json:"city"`
	Region string `json:"region"`
	PostalCode string `json:"postal_code"`
	CountryCode string `json:"country_code"`
	kekContactId string
	Company Company
}
