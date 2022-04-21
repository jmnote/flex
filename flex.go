package flex

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"os"
	"regexp"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

type Flex struct {
	Object interface{}
}

// creation

func New() *Flex {
	f := new(Flex)
	return f
}

func NewFromObject(object interface{}) *Flex {
	f := new(Flex)
	f.Object = object
	return f
}

func NewFromJSON(jsonString string) (*Flex, error) {
	return NewFromJSONBytes([]byte(jsonString))
}

func NewFromJSONFile(filepath string) (*Flex, error) {
	jsonBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return NewFromYAMLBytes(jsonBytes)
}

func NewFromJSONBytes(jsonBytes []byte) (*Flex, error) {
	f := new(Flex)
	if err := json.Unmarshal(jsonBytes, &f.Object); err != nil {
		return nil, err
	}
	return f, nil
}

func NewFromYAML(yamlString string) (*Flex, error) {
	return NewFromYAMLBytes([]byte(yamlString))
}

func NewFromYAMLFile(filepath string) (*Flex, error) {
	yamlBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return NewFromYAMLBytes(yamlBytes)
}

func NewFromYAMLBytes(yamlBytes []byte) (*Flex, error) {
	f := new(Flex)
	if err := yaml.Unmarshal(yamlBytes, &f.Object); err != nil {
		return nil, err
	}
	return f, nil
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
		val.Object = cast.ToStringMap(f.Object)[key1[1:]]
		return val.Get(key[len(key1):])
	}
	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		val.Object = cast.ToSlice(f.Object)[cast.ToInt(key1[1:len(key1)-1])]
		return val.Get(key[len(key1):])
	}
	r, _ = regexp.Compile(`^\[([0-9]*):([0-9]*)\]`)
	matches := r.FindStringSubmatch(key)
	if len(matches) == 3 {
		slice := cast.ToSlice(f.Object)
		start := 0
		end := len(slice)
		if matches[1] != "" {
			start = cast.ToInt(matches[1])
		}
		if matches[2] != "" {
			end = cast.ToInt(matches[2])
		}
		val.Object = slice[start:end]
		return val.Get(key[len(matches[0]):])
	}
	return nil
}

func (f *Flex) Set(key string, object interface{}) *Flex {
	if key == "" || key == "." {
		f.Object = object
		return f
	}
	r, _ := regexp.Compile(`^\.[a-zA-Z0-9\-_]+`)
	key1 := r.FindString(key)
	if key1 != "" {
		val := f.Get(key1)
		val.Set(key[len(key1):], object)
		m := cast.ToStringMap(f.Object)
		m[key1[1:]] = val.Object
		f.Object = m
		return f
	}
	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		val := f.Get(key1)
		val.Set(key[len(key1):], object)
		s := cast.ToSlice(f.Object)
		s[cast.ToInt(key1[1:len(key1)-1])] = val.Object
		f.Object = s
		return f
	}
	return f
}

func Append(f *Flex, objects ...interface{}) *Flex {
	slice := f.ToSlice()
	slice = append(slice, objects...)
	f.Object = slice
	return f
}

func (f *Flex) GetBool(key string) bool {
	return f.Get(key).ToBool()
}

func (f *Flex) ToBool() bool {
	return cast.ToBool(f.Object)
}

func (f *Flex) GetFloat64(key string) float64 {
	return f.Get(key).ToFloat64()
}

func (f *Flex) ToFloat64() float64 {
	return cast.ToFloat64(f.Object)
}

func (f *Flex) GetInt(key string) int {
	return f.Get(key).ToInt()
}

func (f *Flex) ToInt() int {
	return cast.ToInt(f.Object)
}

func (f *Flex) GetIntSlice(key string) []int {
	return f.Get(key).ToIntSlice()
}

func (f *Flex) ToIntSlice() []int {
	return cast.ToIntSlice(f.Object)
}

func (f *Flex) GetObject(key string) interface{} {
	return f.Get(key).Object
}

func (f *Flex) ToObject() interface{} {
	return f.Object
}

func (f *Flex) GetSlice(key string) []interface{} {
	return f.Get(key).ToSlice()
}

func (f *Flex) ToSlice() []interface{} {
	var slice []interface{}
	switch t := f.Object.(type) {
	case []interface{}:
		slice = append(slice, t...)
	case []string:
		for _, value := range t {
			slice = append(slice, value)
		}
	case []int:
		for _, value := range t {
			slice = append(slice, value)
		}
	}
	return slice
}

func (f *Flex) GetObjectMap(key string) map[string]interface{} {
	return f.Get(key).ToObjectMap()
}

func (f *Flex) ToObjectMap() map[string]interface{} {
	return cast.ToStringMap(f.Object)
}

func (f *Flex) GetString(key string) string {
	return f.Get(key).ToString()
}

func (f *Flex) ToString() string {
	return cast.ToString(f.Object)
}

func (f *Flex) GetStringMap(key string) map[string]string {
	return f.Get(key).ToStringMap()
}

func (f *Flex) ToStringMap() map[string]string {
	return cast.ToStringMapString(f.Object)
}

func (f *Flex) GetStringSlice(key string) []string {
	return f.Get(key).ToStringSlice()
}

func (f *Flex) ToStringSlice() []string {
	return cast.ToStringSlice(f.Object)
}

func (f *Flex) ToJSON() string {
	val, _ := json.Marshal(f.Object)
	return string(val)
}

func (f *Flex) ToYAML() string {
	var b bytes.Buffer
	encoder := yaml.NewEncoder(&b)
	encoder.SetIndent(2)
	encoder.Encode(f.Object)
	return b.String()
}

func (f *Flex) String() string {
	return f.ToJSON()
}
