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
	return f.GetFlex(key).object
}

func (f *Flex) GetBool(key string) bool {
	return cast.ToBool(f.GetFlex(key).object)
}

func (f *Flex) GetFloat64(key string) float64 {
	return cast.ToFloat64(f.GetFlex(key).object)
}

func (f *Flex) GetInt(key string) int {
	return cast.ToInt(f.GetFlex(key).object)
}

func (f *Flex) GetIntSlice(key string) []int {
	return cast.ToIntSlice(f.GetFlex(key).object)
}

func (f *Flex) GetSlice(key string) []interface{} {
	return cast.ToSlice(f.GetFlex(key).object)
}

func (f *Flex) GetString(key string) string {
	return cast.ToString(f.GetFlex(key).object)
}

func (f *Flex) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(f.GetFlex(key).object)
}

func (f *Flex) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(f.GetFlex(key).object)
}

func (f *Flex) GetStringSlice(key string) []string {
	return cast.ToStringSlice(f.GetFlex(key).object)
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

func (f *Flex) String() string {
	return "Flex{" + fmt.Sprintf("%v", f.object) + "}"
}
