package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	instance := vod.NewInstance()
	ret, _ := instance.GetVideoPlayAuth([]string{"v0282cd70000blsd0bkthbi41ag6kpcg"}, []string{}, []string{})
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetVideoPlayAuth([]string{"v0282cd70000blsd0bkthbi41ag6kpcg"}, []string{"evideo"}, []string{})
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
