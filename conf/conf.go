package conf

import (
	"fmt"
	"os"
	"strings"

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

func Strings(key string) []string {
	return Configure.Strings(key)
}

func Bool(key string) bool {
	b, _ := Configure.Bool(key)
	return b
}

//destpath=${rpstir2::datadir}/rsyncrepo   --> replace ${rpstir2::datadir}
//-->/root/rpki/data/rsyncrepo --> get /root/rpki/data/rsyncrepo
func VariableString(key string) string {
	if len(key) == 0 || len(String(key)) == 0 {
		return ""
	}
	value := String(key)
	start := strings.Index(value, "${")
	end := strings.Index(value, "}")
	if start >= 0 && end > 0 && start < end {
		//${rpstir2::datadir}/rsyncrepo -->rpstir2::datadir
		replaceKey := string(value[start+len("${") : end])
		if len(replaceKey) == 0 || len(String(replaceKey)) == 0 {
			return value
		}
		//rpstir2::datadir -->get  "/root/rpki/data"
		replaceValue := String(replaceKey)
		prefix := string(value[:start])
		suffix := string(value[end+1:])
		///root/rpki/data/rsyncrepo
		newValue := prefix + replaceValue + suffix
		return newValue
	}
	return ""

}
