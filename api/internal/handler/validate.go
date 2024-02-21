package handler

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/friendsofgo/errors"
)

var validTLDs = []string{"com", "org", "net"}

// Group User
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

// Group Friends
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

func (c CommonFriends) validate() error {
	if err := validateEmails(c.Emails); err != nil {
		return err
	}
	return nil
}

// Group Subscription
func (s AddSubscription) validate() error {
	emails := []string{s.Requester, s.Target}
	if err := validateEmails(emails); err != nil {
		return err
	}
	return nil
}

func (b Block) validate() error {
	emails := []string{b.Requester, b.Target}
	if (len(emails) != 2) || (emails[0] == emails[1]) {
		return errors.New("Please insert two different emails")
	}

	if requesterError := validateEmail(b.Requester); requesterError != nil {
		return errors.New("Invalid format of requester email: " + requesterError.Error())
	}
	if targetError := validateEmail(b.Target); targetError != nil {
		return errors.New("Invalid format of target email: " + targetError.Error())
  }
  return nil
}

// Local Functions
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

	matches := tldRegex.FindStringSubmatch(email)

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
