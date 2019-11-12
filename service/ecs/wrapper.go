package ecs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

//func (p *ECS) RunInstances(input *RunInstancesInput) (*Reservation, error) {
//	query := url.Values{}
//
//	query.Add("AvailabilityZone", *input.AvailabilityZone)
//	query.Add("Count", *input.Count)
//	query.Add("ImageId", *input.ImageId)
//	query.Add("FlavorId", *input.InstanceType)
//	query.Add("KeypairId",*input.KeyName)
//	query.Add("Name", *input.Name) //if we create many instances, system will append minor integers
//	query.Add("SecurityGroupIds", *input.SecurityGroupIds[0])
//	query.Add("RootVolumeSize", *input.RootVolumeSize)
//	query.Add("RootVolumeType", *input.RootVolumeType)
//	query.Add("ProjectName", "default-project")
//	query.Add("SwitchId", *input.SubnetId)
//
//	respBody, status, err := p.Query("RunInstances", query)
//
//	if err != nil {
//		return nil, err
//	}
//	if status != http.StatusOK {
//		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
//	}
//
//	resp := new(Reservation)
//	if err := json.Unmarshal(respBody, resp); err != nil {
//		return nil, errors.Wrap(err, "unmarshal body failed")
//	}
//	return resp, nil
//}


// DescribeInstances API operation for ECS
func (p *ECS) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	query := url.Values{}

	for _,v := range input.Filters{
		query[v.Name] = v.Values
	}

	respBody, status, err := p.Query("GetInstances", query)

	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := new(DescribeInstancesOutput)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal body failed")
	}

	return resp, nil
}




