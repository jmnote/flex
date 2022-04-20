package flex

import (
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
	f, err := NewFromYAMLString(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	var want string
	var got string

	got = FmtStringSharp(f.Get(".spec.containers"))
	want = "&flex.Flex{object:[]interface {}{map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count\"}, map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count-log-1\"}, map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count-log-2\"}}}"
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec").Get(".containers"))
	want = "&flex.Flex{object:[]interface {}{map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count\"}, map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count-log-1\"}, map[string]interface {}{\"image\":\"busybox:1.28\", \"name\":\"count-log-2\"}}}"
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[1]"))
	want = FmtStringSharp(f.Get(".spec").Get(".containers").Get("[1]"))
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec").Get(".containers").GetString("[1].image")
	want = `busybox:1.28`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
