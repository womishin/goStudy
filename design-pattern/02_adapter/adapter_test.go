package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	t.Log(target.Request())
}