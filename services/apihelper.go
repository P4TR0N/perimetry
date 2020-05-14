package services

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"perimetry/v1/common"
)

var (
	perimetryAppKey  = flag.String("PerimetryAppKey", common.PerimetryAppKey, "PerimetryAppKey URL and port")
)

// SystemStatus возвращает 200, в случаа, если http сервер запущен и работает корректно. 
func SystemStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// AppKey приложения Perimetry для авторизации в FushionAuth
func AppKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	appKey := `{"applicationId": " `+ *perimetryAppKey + `"}`
	w.Write([]byte(appKey))
}

func PerimetryAppVersion(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	file, err := os.Open("version.json")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for{
		n, err := file.Read(data)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		w.Write([]byte(string(data[:n])))
	}
}