package flex

import (
	"github.com/stretchr/testify/assert"
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
	want = `[]interface {}{map[string]interface {}{"image":"busybox:1.28", "name":"count"}, map[string]interface {}{"image":"busybox:1.28", "name":"count-log-1"}, map[string]interface {}{"image":"busybox:1.28", "name":"count-log-2"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.GetFlex(".spec").Get(".containers"))
	want = `[]interface {}{map[string]interface {}{"image":"busybox:1.28", "name":"count"}, map[string]interface {}{"image":"busybox:1.28", "name":"count-log-1"}, map[string]interface {}{"image":"busybox:1.28", "name":"count-log-2"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers[1]"))
	want = FmtStringSharp(f.GetFlex(".spec").GetFlex(".containers").Get("[1]"))
	assert.Equal(t, want, got, "not equal")

	got = f.GetFlex(".spec").GetFlex(".containers").GetString("[1].image")
	want = `busybox:1.28`
	assert.Equal(t, want, got, "not equal")

}
