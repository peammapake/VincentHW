package exception

var (
	// System and Connection Errors
	UnknownError    = NewError(UNKNOWN, "Internal server error")
	DatabaseError   = NewError(DATABASE_ERROR, "Database error")
	ConnectionError = NewError(CONNECTION_ERROR, "Connection error")
)
