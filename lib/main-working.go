package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type HeaderType struct {
	DateTime             string `json:"date_time"`
	Instance             string `json:"instance"`
	Severity             string `json:"severity"`
	InstanceIdentityType `json:"instance_identity"`
	Part                 Part `json:"part"`
}

type InstanceIdentityType struct {
	DevPayProductCodes string    `json:"devpayProductCodes"`
	PrivateIp          string    `json:"privateIp"`
	AvailabilityZone   string    `json:"availabilityZone"`
	AccountId          string    `json:"accountId"`
	Version            string    `json:"version"`
	InstanceId         string    `json:"instanceId"`
	BillingProducts    []string  `json:"billingProducts"`
	InstanceType       string    `json:"instanceType"`
	ImageId            string    `json:"imageId"`
	PendingTime        time.Time `json:"pendingTime"`
	Architecture       string    `json:"architecture"`
	KernelId           string    `json:"kernelId"`
	RamdiskId          string    `json:"ramdiskId"`
	Region             string    `json:"region"`
}

type Part struct {
	Part string      `json:"part"`
	Data interface{} `json:"data"`
}

func (t *HeaderType) AddPart(part string) {
	t.Part.Part = part
}

func (t *HeaderType) AddPartData(data []byte) {
	var f interface{}
	json.Unmarshal(data, &f)
	t.Part.Data = f
}

func ProfileDate() string {
	t := time.Now()
	formatedTime := t.Format(time.RFC3339)
	return string(formatedTime)
}

//Add Metdata to HeaderType
func AddMetadata(m *HeaderType) error {
	// Create a EC2Metadata client with additional configuration
	svc := ec2metadata.New(session.New())
	var identitydocument ec2metadata.EC2InstanceIdentityDocument

	m.DateTime = ProfileDate()
	m.Severity = "INFO"
	if svc.Available() {
		identitydocument, _ = svc.GetInstanceIdentityDocument()

		m.InstanceIdentityType.PrivateIp = identitydocument.PrivateIP
		m.InstanceIdentityType.AvailabilityZone = identitydocument.AvailabilityZone
		m.InstanceIdentityType.AccountId = identitydocument.AccountID
		m.InstanceIdentityType.Version = identitydocument.Version
		m.InstanceIdentityType.InstanceId = identitydocument.InstanceID
		//m.InstanceIdentityType.BillingProducts = append(m.InstanceIdentityType.BillingProducts, identitydocument.BillingProducts)
		m.InstanceIdentityType.InstanceType = identitydocument.InstanceType
		m.InstanceIdentityType.ImageId = identitydocument.ImageID
		m.InstanceIdentityType.PendingTime = identitydocument.PendingTime
		m.InstanceIdentityType.Architecture = identitydocument.Architecture
		m.InstanceIdentityType.KernelId = identitydocument.KernelID
		m.InstanceIdentityType.RamdiskId = identitydocument.RamdiskID
		m.InstanceIdentityType.Region = identitydocument.Region

	} else {
		m.DateTime = ProfileDate()
		//m.Instance = "i-123456"
		//m.Severity = "INFO"
		//m.InstanceIdentityType.DevPayProductCodes = "null"
		//m.InstanceIdentityType.PrivateIp = "10.1.1.4"
		//m.InstanceIdentityType.AvailabilityZone = "us-west-2c"
		//m.InstanceIdentityType.AccountId = "113742210699"
		//m.InstanceIdentityType.Version = "2010-08-31"
		//m.InstanceIdentityType.InstanceId = "i-0a1bf6d6"
		//m.InstanceIdentityType.BillingProducts = append(m.InstanceIdentityType.BillingProducts, "bp-63a5400a")
		//m.InstanceIdentityType.InstanceType = "m3.large"
		//m.InstanceIdentityType.ImageId = "ami-b061acd0"
		//m.InstanceIdentityType.PendingTime = 2016-10-11T20:43:07Z
		//m.InstanceIdentityType.Architecture = "x86_64"
		//m.InstanceIdentityType.KernelId = "null"
		//m.InstanceIdentityType.RamdiskId = "null"
		//m.InstanceIdentityType.Region = "us-west-2"
	}
	return nil
}

func Software() interface{} {
	software := new(HeaderType)
	_ = AddMetadata(software)
	var pkg_query string

	switch os := runtime.GOOS; os {
	case "darwin":
		pkg_query = "select * from homebrew_packages"
	case "linux":
		pkg_query = "select * from rpm_packages"
	case "windows":
		pkg_query = "select * from programs"
	}

	software.AddPart("software")
	pkg_cmd := "osqueryi"
	pkg_list, _ := exec.Command(pkg_cmd, pkg_query, "--json").Output()
	software.AddPartData(pkg_list)
	return software
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteProfile(profile []byte, path string) {
	p, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	p.Write(profile)
}

func main() {
	path := "/tmp/agent.log"

	//Get 'software' profile
	software := Software()
	sjson, _ := json.Marshal(software)
	fmt.Printf("%s", sjson)
	WriteProfile(sjson, path)
	procs := getProcs()
	pjson, _ := json.Marshal(procs)
	fmt.Printf("%s", pjson)
}
