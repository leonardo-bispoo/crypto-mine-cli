package save

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func createCSVFile(fileName string) (*os.File, *csv.Writer, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, nil, err
	}

	name := fmt.Sprintf("%s.csv", fileName)

	filePath := filepath.Join(downloadFolderPath, name)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create file %q: %v\n", name, err)
	}

	writer := csv.NewWriter(file)

	return file, writer, nil
}
