package kekcontact


type Company struct {
	Name string
	Phone string
	Address string
	City string
	Region string
	PostalCode string
	Id string
	kekCompanyId string
}

type Contact struct {
	Name string
	Email string
	Phone string
	Id string
	Address string
	City string
	Region string
	PostalCode string
	CountryCode string
	kekContactId string
	*Company
}
