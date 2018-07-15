package gomanage

import "testing"

func TestBasicManager_Add(t *testing.T) {
	bm := NewBasicManager()

	// basic add
	ret := bm.Add("testm", managerFuncForTests)
	if ret != nil {
		t.Error("Failed adding a new manager function")
	}

	// try adding again with the same key
	ret = bm.Add("testm", managerFuncForTests)
	if ret == nil {
		t.Error("Shouldn't be able to add another manager function with the same key")
	}

	// try adding again with a different key
	ret = bm.Add("testm1", managerFuncForTests)
	if ret != nil {
		t.Error("Failed adding a new manager function with a different key")
	}
}

func TestBasicManager_Remove(t *testing.T) {
	bm := NewBasicManager()
	// basic add
	bm.Add("testm", managerFuncForTests)

	f, e := bm.Remove("testm")
	if e != nil {
		t.Error("Error removing function", e)
	}
	if f == nil {
		t.Error("Removed function wasn't returned")
	}
	ret, _ := f("testparam")
	if ret != "testparam" {
		t.Error("Problem calling the removed function")
	}

	_, e = bm.Remove("does_not_exist")
	if e == nil {
		t.Error("No error was returned while trying to remove non existing key")
	}
}

var managerFuncForTests = func (params ...string) (string, error) {
	return params[0], nil
}