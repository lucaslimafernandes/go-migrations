package pkggomigrations

import (
	"fmt"
	"os"
	"time"
)

// write_in logs the migration information to some storage (could be a file or a database)
func write_in(s *string, flname string, rowsAff int, lastId int, path string) error {

	mPath := fmt.Sprintf("%v/applied", path)
	_ = os.Mkdir(mPath, os.ModePerm)

	dt := time.Now()
	formattedTime := dt.Format("2006-01-02 15:04:05")

	*s = fmt.Sprintf("%v\n--Applied at: %s", *s, formattedTime)
	*s = fmt.Sprintf("%v\n--Rows affected: %d", *s, rowsAff)
	*s = fmt.Sprintf("%v\n--Last Id inserted: %d", *s, lastId)

	fmt.Println(*s)

	filePath := fmt.Sprintf("%s/%s", mPath, flname)
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
