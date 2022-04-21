package test

import (
	"github.com/jmnote/flex"
	"testing"
)

func TestSetFlex(t *testing.T) {
	var got string
	var want string

	jsonString := `{"a":"hello","b":{"nums":[1,2,3],"lorem":"ipsum"}}`
	f, err := flex.NewFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	got = f.Get(".b").ToJSON()
	want = `{"lorem":"ipsum","nums":[1,2,3]}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f.Set(".b.lorem", "world")
	got = f.Get(".").ToJSON()
	want = `{"a":"hello","b":{"lorem":"world","nums":[1,2,3]}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f, err = flex.NewFromJSON(`["a",{"nums":[1,2,3,4]},42]`)
	if err != nil {
		t.Fatal(err)
	}

	f.Set("[1].nums[2]", "world")
	got = f.Get(".").ToJSON()
	want = `["a",{"nums":[1,2,"world",4]},42]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
