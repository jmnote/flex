package test

import (
	"github.com/jmnote/flex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYAMLString(t *testing.T) {
	var got string
	var want string

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
	f, err := flex.NewFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	got = f.Get(".spec").ToJSON()
	want = `{"containers":[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers").ToJSON()
	want = `[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0]").ToJSON()
	want = `{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0].env").ToJSON()
	want = `[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0].env[1].name").ToJSON()
	want = `"DEMO_FAREWELL"`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestYAMLFile(t *testing.T) {
	var want string
	var got string

	f, err := flex.NewFromYAMLFile("sample.yaml")
	if err != nil {
		t.Fatal(err)
	}

	got = f.Get(".spec").ToJSON()
	want = `{"containers":[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers").ToJSON()
	want = `[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0]").ToJSON()
	want = `{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0].env").ToJSON()
	want = `[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers[0].env[1].name").ToJSON()
	want = `"DEMO_FAREWELL"`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
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
object-map:
  a: hello
  b: 42
string-map:
  a: hello
  b: world
string-slice:
- a
- b
- c
`
	f, err := flex.NewFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, f.GetObject(".nil"), "")
	assert.Equal(t, true, f.GetBool(".bool"), "")
	assert.Equal(t, 3.14, f.GetFloat64(".float64"), "")
	assert.Equal(t, 42, f.GetInt(".int"), "")
	assert.Equal(t, []int{1, 2, 3, 4}, f.GetIntSlice(".int-slice"), "")
	assert.Equal(t, "hello", f.GetString(".string"), "")
	assert.Equal(t, map[string]interface{}{"a": "hello", "b": 42}, f.GetObjectMap(".object-map"), "")
	assert.Equal(t, map[string]string{"a": "hello", "b": "world"}, f.GetStringMap(".string-map"), "")
	assert.Equal(t, []string{"a", "b", "c"}, f.GetStringSlice(".string-slice"), "")
}

func TestYAMLToYAML(t *testing.T) {
	var want string
	var got string

	f, err := flex.NewFromYAMLFile("sample.yaml")
	if err != nil {
		t.Fatal(err)
	}

	got = f.ToYAML()
	want = `kind: TestCase
spec:
  containers:
    - env:
        - name: DEMO_GREETING
          value: Hello from the environment
        - name: DEMO_FAREWELL
          value: Such a sweet sorrow
      image: gcr.io/google-samples/node-hello:1.0
      name: envar-demo-container
`
	assert.Equal(t, want, got, "")

	got = f.Get(".spec.containers").ToYAML()
	want = `- env:
    - name: DEMO_GREETING
      value: Hello from the environment
    - name: DEMO_FAREWELL
      value: Such a sweet sorrow
  image: gcr.io/google-samples/node-hello:1.0
  name: envar-demo-container
`
	assert.Equal(t, want, got, "")
}

func TestYAMLToJSON(t *testing.T) {
	var want string
	var got string

	f, err := flex.NewFromYAMLFile("sample.yaml")
	if err != nil {
		t.Fatal(err)
	}

	got = f.ToJSON()
	want = `{"kind":"TestCase","spec":{"containers":[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers").ToJSON()
	want = `[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
