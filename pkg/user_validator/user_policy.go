package uservalidator

const (
	// name
	MIN_FIRSTNAME_SIZE = 1
	MAX_FIRSTNAME_SIZE = 50
	MIN_LASTNAME_SIZE  = 1
	MAX_LASTNAME_SIZE  = 50

	// username
	MIN_USERNAME_SIZE = 5
	MAX_USERNAME_SIZE = 30
	USERNAME_PATTERN = "^[a-zA-Z]+$"

	// password
	PASSWORD_PATTERN = `^[a-zA-Z\d!@#$%^&*]{8,16}$`
)
