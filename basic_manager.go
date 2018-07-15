package gomanage

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// basic implementation

type basicManager struct {

	managerFunctions map[string]*ManagerFuncSpec
}

func NewBasicManager() Manager {
	return &basicManager{make(map[string]*ManagerFuncSpec)}
}

func (bm *basicManager) Add(key string, managerFuncSpec *ManagerFuncSpec) error {
	if _, ok := bm.managerFunctions[key]; ok {
		return errors.New("The manager function already exists. Use different key or remove existing")
	}
	bm.managerFunctions[key] = managerFuncSpec
	return nil
}

func (bm *basicManager) Remove(key string) (*ManagerFuncSpec, error) {
	if function, ok := bm.managerFunctions[key]; ok {
		delete(bm.managerFunctions, key)
		return function, nil
	} else {
		return nil, errors.New("The manager function does not exist")
	}
}


func (bm *basicManager) List() map[string]*ManagerFuncSpec {
	return copyOf(bm.managerFunctions)
}


func (bm *basicManager) ListEndpoints(spec string) string {
	panic("Not implemented")
}


func (bm *basicManager) Start(port int) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	return nil
}

func copyOf(origin map[string]*ManagerFuncSpec) map[string]*ManagerFuncSpec {
	targetMap := make(map[string]*ManagerFuncSpec)
	for key, value := range origin {
		targetMap[key] = value
	}
	return targetMap
}

