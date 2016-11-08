package main

import (
	//"encoding/json"
	//"fmt"
	//"github.com/aws/aws-sdk-go/aws/ec2metadata"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"os"
	//"os/exec"
	"profiler/lib/metadata"
        "profiler/lib/netstat"
	"profiler/lib/os"
	"profiler/lib/procs"
	"profiler/lib/profiler"
	"profiler/lib/software"
	"profiler/lib/unix_files"
	//"runtime"
	//"time"
)

func main() {
	path := "/tmp/agent.log"

        //#Log all parts netstat, os, metadata, software, proc, unixFiles. 
        //netstat
        netstat_profile := profiler.New()
        _ = metadata.GetMetadata(netstat_profile)
        _ = netstat.GetNetStat(netstat_profile)
        netstat_profile.Print()

        //os
        os_profile := profiler.New()
        _ = metadata.GetMetadata(os_profile)
        _ = os.GetOs(os_profile)
        os_profile.Print()


        //Sofware 
	software_profile := profiler.New()
	_ = metadata.GetMetadata(software_profile)
	_ = software.GetSoftware(software_profile)
        software_profile.Print()
       

        //Procs
        procs_profile := profiler.New() 
        _ = metadata.GetMetadata(procs_profile)
	_ = procs.GetProcs(procs_profile)
	procs_profile.Print()

        //Procs
        unix_files_profile := profiler.New() 
        _ = metadata.GetMetadata(unix_files_profile)
	_ = unix_files.GetUnixFiles(unix_files_profile)
	unix_files_profile.Print()
        unix_files_profile.WriteProfile(path)

}
