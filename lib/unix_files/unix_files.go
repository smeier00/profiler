package unix_files

import (
	"os/exec"
	"runtime"
        "encoding/json"
        "fmt"
)

type Profiler interface {
        AddMetadata([]byte)
        AddPart(string)
        AddPartData([]byte)
}

type listtype []map[string]string

func GetUnixFiles(p Profiler) error {
	var query string

	switch os := runtime.GOOS; os {
	case "darwin":
		query = "select * from users"
	case "linux":
		query = "select * from users"
	case "windows":
		query = "select * from users"
	}

	p.AddPart("unix_files")
	cmd := "osqueryi"
	user_list, _ := exec.Command(cmd, query, "--json").Output()
        //Unmarshal response, loop and write
        var list listtype 
        json.Unmarshal(user_list, &list)
        fmt.Println(list[1])
        for l := range list {
          //fmt.Printf(l)
          pjson, _ := json.Marshal(l)
          p.AddPartData(pjson)
          //fmt.Println(string(pjson))
        }

	//p.AddPartData(user_list)
	return nil
}
