package pkggomigrations

import (
	"fmt"
	"os"
	"time"
)

func Write_in() {

	err := os.Mkdir("migrations/applied", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create '/migrations/applied' path: ", err)
	}

	fmt.Println("Write this function to write in sql files: Applied datetime")

	fi, _ := ReadMigration("0001", ".up.sql")

	dt := time.Now()
	formattedTime := dt.Format("2006-01-02 15:04:05")

	fi = fmt.Sprintf("%v\n--Applied at: %s", fi, formattedTime)

	fmt.Println(fi)

	filePath := fmt.Sprintf("migrations/applied/%s", "filename.sql")
	newFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file: ", err)
	}
	defer newFile.Close()

	_, err = newFile.WriteString(fi)
	if err != nil {
		fmt.Println("Failed to create file: ", err)
	}

}
