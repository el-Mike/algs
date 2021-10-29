package queue

const (
	QUEUE_EMPTY_ERROR = "Queue is empty!"
	QUEUE_FULL_ERROR  = "Queue is full!"
)

type QueueError struct {
	Message string
}

func (e *QueueError) Error() string {
	return e.Message
}

func NewQueueEmptyError() *QueueError {
	return &QueueError{
		Message: QUEUE_EMPTY_ERROR,
	}
}

func NewQueueFullError() *QueueError {
	return &QueueError{
		Message: QUEUE_FULL_ERROR,
	}
}
