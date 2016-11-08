package netstat

import (
	"os/exec"
	"runtime"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
        AddPartData([]byte)
}

func GetNetStat(p Profiler) error {
	//"time"
	//proc := new(HeaderType)
	//_ = AddMetadata(proc)
	var net_query string

	switch os := runtime.GOOS; os {
	case "darwin":
		net_query = "select * from process_open_sockets"
	case "linux":
		net_query = "select * from process_open_sockets"
	case "windows":
		net_query = "select * from process_open_sockets"
	}

	p.AddPart("netstat")
	net_cmd := "osqueryi"
	net_list, _ := exec.Command(net_cmd, net_query, "--json").Output()
	p.AddPartData(net_list)
	return nil
}
