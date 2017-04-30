package view

import (
	"bytes"
	"strings"
	"testing"
)

type testProfileModel struct{}

func (pm testProfileModel) GithubName() string {
	return "bleeppurple"
}

func (pm testProfileModel) FirstName() string {
	return "james"
}

func (pm testProfileModel) LastName() string {
	return "laux"
}

func (pm testProfileModel) Address() string {
	return "a nice road, in a nice place, goodpostcode"
}

func (pm testProfileModel) Company() string {
	return "Laux Development"
}

func (pm testProfileModel) JobRole() string {
	return "Developer"
}

func (pm testProfileModel) DOB() string {
	return "20/01/1990"
}

func (pm testProfileModel) Gender() string {
	return "male"
}

func TestMain(t *testing.T) {

	m := testProfileModel{}
	v := View{}

	buf := new(bytes.Buffer)
	v.Profile(buf, m)

	if !(strings.Contains(buf.String(), "goodpostcode")) {
		t.Errorf("expected to find 'goodpostcode' but got an error")
	}
}
