package facade

import "testing"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	t.Log(ret)
}