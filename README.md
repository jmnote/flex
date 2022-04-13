# flex

flexible object in golang
* read from yaml, json
* access values by string key
* inspired by viper

## Creating a Flex

In Flex, there are a few ways to create a flex.
The following functions exist:

* `NewFromJSONString(rawString string) : (*Flex, error)`
* `NewFromJSONFile(filepath string) : (*Flex, error)`
* `NewFromJSONBytes(rawBytes []byte) : (*Flex, error)`
* `NewFromYAMLString(rawString string) : (*Flex, error)`
* `NewFromYAMLFile(filepath string) : (*Flex, error)`
* `NewFromYAMLBytes(rawBytes []byte) : (*Flex, error)`

## Getting Values From Flex

In Flex, there are a few ways to get a value depending on the value’s type.
The following methods exist:

* `Get(key string) : interface{}`
* `GetBool(key string) : bool`
* `GetFloat64(key string) : float64`
* `GetInt(key string) : int`
* `GetIntSlice(key string) : []int`
* `GetString(key string) : string`
* `GetStringMap(key string) : map[string]interface{}`
* `GetStringMapString(key string) : map[string]string`
* `GetStringSlice(key string) : []string`
