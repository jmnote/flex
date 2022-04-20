package flex

import (
	"testing"
)

func TestNilFlex(t *testing.T) {
	var want string
	var got string

	f := NewFromObject(nil)
	want = `&flex.Flex{object:interface {}(nil)}`
	got = FmtStringSharp(f.Get(""))
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestMultiFlex(t *testing.T) {
	var want string
	var got string

	f1, err := NewFromJSONString(`{"a":42,"b":"x"}`)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := NewFromJSONString("[1,2,3,4]")
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f1.Get(""))
	want = `&flex.Flex{object:map[string]interface {}{"a":42, "b":"x"}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f2.Get("[3]"))
	want = `&flex.Flex{object:4}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f3 := MultiFlex(f1, f2)

	got = FmtStringSharp(f3)
	want = `&flex.Flex{object:[]interface {}{map[string]interface {}{"a":42, "b":"x"}, []interface {}{1, 2, 3, 4}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f3.Get("[1]"))
	want = FmtStringSharp(f2)
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f3.Get("[0].a"))
	want = `&flex.Flex{object:42}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
