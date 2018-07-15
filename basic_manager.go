package gomanage

import "errors"

// basic implementation

type basicManager struct {

	managerFunctions map[string]ManagerFunc
}

func NewBasicManager() Manager {
	return &basicManager{make(map[string]ManagerFunc)}
}

func (bm *basicManager) Add(key string, managerFunc ManagerFunc) error {
	if _, ok := bm.managerFunctions[key]; ok {
		return errors.New("The manager function already exists. Use different key or remove existing")
	}
	bm.managerFunctions[key] = managerFunc
	return nil
}

func (bm *basicManager) Remove(key string) (ManagerFunc, error) {
	if function, ok := bm.managerFunctions[key]; ok {
		delete(bm.managerFunctions, key)
		return function, nil
	} else {
		return nil, errors.New("The manager function does not exist")
	}
}


func (bm *basicManager) List() map[string]ManagerFunc {
	return copyOf(bm.managerFunctions)
}


func (bm *basicManager) ListEndpoints(spec string) string {
	panic("Not implemented")
}


func (bm *basicManager) Start(port int) error {
	return nil
}

func copyOf(origin map[string]ManagerFunc) map[string]ManagerFunc {
	targetMap := make(map[string]ManagerFunc)
	for key, value := range origin {
		targetMap[key] = value
	}
	return targetMap
}

