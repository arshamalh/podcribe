package convertor

import (
	"errors"
	"path"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type I interface {
	// Convert a file format to another and returns the new path or raise an error
	Convert(file2convert string) (converted_file_path string, err error)
}

type convertor struct {
}

func New() *convertor {
	return &convertor{}
}

// Convert turns mp3 file to wav file
// Default command is:
// ffmpeg -i input.mp3 -ar 16000 -ac 1 -c:a pcm_s16le output.wav
func (c convertor) Convert(file2convert string) (string, error) {
	converted_file_path, err := convertedFilePath(file2convert)
	if err != nil {
		return "", err
	}
	if err := ffmpeg.Input(file2convert).Output(converted_file_path,
		ffmpeg.KwArgs{
			"c:a": "pcm_s16le",
			"ar":  16000,
			"ac":  1,
		}).Run(); err != nil {
		return "", err
	}
	return converted_file_path, nil
}

func convertedFilePath(file2convert string) (string, error) {
	_, file := path.Split(file2convert)
	if strings.HasSuffix(file, ".wav") {
		return path.Join("files", file), nil
	}

	extLessName, found := strings.CutSuffix(file, ".mp3")
	if !found {
		return "", errors.New("entered file extension is not valid for conversion, use a .mp3 or .wav file")
	}
	return path.Join("files", extLessName+".wav"), nil
}
