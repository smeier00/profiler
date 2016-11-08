package os

import (
	"os/exec"
	"runtime"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
        AddPartData([]byte)
}

func GetOs(p Profiler) error {
	var query string

	switch os := runtime.GOOS; os {
	case "darwin":
		query = "select * from system_info"
	case "linux":
		query = "select * from system_info"
	case "windows":
		query = "select * from system_info"
	}

	p.AddPart("os")
	cmd := "osqueryi"
	os, _ := exec.Command(cmd, query, "--json").Output()
	p.AddPartData(os)
	return nil
}
