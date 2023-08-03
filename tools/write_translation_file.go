package tools

import "os"

// Get a translation text and write it to a text file and returns the path
func WriteTranslation(filename, translation string) (filepath string, err error) {
	filepath = filename
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	if _, err = file.Write([]byte(translation)); err != nil {
		return "", err
	}

	return filepath, nil
}
