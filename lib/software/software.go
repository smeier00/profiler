package software

import (
	//"fmt"
	"os/exec"
	"runtime"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
        AddPartData([]byte)
}

func GetSoftware(p Profiler) error  {
	//software = new(HeaderType)
        //	_ = AddMetadata(software)
	var pkg_query string
	var pkg_list []byte

	switch os := runtime.GOOS; os {
	case "darwin":
		pkg_query = "select * from homebrew_packages"
	case "linux":
		pkg_query = "select * from rpm_packages"
	case "windows":
		pkg_query = "select * from programs"
	}

	p.AddPart("software")
	pkg_cmd := "osqueryi"
	pkg_list, _ = exec.Command(pkg_cmd, pkg_query, "--json").Output()
        
	p.AddPartData(pkg_list)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
