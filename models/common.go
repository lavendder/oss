package models

import "encoding/json"

func StructToMap(obj interface{}) map[string]interface{} {

	m := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)

	return m
}