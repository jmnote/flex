package flex

import (
	"encoding/json"
	"os"
	"regexp"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

type Flex struct {
	rawString string
	data      interface{}
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
	f.rawString = string(rawBytes)

	if err := json.Unmarshal(rawBytes, &f.data); err != nil {
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
	f.rawString = string(rawBytes)

	if err := yaml.Unmarshal(rawBytes, &f.data); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *Flex) GetBool(key string) bool {
	return cast.ToBool(f.Get(key))
}

func (f *Flex) GetFloat64(key string) float64 {
	return cast.ToFloat64(f.Get(key))
}

func (f *Flex) GetInt(key string) int {
	return cast.ToInt(f.Get(key))
}

func (f *Flex) GetIntSlice(key string) []int {
	return cast.ToIntSlice(f.Get(key))
}

func (f *Flex) GetString(key string) string {
	return cast.ToString(f.Get(key))
}

func (f *Flex) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(f.Get(key))
}

func (f *Flex) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(f.Get(key))
}

func (f *Flex) GetStringSlice(key string) []string {
	return cast.ToStringSlice(f.Get(key))
}

func (f *Flex) Get(key string) interface{} {
	return getDeep(key, f.data)
}

func getDeep(key string, object interface{}) interface{} {
	if key == "" || key == "." {
		return object
	}
	r, _ := regexp.Compile(`^\.[a-zA-Z0-9\-]+`)
	key1 := r.FindString(key)
	if key1 != "" {
		realKey1 := key1[1:]
		newKey := key[len(key1):]
		newObject := cast.ToStringMap(object)[realKey1]
		return getDeep(newKey, newObject)
	}

	r, _ = regexp.Compile(`^\[[0-9]+\]`)
	key1 = r.FindString(key)
	if key1 != "" {
		realKey1 := cast.ToInt(key1[1 : len(key1)-1])
		newKey := key[len(key1):]
		newObject := cast.ToSlice(object)[realKey1]
		return getDeep(newKey, newObject)
	}
	return nil
}
