package base

type Restful interface {
	Get() *statusError
	Post() *statusError
	Put() *statusError
	Delete() *statusError
}
