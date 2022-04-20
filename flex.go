package flex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

type Flex struct {
	object interface{}
}

// creation

func NewFromObject(object interface{}) *Flex {
	f := new(Flex)
	f.object = object
	return f
}

func NewFromJSONString(rawString string) (*Flex, error) {
	return NewFromJSONBytes([]byte(rawString))
}

func NewFromJSONFile(filepath string) (*Flex, error) {
	rawBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return NewFromYAMLBytes(rawBytes)
}

func NewFromJSONBytes(rawBytes []byte) (*Flex, error) {
	f := new(Flex)
	if err := json.Unmarshal(rawBytes, &f.object); err != nil {
		return nil, err
	}
	return f, nil
}

func NewFromYAMLString(rawString string) (*Flex, error) {
	return NewFromYAMLBytes([]byte(rawString))
}

func NewFromYAMLFile(filepath string) (*Flex, error) {
	rawBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return NewFromYAMLBytes(rawBytes)
}

func NewFromYAMLBytes(rawBytes []byte) (*Flex, error) {
	f := new(Flex)
	if err := yaml.Unmarshal(rawBytes, &f.object); err != nil {
		return nil, err
	}
	return f, nil
}

// slice methods
func MultiFlex(flexes ...*Flex) *Flex {
	var slice []interface{}
	for _, f := range flexes {
		slice = append(slice, f.object)
	}
	return NewFromObject(slice)
}

// get methods
func (f *Flex) Get(key string) *Flex {
	if key == "" || key == "." {
		return f
	}
	val := new(Flex)
	r, _ := regexp.Compile(`^\.[a-zA-Z0-9\-_]+`)
	key1 := r.FindString(key)
	if key1 != "" {
		val.object = cast.ToStringMap(f.object)[key1[1:]]
		return val.Get(key[len(key1):])
	}
	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		val.object = cast.ToSlice(f.object)[cast.ToInt(key1[1:len(key1)-1])]
		return val.Get(key[len(key1):])
	}
	return nil
}

func (f *Flex) Set(key string, object interface{}) *Flex {
	if key == "" || key == "." {
		f.object = object
		return f
	}
	r, _ := regexp.Compile(`^\.[a-zA-Z0-9\-_]+`)
	key1 := r.FindString(key)
	if key1 != "" {
		val := f.Get(key1)
		val.Set(key[len(key1):], object)
		m := cast.ToStringMap(f.object)
		m[key1[1:]] = val.object
		f.object = m
		return f
	}
	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		val := f.Get(key1)
		val.Set(key[len(key1):], object)
		s := cast.ToSlice(f.object)
		s[cast.ToInt(key1[1:len(key1)-1])] = val.object
		f.object = s
		return f
	}
	return f
}

func (f *Flex) GetBool(key string) bool {
	return cast.ToBool(f.Get(key).object)
}

func (f *Flex) GetFloat64(key string) float64 {
	return cast.ToFloat64(f.Get(key).object)
}

func (f *Flex) GetInt(key string) int {
	return cast.ToInt(f.Get(key).object)
}

func (f *Flex) GetIntSlice(key string) []int {
	return cast.ToIntSlice(f.Get(key).object)
}

func (f *Flex) GetObject(key string) interface{} {
	return f.Get(key).object
}

func (f *Flex) GetSlice(key string) []interface{} {
	return cast.ToSlice(f.Get(key).object)
}

func (f *Flex) GetObjectMap(key string) map[string]interface{} {
	return cast.ToStringMap(f.Get(key).object)
}

func (f *Flex) GetString(key string) string {
	return cast.ToString(f.Get(key).object)
}

func (f *Flex) GetStringMap(key string) map[string]string {
	return cast.ToStringMapString(f.Get(key).object)
}

func (f *Flex) GetStringSlice(key string) []string {
	return cast.ToStringSlice(f.Get(key).object)
}

func (f *Flex) GetJSON(key string) string {
	val, _ := json.Marshal(f.Get(key).object)
	return string(val)
}

func (f *Flex) GetYAML(key string) string {
	var b bytes.Buffer
	encoder := yaml.NewEncoder(&b)
	encoder.SetIndent(2)
	encoder.Encode(f.Get(key).object)
	return b.String()
}

func (f *Flex) String() string {
	return "Flex{" + fmt.Sprintf("%v", f.object) + "}"
}
