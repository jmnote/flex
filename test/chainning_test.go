package test

import (
	"github.com/jmnote/flex"
	"testing"
)

func TestChaining(t *testing.T) {
	yamlString := `
kind: TestCase
spec:
  containers:
  - name: count
    image: busybox:1.28
  - name: count-log-1
    image: busybox:1.28
  - name: count-log-2
    image: busybox:1.28

`
	f, err := flex.NewFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	var want string
	var got string

	got = f.Get(".spec.containers").ToJSON()
	want = `[{"image":"busybox:1.28","name":"count"},{"image":"busybox:1.28","name":"count-log-1"},{"image":"busybox:1.28","name":"count-log-2"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec").Get(".containers").ToJSON()
	want = `[{"image":"busybox:1.28","name":"count"},{"image":"busybox:1.28","name":"count-log-1"},{"image":"busybox:1.28","name":"count-log-2"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[1]").ToJSON()
	want = f.Get(".spec").Get(".containers").Get("[1]").ToJSON()
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec").Get(".containers").GetString("[1].image")
	want = `busybox:1.28`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
