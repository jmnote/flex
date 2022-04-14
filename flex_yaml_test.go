package flex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYAMLString(t *testing.T) {
	var want string
	var got string

	yamlString := `
kind: TestCase
spec:
  containers:
  - name: envar-demo-container
    image: gcr.io/google-samples/node-hello:1.0
    env:
    - name: DEMO_GREETING
      value: "Hello from the environment"
    - name: DEMO_FAREWELL
      value: "Such a sweet sorrow"
`
	f, err := NewFromYAMLString(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	got = FmtToStringDetail(f.Get(".spec"))
	want = `map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers"))
	want = `[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers[0]"))
	want = `map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers[0].env"))
	want = `[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}`
	assert.Equal(t, want, got, "not equal")

	want = `"DEMO_FAREWELL"`
	got = FmtToStringDetail(f.Get(".spec.containers[0].env[1].name"))
	assert.Equal(t, want, got, "not equal")
}

func TestYAMLFile(t *testing.T) {
	var want string
	var got string

	f, err := NewFromYAMLFile("sample.yaml")
	if err != nil {
		t.Fatal(err)
	}

	got = FmtToStringDetail(f.Get(".spec"))
	want = `map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers"))
	want = `[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers[0]"))
	want = `map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}`
	assert.Equal(t, want, got, "not equal")

	got = FmtToStringDetail(f.Get(".spec.containers[0].env"))
	want = `[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}`
	assert.Equal(t, want, got, "not equal")

	want = `"DEMO_FAREWELL"`
	got = FmtToStringDetail(f.Get(".spec.containers[0].env[1].name"))
	assert.Equal(t, want, got, "not equal")
}

func TestYAMLStringTypes(t *testing.T) {
	yamlString := `
nil: ~
bool: true
float64: 3.14
int: 42
int-slice:
- 1
- 2
- 3
- 4
string: hello
string-map:
  a: hello
  b: 42
string-map-string:
  a: hello
  b: world
string-slice:
- a
- b
- c
`
	f, err := NewFromYAMLString(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, f.GetFlex(".nil").Object(), "not equal")
	assert.Equal(t, true, f.GetFlex(".bool").Bool(), "not equal")
	assert.Equal(t, 3.14, f.GetFlex(".float64").Float64(), "not equal")
	assert.Equal(t, 42, f.GetFlex(".int").Int(), "not equal")
	assert.Equal(t, []int{1, 2, 3, 4}, f.GetFlex(".int-slice").IntSlice(), "not equal")
	assert.Equal(t, "hello", f.GetFlex(".string").String(), "not equal")
	assert.Equal(t, map[string]interface{}{"a": "hello", "b": 42}, f.GetFlex(".string-map").StringMap(), "not equal")
	assert.Equal(t, map[string]string{"a": "hello", "b": "world"}, f.GetFlex(".string-map-string").StringMapString(), "not equal")
	assert.Equal(t, []string{"a", "b", "c"}, f.GetFlex(".string-slice").StringSlice(), "not equal")

	assert.Equal(t, nil, f.Get(".nil"), "not equal")
	assert.Equal(t, true, f.GetBool(".bool"), "not equal")
	assert.Equal(t, 3.14, f.GetFloat64(".float64"), "not equal")
	assert.Equal(t, 42, f.GetInt(".int"), "not equal")
	assert.Equal(t, []int{1, 2, 3, 4}, f.GetIntSlice(".int-slice"), "not equal")
	assert.Equal(t, "hello", f.GetString(".string"), "not equal")
	assert.Equal(t, map[string]interface{}{"a": "hello", "b": 42}, f.GetStringMap(".string-map"), "not equal")
	assert.Equal(t, map[string]string{"a": "hello", "b": "world"}, f.GetStringMapString(".string-map-string"), "not equal")
	assert.Equal(t, []string{"a", "b", "c"}, f.GetStringSlice(".string-slice"), "not equal")
}
