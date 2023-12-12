package handler

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/friendsofgo/errors"
)

var validTLDs = []string{"com", "org", "net"}

func (u User) validate() error {
	if len(u.Name) == 0 {
		return errors.New("Name cannot be blank")
	}
	firstChar := rune(u.Name[0])
	if !unicode.IsUpper(firstChar) {
		return errors.New("Name Invalid")
	}
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}
func (f Friends) validate() error {
	if err := validateEmails(f.Emails); err != nil {
		return err
	}
	return nil
}
func (e FriendsList) validate() error {
	if err := validateEmail(e.Email); err != nil {
		return err
	}
	return nil
}
func (i UpdateTopic) validate() (string, error) {
	if err := validateEmail(i.Sender); err != nil {
		return "", err
	}

	extractEmailPattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
	re := regexp.MustCompile(extractEmailPattern)
	mentionedEmail := re.FindString(i.Text)
	if err := validateEmail(mentionedEmail); err != nil {
		return "", err
	}
	return mentionedEmail, nil
}

func validateEmail(email string) error {
	// Check Email length
	lengthIsValid := 0 < len(email) && len(email) <= 320
	if !(lengthIsValid) {
		return errors.New("Invalid Email Length")
	}

	// Check Email format
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, email)
	if !(match) {
		return errors.New("Invalid Email Format")
	}

	// Check Email TLD
	tldRegex := regexp.MustCompile(fmt.Sprintf("\\.(%s)$", strings.Join(validTLDs, "|")))

	// Find the TLD in the email
	matches := tldRegex.FindStringSubmatch(email)

	// Check if a valid TLD is found
	if len(matches) == 0 {
		return errors.New("Invalid Email TLD")
	}
	return nil
}
func validateEmails(emails []string) error {
	// Check number of input and avoid repeat inputs
	if (len(emails) != 2) || (emails[0] == emails[1]) {
		return errors.New("Please insert two different emails")
	}

	// Check each email with validateEmail function
	for _, email := range emails {
		if err := validateEmail(email); err != nil {
			return err
		}
	}
	return nil
}
