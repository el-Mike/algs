package hashtable

const (
	KEY_NOT_EMPTY_ERROR = "Key is not empty!"
)

type HashtableError struct {
	Message string
}

func (e *HashtableError) Error() string {
	return e.Message
}

func NewKeyNotEmptyError() *HashtableError {
	return &HashtableError{
		Message: KEY_NOT_EMPTY_ERROR,
	}
}
