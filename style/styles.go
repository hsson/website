package style

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hsson/go-website/util"
)

const cssPattern = "*.css"

type stringChannelRes struct {
	result string
	err    error
}

// Generate will create CSS stylesheets in the specified subdirectory
// inside the specified artifacts directory based on the files in the specified
// input directory.
func Generate(inputDir, artifactDir, subDir string) ([]string, error) {
	fmt.Println("Generating stylesheets...")
	outDir := filepath.Join(artifactDir, subDir)
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created %s\n", outDir)
	sheets, err := copyRegulars(inputDir, artifactDir, subDir)

	return sheets, err
}

func copyRoutine(inputFile, outputDir string, resChannel chan stringChannelRes) {
	_, filename := filepath.Split(inputFile)
	outPath := filepath.Join(outputDir, filename)
	err := util.CopyFile(inputFile, outPath)

	resChannel <- stringChannelRes{filename, err}
}

func copyRegulars(inputDir, artifactDir, subDir string) (result []string, err error) {
	stylesheets, err := filepath.Glob(filepath.Join(inputDir, cssPattern))
	if err != nil {
		return result, err
	}

	resultChannel := make(chan stringChannelRes)
	for _, sheet := range stylesheets {
		go copyRoutine(sheet, filepath.Join(artifactDir, subDir), resultChannel)
	}
	for i := 0; i < len(stylesheets); i++ {
		res := <-resultChannel
		if res.err != nil {
			fmt.Printf("Failed to create \"%s\"\n", res.result)
			return result, res.err
		}
		fmt.Printf("Created %s\n", filepath.Join(artifactDir, subDir, res.result))
		result = append(result, filepath.Join(subDir, res.result))
	}
	return result, err
}
