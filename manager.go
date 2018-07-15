package gomanage

type ManagerFunc func (map[string]string) (string, error)

type ManagerFuncSpec struct {
	managerFunc ManagerFunc
	paramNames []string
}

// Func creates new manager function according to ManagerFuncSpec
func Func(managerFunc ManagerFunc, paramNames ...string) *ManagerFuncSpec {
	return &ManagerFuncSpec{managerFunc, paramNames}
}

type Manager interface {

	// Add adds a function spec that's exposed as a manager endpoint
	Add(key string, managerFuncSpec *ManagerFuncSpec) error

	// Remove an already added manager function
	Remove(key string) (*ManagerFuncSpec, error)

	// List lists the manager functions available as a map
	List() map[string]*ManagerFuncSpec

	// ListEndpoints lists the available manager endpoints in json format
	// spec can be something like swagger
	ListEndpoints(spec string) string

	// Start starts the manager http server on the given port
	Start(port int) error

}