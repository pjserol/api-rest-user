package utils

import "regexp"

// ValidateEmail validate the format of the email
func ValidateEmail(email string) bool {
	var rxEmail = regexp.MustCompile("^.+@.+\\..+$")

	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}

	return true
}
