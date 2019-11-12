package ecs

//type RunInstancesInput struct {
//
//	// AccountID can be found in https://vconsole.bytedance.net/user/basics/
//	// AccountId *string `type:"string"`
//
//	// The Availability Zone in which to create the Capacity Reservation.
//	AvailabilityZone *string `type:"string"`
//
//	// The number of instance
//	Count *string `type:"string"`
//
//	// The ID of the image. An image ID is required to launch an instance and must be
//	// specified here or in a launch template.
//	ImageId *string `type:"string"`
//
//	InstanceType *string `type:"string" enum:"InstanceType"`
//
//	// If you do not specify a key pair, you can't connect to the instance unless
//	// you choose an AMI that is configured to allow users another way to log in.
//	KeyName *string `type:"string"`
//
//	// The name of instance. If you create more than one machine, the minor number
//	// will be appended to the name
//	Name *string `type:"string"`
//
//	// RoleId can be found in https://vconsole.bytedance.net/user/basics/
//	// RoleId *string `type:"string"`
//
//	// TODO: EBS is not supported
//	// The volume is attached to instance
//	RootVolumeSize *string `type:"string"`
//
//	RootVolumeType *string `type:"string"`
//
//	// If you specify a network interface, you must specify any security groups
//	// as part of the network interface.
//	SecurityGroupIds []*string `locationName:"SecurityGroupId" locationNameList:"SecurityGroupId" type:"list"`
//
//	// If you specify a network interface, you must specify any subnets as part
//	// of the network interface.
//	SubnetId *string `type:"string"`
//}

type DescribeInstancesInput struct {

	// The filters.
	//
	//    * availability-zone - The Availability Zone of the instance.
	//
	//    * image-id - The ID of the image used to launch the instance.
	//
	//    * InstanceIds - The IDs of the instance. (for example, {"VmXXX0736", "VmXXX0835"})
	//
	//    * instance-type - The type of instance (for example, t2.micro).
	//
	//	  * ProjectName - The name of the project (for example, default).
	//
	Filters []*Filter
}

type Filter struct {
	// The name of the filter. Filter names are case-sensitive.
	Name string

	// The filter values. Filter values are case-sensitive.
	Values []string
}

type DescribeInstancesOutput struct {

	Metadata ResponseMetadata `json:"ResponseMetadata"`

	// The Reservation
	Reservation Reservation  `json:"Result"`
}

type ResponseMetadata struct{

	// RequestId is for debugging purpose
	RequestId string

	// The action to call downstream service, (for example, GetInstances)
	Action string

	// The service name, (for example, ecs)
	Service string

	// Region for data center, (for example, cn-north-2)
	Region string
}

type Reservation struct {
	// The instances whose status are active
	Instances []Instance

	// The ID of the reservation. TODO: it should be provided by downstream service
	ReservationId string
}

type Instance struct {
	// The serial number of VM (for example, Vm65977797013XXXXXXX)
	Id string

	// The time instance got created
	CreatedAt string

	// The time instance got updated
	UpdateAt string

	// The instance image, which is used to restore VM
	Image Image

	// The instance type, (for example, ecs.s2.2xlarge16, general compute)
	Flavor Flavor

	// The name of instance
	Name string

	Region string

	// The available zone of region
	AvailableZone string

	// Private IP
	Ip string

	Status string
}

type Flavor struct {
	Id string
	CreatedAt string
	UpdateAt string

	// For example, ecs.s2.2xlarge16
	Name string

	Region string

	// The number of CPU
	Cpu int

	// The number of GPU
	Gpu int

	// The size of memory
	Memory int

	// The type of network, (for example, classic)
	NetworkType string

	// for example, FlavorType = general
	FlavorType string
	FlavorTypeDescription string

	// for example, public or Internal
	Visibility string
}

type Image struct {

	Id string

	CreatedAt string

	UpdateAt string

	Name string

	Region string

	// for example, 2ac7ba0e-8c08-41a3-8e32-6b5d6910bd16
	SysRef string

	DataRef string

	// for example, Debian
	OsName string

	// for example, public or Internal
	Visibility string

	// follow each tenant's project
	ProjectId int
}

type Port struct {

	Id string

	CreatedAt string

	UpdateAt string

	Type string

	// the id created for subnet
	SwitchId string
	FixedIp string
	Mac string
	InstanceId string
	NicIndex int

	// the id of VPC
	VpcId string

	VpcName string

	// security groups to control access
	BindingSecurityGroups []string
}



