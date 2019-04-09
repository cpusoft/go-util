package util

import (
	"fmt"
	"os"

	config "github.com/astaxie/beego/config"
	osutil "github.com/cpusoft/goutil/osutil"
)

var Configure config.Configer

// load configure file
func init() {
	/*
		iniFile := config.NewINIFile(util.GetParentPath() + string(os.PathSeparator) + "conf/slurm.conf")
		Configure = config.NewConfig([]config.Provider{iniFile})
		if err := Configure.Load(); err != nil {
			fmt.Println("conf:", err)
		}
		fmt.Println("conf:", *Configure)
	*/
	var err error
	Configure, err = config.NewConfig("ini", osutil.GetParentPath()+string(os.PathSeparator)+"conf"+string(os.PathSeparator)+"project.conf")
	if err != nil {
		fmt.Println("conf init err: ", err)
	}

}

func String(key string) string {
	return Configure.String(key)
}

func Int(key string) int {
	i, _ := Configure.Int(key)
	return i
}
