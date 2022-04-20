package flex

import (
	"testing"
)

func TestJSONString(t *testing.T) {
	var got string
	var want string

	jsonString := `{"kind":"TestCase","spec":{"containers":[{"name":"envar-demo-container","image":"gcr.io/google-samples/node-hello:1.0","env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}]}]}}`
	f, err := NewFromJSONString(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f.Get(".spec"))
	want = `&flex.Flex{object:map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers"))
	want = `&flex.Flex{object:[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0]"))
	want = `&flex.Flex{object:map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0].env"))
	want = `&flex.Flex{object:[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0].env[1].name"))
	want = `&flex.Flex{object:"DEMO_FAREWELL"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestJSONFile(t *testing.T) {
	var got string
	var want string

	f, err := NewFromJSONFile("sample.json")
	if err != nil {
		t.Fatal(err)
	}

	got = FmtStringSharp(f.Get(".spec"))
	want = `&flex.Flex{object:map[string]interface {}{"containers":[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers"))
	want = `&flex.Flex{object:[]interface {}{map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0]"))
	want = `&flex.Flex{object:map[string]interface {}{"env":[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}, "image":"gcr.io/google-samples/node-hello:1.0", "name":"envar-demo-container"}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0].env"))
	want = `&flex.Flex{object:[]interface {}{map[string]interface {}{"name":"DEMO_GREETING", "value":"Hello from the environment"}, map[string]interface {}{"name":"DEMO_FAREWELL", "value":"Such a sweet sorrow"}}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = FmtStringSharp(f.Get(".spec.containers[0].env[1].name"))
	want = `&flex.Flex{object:"DEMO_FAREWELL"}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestJSONToYAML(t *testing.T) {
	var got string
	var want string

	f, err := NewFromJSONFile("sample.json")
	if err != nil {
		t.Fatal(err)
	}

	got = f.GetYAML("")
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
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.GetYAML(".spec.containers")
	want = `- env:
    - name: DEMO_GREETING
      value: Hello from the environment
    - name: DEMO_FAREWELL
      value: Such a sweet sorrow
  image: gcr.io/google-samples/node-hello:1.0
  name: envar-demo-container
`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}

func TestJSONToJSON(t *testing.T) {
	var want string
	var got string

	f, err := NewFromJSONFile("sample.json")
	if err != nil {
		t.Fatal(err)
	}

	got = f.GetJSON("")
	want = `{"kind":"TestCase","spec":{"containers":[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]}}`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.GetJSON(".spec.containers")
	want = `[{"env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}],"image":"gcr.io/google-samples/node-hello:1.0","name":"envar-demo-container"}]`
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}
}
