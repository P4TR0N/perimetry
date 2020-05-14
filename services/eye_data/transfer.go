package eye_data

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func TransferToPDFGenerator(json_folder, pdf_path, birthday, id, eye string){
	fmt.Println("Функция отправки данных в PDF")
	fmt.Println(json_folder, pdf_path, birthday, id, eye)
	payload := url.Values{}
	payload.Add("json_folder", json_folder)
	payload.Add("pdf_path", pdf_path)
	payload.Add("birthday", birthday)
	payload.Add("id", id)
	payload.Add("eye", eye)

	res, err := http.Get("http://localhost:5041/generate_pdf?" + payload.Encode())
	if err != nil {
		log.Panic(err)
	}
	status, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", status)
}