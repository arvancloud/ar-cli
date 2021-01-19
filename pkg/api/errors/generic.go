package errors

type GenericApiCallError interface {
	Error() string
	Body() []byte
	Model() interface{}
}
