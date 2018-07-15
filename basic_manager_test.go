package gomanage

import "testing"

func TestBasicManager_Add(t *testing.T) {
	bm := NewBasicManager()

	// basic add
	ret := bm.Add("testm", managerFuncSpecForTests)
	if ret != nil {
		t.Error("Failed adding a new manager function")
	}

	// try adding again with the same key
	ret = bm.Add("testm", managerFuncSpecForTests)
	if ret == nil {
		t.Error("Shouldn't be able to add another manager function with the same key")
	}

	// try adding again with a different key
	ret = bm.Add("testm1", managerFuncSpecForTests)
	if ret != nil {
		t.Error("Failed adding a new manager function with a different key")
	}
}

func TestBasicManager_Remove(t *testing.T) {
	bm := NewBasicManager()
	// basic add
	bm.Add("testm", managerFuncSpecForTests)

	f, e := bm.Remove("testm")
	if e != nil {
		t.Error("Error removing function", e)
	}
	if f == nil {
		t.Error("Removed function wasn't returned")
	}
	ret, _ := f.managerFunc(map[string]string{"first":"testparam"})
	if ret != "testparam" {
		t.Error("Problem calling the removed function")
	}

	_, e = bm.Remove("does_not_exist")
	if e == nil {
		t.Error("No error was returned while trying to remove non existing key")
	}
}

func TestBasicManager_Start(t *testing.T) {
	bm := NewBasicManager()
	bm.Add("testm", managerFuncSpecForTests)
	bm.Start(8080)
}

var managerFuncForTests = func (params map[string]string) (string, error) {
	return params["first"], nil
}

var managerFuncSpecForTests = Func(managerFuncForTests, "first")