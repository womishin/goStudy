package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("tom")
	t.Log(s)
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("tom")
	t.Log(s)
}
