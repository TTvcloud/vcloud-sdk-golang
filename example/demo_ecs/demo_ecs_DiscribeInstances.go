package main

import (
	"fmt"
	"vcloud-sdk-golang/base"
	"vcloud-sdk-golang/bytedance"
	"vcloud-sdk-golang/service/ecs"
)

func main() {
	session := bytedance.Session{
		Config: bytedance.Config{
			Region:base.RegionCnNorth2,
		},
	}

	client := ecs.New(session)

	// call below method if you dont set ak and sk in ～/.vcloud/config
	//client.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	request := ecs.DescribeInstancesInput{
		Filters: []*ecs.Filter{
			{
				Name:"InstanceIds",
				Values:[]string{"Vm6597779701332340736"},
			},
			{
				Name:"ProjectName",
				Values:[]string{"default"},
			},
		},
	}

	res, err := client.DescribeInstances(&request)

	if err != nil{
		fmt.Println("Failed to Describe Instances", err)
	}else{
		for _,v := range res.Reservation.Instances{
			fmt.Println("Active Instance Name: ", v.Name)
			fmt.Println("Active Instance ID: ", v.Id)
		}
		fmt.Println("Successfully Described Instances")
	}

}
