package httpserver

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	belogs "github.com/astaxie/beego/logs"
	"github.com/cpusoft/go-json-rest/rest"
)

// result:ok/fail
type HttpResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

// setup Http Server, listen on port
func ListenAndServe(port string, router *rest.App) {

	api := rest.NewApi()
	MyAccessProdStack := rest.AccessProdStack
	MyAccessProdStack[0] = &rest.AccessLogApacheMiddleware{
		Logger: belogs.GetLogger("access"),
		Format: rest.CombinedLogFormat,
	}
	api.Use(MyAccessProdStack...)
	api.SetApp(*router)
	belogs.Emergency(http.ListenAndServe(port, api.MakeHandler()))
}

// setup Https Server, listen on port. need crt and key files
func ListenAndServeTLS(port string, crtFile string, keyFile string, router *rest.App) {

	api := rest.NewApi()
	MyAccessProdStack := rest.AccessProdStack
	MyAccessProdStack[0] = &rest.AccessLogApacheMiddleware{
		Logger: belogs.GetLogger("access"),
		Format: rest.CombinedLogFormat,
	}
	api.Use(MyAccessProdStack...)
	api.SetApp(*router)
	//belogs.Emergency(http.ListenAndServe(port, api.MakeHandler()))
	belogs.Emergency(http.ListenAndServeTLS(port, crtFile, keyFile, api.MakeHandler()))
}

// return: map[fileFormName]=fileName, such as map["file1"]="aabbccdd.txt"
func ReceiveFiles(receiveDir string, r *http.Request) (receiveFiles map[string]string, err error) {
	belogs.Debug("ReceiveFiles(): receiveDir:", receiveDir)

	reader, err := r.MultipartReader()
	if err != nil {
		belogs.Error("ReceiveFiles(): err:", err)
		return nil, err
	}
	receiveFiles = make(map[string]string)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if !strings.HasSuffix(receiveDir, string(os.PathSeparator)) {
			receiveDir = receiveDir + string(os.PathSeparator)
		}
		file := receiveDir + part.FileName()
		form := strings.TrimSpace(part.FormName())
		belogs.Debug("ReceiveFiles():FileName:", part.FileName(), "   FormName:", part.FormName())
		if part.FileName() == "" { // this is FormData
			data, _ := ioutil.ReadAll(part)
			ioutil.WriteFile(file, data, 0644)
		} else { // This is FileData
			dst, _ := os.Create(file)
			defer dst.Close()
			io.Copy(dst, part)
		}
		receiveFiles[form] = file
	}
	return receiveFiles, nil
}
