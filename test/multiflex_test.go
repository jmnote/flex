package test

import (
	"github.com/jmnote/flex"
	"testing"
)

func TestNilFlex(t *testing.T) {
	var want string
	var got string

	f := flex.NewFromObject(nil)
	want = `null`
	got = f.ToJSON()
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestToSlice(t *testing.T) {
	var got string
	var want string

	got = flex.NewFromObject(flex.NewFromObject([]int{1, 2, 3, 4}).ToSlice()).ToJSON()
	want = `[1,2,3,4]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = flex.NewFromObject(flex.NewFromObject([]string{"hello", "world"}).ToSlice()).ToJSON()
	want = `["hello","world"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = flex.NewFromObject(flex.NewFromObject([]interface{}{1, 2, "hello", "world"}).ToSlice()).ToJSON()
	want = `[1,2,"hello","world"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	var f *flex.Flex
	var err error

	f, err = flex.NewFromJSON(`[1,2,3,4]`)
	if err != nil {
		t.Fatal(err)
	}

	got = flex.NewFromObject(flex.NewFromObject(f.Object).ToSlice()).ToJSON()
	want = `[1,2,3,4]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f, err = flex.NewFromJSON(`["hello","world"]`)
	if err != nil {
		t.Fatal(err)
	}
	got = flex.NewFromObject(flex.NewFromObject(f.Object).ToSlice()).ToJSON()
	want = `["hello","world"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f, err = flex.NewFromJSON(`[1,2,"hello","world"]`)
	if err != nil {
		t.Fatal(err)
	}

	got = flex.NewFromObject(flex.NewFromObject(f.Object).ToSlice()).ToJSON()
	want = `[1,2,"hello","world"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestMultiFlex(t *testing.T) {
	var want string
	var got string

	f1, err := flex.NewFromJSON(`{"a":42,"b":"x"}`)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := flex.NewFromJSON("[1,2,3,4]")
	if err != nil {
		t.Fatal(err)
	}

	got = f1.ToJSON()
	want = `{"a":42,"b":"x"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f2.Get("[3]").ToJSON()
	want = `4`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f3 := flex.NewFromObject([]int{1, 2, 3, 4})
	got = f3.ToJSON()
	want = `[1,2,3,4]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f3 = flex.Append(f3, "hello")
	got = f3.ToJSON()
	want = `[1,2,3,4,"hello"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f3.Get("[1]").ToJSON()
	want = `2`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f4 := flex.Append(f3, f2.Object)

	got = f4.ToJSON()
	want = `[1,2,3,4,"hello",[1,2,3,4]]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f5 := flex.Append(flex.New(), 1, 2, 3, 4, "hello")
	got = f5.ToJSON()
	want = `[1,2,3,4,"hello"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

}

func TestMap(t *testing.T) {
	var got string
	var want string

	f := flex.New()

	got = f.ToJSON()
	want = `null`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".hello").ToJSON()
	want = `null`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f.Set(".hello", "world")

	got = f.ToJSON()
	want = `{"hello":"world"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".hello").ToJSON()
	want = `"world"`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestSlice(t *testing.T) {
	var got string
	var want string

	f := flex.New()
	f = flex.Append(f, "hello")
	f = flex.Append(f, "world", "lorem", "ipsum")

	got = f.ToJSON()
	want = `["hello","world","lorem","ipsum"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get("[1]").ToJSON()
	want = `"world"`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestSlice2(t *testing.T) {
	var got string
	var want string

	f, err := flex.NewFromJSON(`{"a":"hello","b":[4,5,6]}`)
	if err != nil {
		t.Fatal(err)
	}

	f2 := flex.Append(f.Get(".b"), "hello")

	got = f2.ToJSON()
	want = `[4,5,6,"hello"]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	f.Set("b", flex.Append(f.Get(".b"), "hello"))
	got = f.ToJSON()
	want = `{"a":"hello","b":[4,5,6]}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
