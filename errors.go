package gophercloud

type InvalidInputError struct {
	Function string
	Argument string
	Value    interface{}
	Message  string
}

func (e *InvalidInputError) Error() string {
	return e.Message
}

type defaultError401 struct {
	*UnexpectedResponseCodeError
}
type defaultError404 struct {
	*UnexpectedResponseCodeError
}
type defaultError405 struct {
	*UnexpectedResponseCodeError
}

func (e defaultError401) Error() string {
	return "Authentication failed"
}
func (e defaultError404) Error() string {
	return "Resource not found"
}
func (e defaultError405) Error() string {
	return "Method not allowed"
}

// Error401er is the interface resource error types implement to override the error message
// from a 401 error.
type Error401er interface {
	Error401(*UnexpectedResponseCodeError) error
}

// Error404er is the interface resource error types implement to override the error message
// from a 404 error.
type Error404er interface {
	Error404(*UnexpectedResponseCodeError) error
}

// Error405er is the interface resource error types implement to override the error message
// from a 405 error.
type Error405er interface {
	Error405(*UnexpectedResponseCodeError) error
}
