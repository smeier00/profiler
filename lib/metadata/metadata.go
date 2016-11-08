package metadata

import (
	//"encoding/json"
	//"fmt"
	//"github.com/aws/aws-sdk-go/aws/ec2metadata"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"os"
	//"os/exec"
	//"runtime"
	"time"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
}

func ProfileDate() string {
        t := time.Now()
        formatedTime := t.Format(time.RFC3339)
        return string(formatedTime)
}

//Add Metdata to HeaderType
func GetMetadata(m Profiler) error {
        // Create a EC2Metadata client with additional configuration
        //svc := ec2metadata.New(session.New())
        //var identitydocument ec2metadata.EC2InstanceIdentityDocument
        var test = []byte(`{
                     "DateTime"  : "1234",
                     "Instance"  : "i-12345",
                     "Severity" :  "INFO"
                    }`)

        m.AddMetadata(test)
        return nil
}

