package flex

import (
	"testing"
)

func TestSetFlex(t *testing.T) {
	var got string
	var want string

	jsonString := `{"a":"hello","b":{"nums":[1,2,3],"lorem":"ipsum"}}`
	f, err := NewFromJSONString(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f.Get(".b"))
	want = `&flex.Flex{object:map[string]interface {}{"lorem":"ipsum", "nums":[]interface {}{1, 2, 3}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f.Set(".b.lorem", "world")
	got = FmtStringSharp(f.Get("."))
	want = `&flex.Flex{object:map[string]interface {}{"a":"hello", "b":map[string]interface {}{"lorem":"world", "nums":[]interface {}{1, 2, 3}}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f, err = NewFromJSONString(`["a",{"nums":[1,2,3,4]},42]`)
	if err != nil {
		t.Fatal(err)
	}

	f.Set("[1].nums[2]", "world")
	got = FmtStringSharp(f.Get("."))
	want = `&flex.Flex{object:[]interface {}{"a", map[string]interface {}{"nums":[]interface {}{1, 2, "world", 4}}, 42}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
