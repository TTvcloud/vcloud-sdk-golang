package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/service/iam"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config

	query := url.Values{}
	query.Set("Limit", "3")

	resp, code, _ := iam.NewInstance().ListUsers(query)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
