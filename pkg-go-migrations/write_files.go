package pkggomigrations

import (
	"fmt"
	"os"
	"time"
)

func write_in(s *string) error {

	err := os.Mkdir("migrations/applied", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create '/migrations/applied' path: ", err)
		return err
	}

	// fmt.Println("Write this function to write in sql files: Applied datetime")

	// fi, _ := ReadMigration("0001", ".up.sql")

	dt := time.Now()
	formattedTime := dt.Format("2006-01-02 15:04:05")

	*s = fmt.Sprintf("%v\n--Applied at: %s", *s, formattedTime)

	fmt.Println(*s)

	filePath := fmt.Sprintf("migrations/applied/%s", "filename.sql")
	newFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file: ", err)
		return err
	}
	defer newFile.Close()

	_, err = newFile.WriteString(*s)
	if err != nil {
		fmt.Println("Failed to create file: ", err)
		return err
	}

	return nil

}
