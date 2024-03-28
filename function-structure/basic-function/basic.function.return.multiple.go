package basicFunction

import (
	"errors"
	"fmt"
)

func getDataFromDB(key string) ([]string, error) {
	if key == "foo" {
		return []string{"data1", "data2", "data3"}, nil
	}
	return nil, errors.New("data not found")
}

func GetData(key string) ([]string, error) {
	data, err := getDataFromDB(key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FunctionReturnMultiple() {
	key := "foo"
	data, err := GetData(key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data retrieved successfully:")
		for i, val := range data {
			fmt.Printf("[%d] %s\n", i+1, val)
		}
	}
}
