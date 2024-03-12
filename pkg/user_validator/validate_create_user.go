package uservalidator

import (
	"fmt"
	"regexp"
)

func ValidateCreateUserRequest(
	firstname string,
	lastname string,
	username string,
	password string,
) map[string]string {
	errs := make(map[string]string)

	// validate first name
	if len(firstname) < MIN_FIRSTNAME_SIZE {
		errs["firstName"] = fmt.Sprintf("First name must be at least %d characters long", MIN_FIRSTNAME_SIZE)
	}
	if len(firstname) > MAX_FIRSTNAME_SIZE {
		errs["firstName"] = fmt.Sprintf("First name must be at most %d characters long", MAX_FIRSTNAME_SIZE)
	}

	// validate last name
	if len(lastname) < MIN_LASTNAME_SIZE {
		errs["lastName"] = fmt.Sprintf("Last name must be at least %d characters long", MIN_LASTNAME_SIZE)
	}
	if len(lastname) > MAX_LASTNAME_SIZE {
		errs["lastName"] = fmt.Sprintf("Last name must be at most %d characters long", MAX_LASTNAME_SIZE)
	}

	// validate username
	if len(username) < MIN_USERNAME_SIZE {
		errs["username"] = fmt.Sprintf("Username must be at least %d characters long", MIN_USERNAME_SIZE)
	}
	if len(username) > MAX_USERNAME_SIZE {
		errs["username"] = fmt.Sprintf("Username must be at most %d characters long", MAX_USERNAME_SIZE)
	}
	if CheckUsernamePattern(username) {
		errs["username"] = "Username pattern does not follow system policy"
	}

	// validate password

	return nil
}

func CheckUsernamePattern(username string) bool {
	re := regexp.MustCompile(USERNAME_PATTERN)

	return re.MatchString(username)
}

func CheckPasswordPattern(password string) bool {
	// Compile the regex pattern
	regex := regexp.MustCompile(PASSWORD_PATTERN)

	// Check if the password matches the pattern
	if !regex.MatchString(password) {
		return false
	}

	// Check for at least one lowercase letter
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one uppercase letter
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one digit
	digitRegex := regexp.MustCompile(`\d`)
	if !digitRegex.MatchString(password) {
		return false
	}

	// Check for at least one special character
	specialRegex := regexp.MustCompile(`[!@#$%^&*]`)
	if !specialRegex.MatchString(password) {
		return false
	}

	return true
}
