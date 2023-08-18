package main

import (
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"root/constants"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func main() {
	fileInfos, err := sqlFiles.ReadDir(constants.FILE_SQL_NAME)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		filePath := filepath.Join(constants.FILE_SQL_NAME, fileInfo.Name()) //sql + / + fileInfo.Name()
		sqlContent, err := sqlFiles.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File: %s\nContent:\n%s\n", filePath, sqlContent)
	}
}
