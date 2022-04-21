package test

import (
	"github.com/jmnote/flex"
	"testing"
)

func TestSlicing(t *testing.T) {
	var got string
	var want string

	f, err := flex.NewFromJSON(`["g","o","l","a","n","g"]`)
	if err != nil {
		t.Fatal(err)
	}

	got = f.Get("[1:4]").ToJSON()
	want = `["o","l","a"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
