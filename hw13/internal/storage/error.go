package storage

//Error type of storage
type Error string

func (e Error) Error() string {
	return string(e)
}

//Errors
const (
	ErrDateBusy Error = "this time is already in use"
	ErrNotFound Error = "event not found"
)
