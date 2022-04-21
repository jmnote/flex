package test

import (
	"github.com/jmnote/flex"
	"testing"
)

func TestJSONString(t *testing.T) {
	var got string
	var want string

	jsonString := `{"kind":"TestCase","spec":{"containers":[{"name":"envar-demo-container","image":"gcr.io/google-samples/node-hello:1.0","env":[{"name":"DEMO_GREETING","value":"Hello from the environment"},{"name":"DEMO_FAREWELL","value":"Such a sweet sorrow"}]}]}}`
	f, err := flex.NewFromJSON(jsonString)
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

func TestJSONFile(t *testing.T) {
	var got string
	var want string

	f, err := flex.NewFromJSONFile("sample.json")
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

func TestJSONToYAML(t *testing.T) {
	var got string
	var want string

	f, err := flex.NewFromJSONFile("sample.json")
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
	if want != got {
		t.Fatal("\nwant =>", want, "\ngot  =>", got)
	}

	got = f.Get(".spec.containers").ToYAML()
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

	f, err := flex.NewFromJSONFile("sample.json")
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
