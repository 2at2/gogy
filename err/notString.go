package err

type NotStringError struct {
	Key string
}

func (e *NotStringError) Error() string {
	return "Argument " + e.Key + " must be string"
}
