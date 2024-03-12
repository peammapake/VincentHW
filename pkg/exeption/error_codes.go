package exception

type Code string

const (
	NOT_FOUND        Code = "NOT_FOUND"
	BAD_REQUEST      Code = "BAD_REQUEST"
	UNKNOWN          Code = "UNKNOWN"
	DATABASE_ERROR   Code = "DATABASE_ERROR"
	INVALID_ARGUMENT Code = "INVALID_ARGUMENT"
	CONNECTION_ERROR Code = "CONNECTION_ERROR"
)
