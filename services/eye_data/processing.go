package eye_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type JsonParser struct {
	Eye          string          `json:"eye"`
	Matrix       [][]int         `json:"matrix"`
	Duration     string          `json:"duration"`
	EyePositions [][]float64     `json:"eye_positions"`
	Events       [][]interface{} `json:"events"`
	Birthday     string          `json:"birthday"`
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}


func ProcessingData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//-----------Парсинг-----------------//
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Не удалось прочитать тело запроса")
	}
	email := r.Header.Get("email")
	send := r.Header.Get("send")
	userID := r.Header.Get("id")
	token := r.Header.Get("token")
	fmt.Println(email,send,userID,token)
	//-----------Парсинг-----------------//
	datatime := time.Now().Format("02.01.2006_15.04.05") //Получить точное время и дату, далее используется для создания файлов //2006-01-02

	var inputjson JsonParser
	err = json.Unmarshal(body, &inputjson)
	if err != nil {
		log.Println(err)
	}

	CreateFolder(userID)
	fmt.Println(inputjson.Duration)
	matrix, _ := json.Marshal(inputjson.Matrix)
	duration, _ := json.Marshal(inputjson.Duration)
	matrixstr := strings.Replace(string(matrix), " ", ", ",0)
	matrixjson := `{"matrix":` + matrixstr + `, "duration":` + string(duration) + `}`
	CreateFile(datatime, userID, string(body), inputjson.Eye, matrixjson)

	json_folder := "/home/tonometry/go-server/perimetry/files/json/" +  userID
	pdf_path := "/home/tonometry/go-server/perimetry/files/pdf/" +  userID + "/" + "field_" + inputjson.Eye + "_" + datatime + ".pdf"
	fmt.Println(json_folder)
	download_pdf :=  "http://files.kuse.pro/pdf/" +  userID + "/" + "field_" + inputjson.Eye + "_" + datatime + ".pdf"

	TransferToPDFGenerator(json_folder, pdf_path, inputjson.Birthday, userID, inputjson.Eye)
	w.Write([]byte(download_pdf))
	if email != "none" {
		SendPDFtoEmail(datatime, email, pdf_path, userID)
	}


}


func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//-----------Парсинг-----------------//
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Не удалось прочитать тело запроса")
	}
	email := r.Header.Get("email")
	send := r.Header.Get("send")
	userID := r.Header.Get("id")
	token := r.Header.Get("token") // Скоро реализую проверку токена
	fmt.Println(email,send,userID,token)
	//-----------Парсинг-----------------//
	datatime := time.Now().Format("02.01.2006_15.04.05") //Получить точное время и дату, далее используется для создания файлов //2006-01-02
	var inputjson JsonParser
	err = json.Unmarshal(body, &inputjson)
	if err != nil {
		log.Println(err)
	}
	CreateFolder(userID)

	json_folder := "/home/tonometry/go-server/perimetry/files/json/" +  userID
	pdf_path := "/home/tonometry/go-server/perimetry/files/pdf/" +  userID + "/" + "field_" + inputjson.Eye + "_" + datatime + ".pdf"
	fmt.Println(json_folder)
	download_pdf :=  "http://files.kuse.pro/pdf/" +  userID + "/" + "field_" + inputjson.Eye + "_" + datatime + ".pdf"
	TransferToPDFGenerator(json_folder, pdf_path, inputjson.Birthday, userID, inputjson.Eye)
	w.Write([]byte(download_pdf))
	if email != "none" {
		SendPDFtoEmail(datatime, email, pdf_path, userID)
	}


}
