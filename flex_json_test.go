package flex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONString(t *testing.T) {
	var want string
	var got string

	jsonString := `{"kind":"TestCase","spec":{"containers":[{"name":"envar-demo-container","image":"gcr.io/google-samples/node-hello:1.0","env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}]}]}}`
	f, err := NewFromJSONString(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f.Get(".spec"))
	want = `map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers"))
	want = `[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers[0]"))
	want = `map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers[0].env"))
	want = `[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}`
	assert.Equal(t, want, got, "not equal")

	want = `"DEMO_FAREWELL"`
	got = FmtStringSharp(f.Get(".spec.containers[0].env[1].name"))
	assert.Equal(t, want, got, "not equal")
}

func TestJSONFile(t *testing.T) {
	var want string
	var got string

	f, err := NewFromJSONFile("sample.json")
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f.Get(".spec"))
	want = `map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers"))
	want = `[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers[0]"))
	want = `map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}`
	assert.Equal(t, want, got, "not equal")

	got = FmtStringSharp(f.Get(".spec.containers[0].env"))
	want = `[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}`
	assert.Equal(t, want, got, "not equal")

	want = `"DEMO_FAREWELL"`
	got = FmtStringSharp(f.Get(".spec.containers[0].env[1].name"))
	assert.Equal(t, want, got, "not equal")
}
