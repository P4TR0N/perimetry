package eye_data

import (
	"fmt"
	"os"
)

func CreateFolder(userID string){

	err := os.MkdirAll("files/json/" +  userID, 0777)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Не удалось создать папку: JSON/" + userID)
	}

	err = os.MkdirAll("files/pdf/" +  userID, 0777)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Не удалось создать папку: PDF/" + userID)
	}

	return
}