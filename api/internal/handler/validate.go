package handler

import (
	"github.com/friendsofgo/errors"
	"regexp"
	"unicode"
)

type Validator interface {
	validate() error
}

func (u User) validate() error {

	if len(u.Name) == 0 {
		return errors.New("Name cannot be blank")
	}
	firstChar := rune(u.Name[0])
	if !unicode.IsUpper(firstChar) {
		return errors.New("Name Invalid")
	}

	if len(u.Email) == 0 {
		return errors.New("Email cannot be blank")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, u.Email)
	if !(match) {
		return errors.New("Email invalid")
	}
	return nil
}

func (f Friends) validate() error {

	if (len(f.Emails) < 2) || (f.Emails[0] == f.Emails[1]) {
		return errors.New("Please insert at least two different emails")
	}

	if len(f.Emails[0]) == 0 {
		return errors.New("The first email is blank")
	}
	if len(f.Emails[1]) == 0 {
		return errors.New("The seconds email is blank")
	}
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`

	match, _ := regexp.MatchString(emailPattern, f.Emails[0])
	if !(match) {
		return errors.New("The first email's invalid")
	}

	match, _ = regexp.MatchString(emailPattern, f.Emails[1])
	if !(match) {
		return errors.New("The seconds email's invalid")
	}
	return nil
}
