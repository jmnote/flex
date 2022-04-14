package flex

import (
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

// get value

func (f *Flex) Get(key string) interface{} {
	return f.GetFlex(key).Object()
}

func (f *Flex) GetBool(key string) bool {
	return f.GetFlex(key).Bool()
}

func (f *Flex) GetFloat64(key string) float64 {
	return f.GetFlex(key).Float64()
}

func (f *Flex) GetInt(key string) int {
	return f.GetFlex(key).Int()
}

func (f *Flex) GetIntSlice(key string) []int {
	return f.GetFlex(key).IntSlice()
}

func (f *Flex) GetSlice(key string) []interface{} {
	return f.GetFlex(key).Slice()
}

func (f *Flex) GetString(key string) string {
	return f.GetFlex(key).String()
}

func (f *Flex) GetStringMap(key string) map[string]interface{} {
	return f.GetFlex(key).StringMap()
}

func (f *Flex) GetStringMapString(key string) map[string]string {
	return f.GetFlex(key).StringMapString()
}

func (f *Flex) GetStringSlice(key string) []string {
	return f.GetFlex(key).StringSlice()
}

// casting to value

func (f *Flex) Object() interface{} {
	return f.object
}

func (f *Flex) Bool() bool {
	return cast.ToBool(f.object)
}

func (f *Flex) Float64() float64 {
	return cast.ToFloat64(f.object)
}

func (f *Flex) Int() int {
	return cast.ToInt(f.object)
}

func (f *Flex) IntSlice() []int {
	return cast.ToIntSlice(f.object)
}

func (f *Flex) Slice() []interface{} {
	return cast.ToSlice(f.object)
}

func (f *Flex) String() string {
	return cast.ToString(f.object)
}

func (f *Flex) StringMap() map[string]interface{} {
	return cast.ToStringMap(f.object)
}

func (f *Flex) StringMapString() map[string]string {
	return cast.ToStringMapString(f.object)
}

func (f *Flex) StringSlice() []string {
	return cast.ToStringSlice(f.object)
}

// get flex

func (f *Flex) GetFlex(key string) *Flex {
	if key == "" || key == "." {
		return f
	}
	val := new(Flex)
	r, _ := regexp.Compile(`^\.[a-zA-Z0-9\-]+`)
	key1 := r.FindString(key)
	if key1 != "" {
		val.object = cast.ToStringMap(f.object)[key1[1:]]
		return val.GetFlex(key[len(key1):])
	}
	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		val.object = cast.ToSlice(f.object)[cast.ToInt(key1[1:len(key1)-1])]
		return val.GetFlex(key[len(key1):])
	}
	return nil
}

// formatting to string

func FmtToString(object interface{}) string {
	return fmt.Sprintf("%v", object)
}

func FmtToStringDetail(object interface{}) string {
	return fmt.Sprintf("%#v", object)
}
