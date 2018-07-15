package gomanage

type ManagerFunc func (...string) (string, error)

type Manager interface {

	// Add adds a function that's exposed as a manager endpoint
	Add(key string, managerFunc ManagerFunc) error

	// Remove an already added manager function
	Remove(key string) (ManagerFunc, error)

	// List lists the manager functions available as a map
	List() map[string]ManagerFunc

	// ListEndpoints lists the available manager endpoints in json format
	// spec can be something like swagger
	ListEndpoints(spec string) string

	// Start starts the manager http server on the given port
	Start(port int) error

}