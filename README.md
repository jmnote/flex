# flex

flexible object in golang
* reading from yaml(string, file, bytes), json(string, file, bytes)
* accessing values by string key
* chaining

## Creating a Flex

In Flex, there are a few ways to create a flex.
The following functions exist:

* `NewFromJSON(jsonString string) (*Flex, error)`
* `NewFromJSONFile(filepath string) (*Flex, error)`
* `NewFromJSONBytes(jsonBytes []byte) (*Flex, error)`
* `NewFromYAML(yamlString string) (*Flex, error)`
* `NewFromYAMLFile(filepath string) (*Flex, error)`
* `NewFromYAMLBytes(yamlBytes []byte) (*Flex, error)`

## Getting Values From Flex

In Flex, there are a few ways to get a value depending on the value’s type.
The following methods exist:

* `Get(key string) *Flex`
* `GetBool(key string) bool`
* `GetFloat64(key string) float64`
* `GetInt(key string) int`
* `GetIntSlice(key string) []int`
* `GetObject(key string) interface{}`
* `GetObjectMap(key string) map[string]interface{}`
* `GetString(key string) string`
* `GetStringMap(key string) map[string]string`
* `GetStringSlice(key string) []string`

## Converting Values From Flex

* `ToBool() bool`
* `ToFloat64() float64`
* `ToInt() int`
* `ToIntSlice() []int`
* `ToObject() interface{}`
* `ToObjectMap() map[string]interface{}`
* `ToString() string`
* `ToStringMap() map[string]string`
* `ToStringSlice() []string`
* `ToJSON() string`
* `ToYAML() string`

## Setting/Appending Value

* `Set(key string, interface{}) *Flex`
* `Append(f *Flex, objects ...interface{}) *Flex`

