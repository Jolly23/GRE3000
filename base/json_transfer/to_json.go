package json_transfer

import (
	"encoding/json"
	"fmt"
)

func ToJson(raw interface{}) string {
	b, err := json.Marshal(raw)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
