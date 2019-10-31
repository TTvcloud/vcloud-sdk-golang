package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	ret, _ := vod.NewInstance().GetVideoPlayAuth([]string{"v0282cd70000blsd0bkthbi41ag6kpcg"}, []string{}, []string{})
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
}
