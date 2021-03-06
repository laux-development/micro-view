package view

import (
	"html/template"
	"io"

	m "github.com/laux-development/micro-view/models"
)

// templates
// parse and cache templates
var templates = template.Must(template.ParseFiles(
	"templates/components/login/login-element.tmpl",
	"templates/components/login/login-cdn-library.tmpl",
	"templates/components/login/login-script.tmpl",
	"templates/components/login/login-return-script.tmpl",
	"pages/login.tmpl",
	"pages/home.tmpl",
	"pages/profile.tmpl",
))

// View
type View struct{}

// Home view
func (v View) Home(wr io.Writer) {
	templates.ExecuteTemplate(wr, "home", nil)
}

// Login view
func (v View) Login(wr io.Writer) {
	templates.ExecuteTemplate(wr, "login", nil)
}

// Profile view
func (v View) Profile(wr io.Writer, p m.ProfileProvider) {

	// create instance variable for view
	profile := struct {
		Address    string
		Company    string
		DOB        string
		FirstName  string
		Gender     string
		GithubName string
		JobRole    string
		LastName   string
	}{
		p.Address(),
		p.Company(),
		p.DOB(),
		p.FirstName(),
		p.Gender(),
		p.GithubName(),
		p.JobRole(),
		p.LastName(),
	}

	templates.ExecuteTemplate(wr, "profile", profile)
}
