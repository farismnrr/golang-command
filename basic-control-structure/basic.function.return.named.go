package basicControl

import "fmt"

func GetDataFromMap(dataDict map[string]string, key string) (data string, err error) {
	if _, ok := dataDict[key]; !ok {
		return key, fmt.Errorf("key %s not found", key)
	}
	data = dataDict[key]

	return
}

func ControlFuncReturnNamed() {
	dataDict := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	key := "key3"
	data, err := GetDataFromMap(dataDict, key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data retrieved successfully:", data)
	}
}
