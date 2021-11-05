package hashtable

const (
	KEY_NOT_EMPTY_ERROR       = "Key is not empty!"
	KEY_ALREADY_EXISTS_ERROR  = "Key already exists!"
	KEY_DOES_NOT_EXIST_ERROR  = "Key does not exist!"
	INCORRECT_ITEM_TYPE_ERROR = "Item is not a correct type!"
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

func NewKeyAlreadyExistsError() *HashtableError {
	return &HashtableError{
		Message: KEY_ALREADY_EXISTS_ERROR,
	}
}

func NewKeyDoesNotExistError() *HashtableError {
	return &HashtableError{
		Message: KEY_DOES_NOT_EXIST_ERROR,
	}
}

func NewIncorrectItemTypeError() *HashtableError {
	return &HashtableError{
		Message: INCORRECT_ITEM_TYPE_ERROR,
	}
}
