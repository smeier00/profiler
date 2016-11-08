package procs

import (
	"os/exec"
	"runtime"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
        AddPartData([]byte)
}

func GetProcs(p Profiler) error {
	var query string

	switch os := runtime.GOOS; os {
	case "darwin":
		query = "select * from processes"
	case "linux":
		query = "select * from processes"
	case "windows":
		query = "select * from processes"
	}

	p.AddPart("procs")
	cmd := "osqueryi"
	proc_list, _ := exec.Command(cmd, query, "--json").Output()
	p.AddPartData(proc_list)
	return nil
}
