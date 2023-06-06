package utils

import (
	"encoding/json"
	"fmt"
)

func GetStructJSON(value interface{}) string {
	result, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}
