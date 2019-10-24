package osutil

import (
	"container/list"
	"io/ioutil"
	"os"
	"os/exec"
	path "path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	belogs "github.com/astaxie/beego/logs"
	hashutil "github.com/cpusoft/goutil/hashutil"
)

func IsExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func IsDir(file string) (bool, error) {
	s, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

func IsFile(file string) (bool, error) {
	s, err := IsDir(file)
	return !s, err
}

// make path.Base() using in windows,
func Base(p string) string {
	p = strings.Replace(p, "\\", "/", -1)
	return path.Base(p)
}

// make path.Split using in win
func Split(p string) (dir, file string) {
	p = strings.Replace(p, "\\", "/", -1)
	return path.Split(p)
}

// path.Ext() using in windows,
func Ext(p string) string {
	p = strings.Replace(p, "\\", "/", -1)
	return path.Ext(p)
}
func GetParentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	dirs := strings.Split(path, string(os.PathSeparator))
	index := len(dirs)
	if len(dirs) > 2 {
		index = len(dirs) - 2
	}
	ret := strings.Join(dirs[:index], string(os.PathSeparator))
	return ret
}

// will deprecated, will use GetAllFilesBySuffixs()
func GetAllFilesInDirectoryBySuffixs(directory string, suffixs map[string]string) *list.List {

	absolutePath, _ := filepath.Abs(directory)
	listStr := list.New()
	filepath.Walk(absolutePath, func(filename string, fi os.FileInfo, err error) error {
		if err != nil || len(filename) == 0 || nil == fi {
			return err
		}
		if !fi.IsDir() {
			suffix := Ext(filename)
			//fmt.Println(suffix)
			if _, ok := suffixs[suffix]; ok {
				listStr.PushBack(filename)
			}
		}
		return nil
	})
	return listStr
}
func GetAllFilesBySuffixs(directory string, suffixs map[string]string) ([]string, error) {

	absolutePath, _ := filepath.Abs(directory)
	files := make([]string, 0)
	filepath.Walk(absolutePath, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil || len(fileName) == 0 || nil == fi {
			belogs.Debug("GetAllFilesBySuffixs():filepath.Walk(): err:", err)
			return err
		}
		if !fi.IsDir() {
			suffix := Ext(fileName)
			if _, ok := suffixs[suffix]; ok {
				files = append(files, fileName)
			}
		}
		return nil
	})
	return files, nil
}

func GetFilesInDir(directory string, suffixs map[string]string) ([]string, error) {
	files := make([]string, 0, 10)
	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if file.IsDir() { // 忽略目录
			continue
		}
		suffix := Ext(file.Name())
		if _, ok := suffixs[suffix]; ok {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

type FileStat struct {
	FilePath string    `json:"filePath"`
	FileName string    `json:"fileName"`
	ModeTime time.Time `json:"modeTime"`
	Size     int64     `json:"size"`
	Hash256  string    `json:"hash256"`
}

func GetAllFileStatsBySuffixs(directory string, suffixs map[string]string) ([]FileStat, error) {

	absolutePath, _ := filepath.Abs(directory)
	fileStats := make([]FileStat, 0)
	filepath.Walk(absolutePath, func(path string, fi os.FileInfo, err error) error {
		if err != nil || len(path) == 0 || nil == fi {
			belogs.Debug("GetAllFileStatsBySuffixs():filepath.Walk(): err:", err)
			return err
		}
		if !fi.IsDir() {

			suffix := Ext(path)
			if _, ok := suffixs[suffix]; ok {
				fileStat := FileStat{}
				fileStat.FilePath, _ = Split(path)
				fileStat.FileName = fi.Name()
				fileStat.ModeTime = fi.ModTime()
				fileStat.Size = fi.Size()
				fileStat.Hash256, _ = hashutil.Sha256File(JoinPathFile(fileStat.FilePath, fileStat.FileName))
				fileStats = append(fileStats, fileStat)
			}
		}
		return nil
	})
	return fileStats, nil

}

func GetFilePathAndFileName(fileAllPath string) (filePath string, fileName string) {
	i := strings.LastIndex(fileAllPath, string(os.PathSeparator))
	return fileAllPath[:i+1], fileAllPath[i+1:]
}

func GetNewLineSep() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	case "linux":
		return "\n"
	default:
		return "\n"

	}
}

func GetPathSeparator() string {
	return string(os.PathSeparator)
}

func JoinPathFile(pathName, fileName string) string {
	if !strings.HasSuffix(pathName, string(os.PathSeparator)) {
		pathName = pathName + string(os.PathSeparator)
	}
	return pathName + fileName
}

func CloseAndRemoveFile(file *os.File) error {
	if file == nil {
		return nil
	}
	s, err := IsExists(file.Name())
	if err != nil {
		belogs.Error("parseMftModel():IsExists:err: ", file.Name(), err)
		return err
	}
	if !s {
		return nil
	}

	err = file.Close()
	if err != nil {
		belogs.Error("parseMftModel():cerfile.Close():err: ", file.Name(), err)
		return err
	}
	err = os.Remove(file.Name())
	if err != nil {
		belogs.Error("parseMftModel():os.Remove:err:", file.Name(), err)
		return nil
	}
	return nil
}
