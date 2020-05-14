package eye_data

import (
	"fmt"
	"os"
)

func CreateFile(datatime, userID, body, eye, matrixstr string) {
	//Создаем исследования по глазам
	field_filename := "files/json/" +  userID + "/" + "field_" + eye + "_" + datatime + ".json"
	fieldfile, err := os.Create(field_filename)
	if err != nil{
		fmt.Println("Не создать json файл: " + field_filename, err)
		os.Exit(1)
	}
	defer fieldfile.Close()
	fieldfile.WriteString(matrixstr)

	//Создаем общий JSON
	procedure_filename := "files/json/" +  userID + "/" + "procedure_" + datatime + ".json"
	procedurefile, err := os.Create(procedure_filename)
	if err != nil{
		fmt.Println("Не создать json файл: " + procedure_filename, err)
		os.Exit(1)
	}
	defer procedurefile.Close()
	procedurefile.WriteString(body)

}