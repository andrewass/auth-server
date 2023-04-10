package user

type InvalidCredentialsError struct {
	msg string
}

func (error *InvalidCredentialsError) Error() string {
	return error.msg
}

type CredentialsConflictError struct {
	msg string
}

func (error *CredentialsConflictError) Error() string {
	return error.msg
}
