package profiler

import (
	"encoding/json"
	"fmt"
	"time"
        "os"
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

func (p *HeaderType) AddPart(part string) {
	p.Part.Part = part
}

func (p *HeaderType) AddPartData(data []byte) {
	var f interface{}
	json.Unmarshal(data, &f)
	p.Part.Data = f
}

func ProfileDate() string {
	t := time.Now()
	formatedTime := t.Format(time.RFC3339)
	return string(formatedTime)
}

//Add Metdata to HeaderType
func (p *HeaderType) AddMetadata(meta []byte) {
	var setmeta HeaderType
	json.Unmarshal(meta, &setmeta)
	p.DateTime = setmeta.DateTime
	p.Severity = setmeta.Severity
	p.Instance = setmeta.Instance
}

func (p *HeaderType) Print() {
	pjson, _ := json.Marshal(p)
	fmt.Println(string(pjson))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func (p *HeaderType) WriteProfile(profile []byte, path string) {
func (p *HeaderType) WriteProfile(path string) {
        //path := "/tmp/agent.log"
        pjson, _ := json.Marshal(p)
	pfile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	pfile.Write(pjson)
}

func New() *HeaderType {
	profile := new(HeaderType)
	return profile
}

//func main() {
//	path := "/tmp/agent.log"
//
//	//Get 'software' profile
//	software := Software()
//	sjson, _ := json.Marshal(software)
//	fmt.Printf("%s", sjson)
//	WriteProfile(sjson, path)
//	procs := getProcs()
//	pjson, _ := json.Marshal(procs)
//	fmt.Printf("%s", pjson)
//}
