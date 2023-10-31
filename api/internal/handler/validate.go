package handler

import (
	"github.com/friendsofgo/errors"
	"regexp"
	"unicode"
)

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
