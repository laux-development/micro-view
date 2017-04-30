package models

type ProfileProvider interface {
	Address() string
	Company() string
	DOB() string
	FirstName() string
	Gender() string
	GithubName() string
	JobRole() string
	LastName() string
}
