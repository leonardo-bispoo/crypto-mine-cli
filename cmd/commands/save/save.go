package save

import (
	"errors"
	"log"
)

func Save(fileType, fileName string) error {
	if fileType != "json" && fileType != "csv" {
		return errors.New("file type not supported, try json or csv")
	}

	if fileType == "json" {
		if err := saveInJSON(fileName); err != nil {
			return err
		}
	}

	if fileType == "csv" {
		if err := saveInCSV(fileName); err != nil {
			return err
		}
	}

	log.Printf("File %s persisted in the Downloads folder in %s format", fileName, fileType)

	return nil
}
