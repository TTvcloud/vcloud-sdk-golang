package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/iam"
)

func main() {
	resp, code, _ := iam.DefaultInstance.ListAccessKeys(nil)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
