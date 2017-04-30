package models

type ProfileProvider interface {
	GithubName() string
	FirstName() string
	LastName() string
}
