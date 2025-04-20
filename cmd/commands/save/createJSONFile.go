package save

import (
	"fmt"
	"os"
	"path/filepath"
)

func createJSONFile(fileName string) (*os.File, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("%s.json", fileName)

	filePath := filepath.Join(downloadFolderPath, name)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create file %q: %s\n", name, err)
	}

	return file, nil
}
