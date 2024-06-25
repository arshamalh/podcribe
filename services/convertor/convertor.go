package convertor

import (
	"io"
	"podcribe/entities"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type I interface {
	// Convert a file format to another and returns the new path or raise an error
	Convert(podcast *entities.Podcast) error
}

type convertor struct {
}

func New() *convertor {
	return &convertor{}
}

// Convert turns mp3 file to wav file
// Default command is:
// ffmpeg -i input.mp3 -ar 16000 -ac 1 -c:a pcm_s16le output.wav
func (c convertor) Convert(podcast *entities.Podcast) error {
	// TODO: "" wav path should not be hard-coded
	return ffmpeg.Input(podcast.Mp3Path).Output(podcast.GetWavPath(""),
		ffmpeg.KwArgs{
			"c:a": "pcm_s16le",
			"ar":  16000,
			"ac":  1,
		}).Run()
}

// TODO: How not to store the input file on Disk and use io.Writer instead.
func (c convertor) ConvertOGGToMP3(oggPath string, mp3Path string) error {
	return ffmpeg.Input(oggPath).Output(mp3Path).Run()
}

func (c convertor) ConvertOGGToMP3Stream(input io.Reader, output io.Writer) error {
	return ffmpeg.
		Input("pipe:", ffmpeg.KwArgs{"format": "ogg"}).
		WithInput(input).
		Output("pipe:", ffmpeg.KwArgs{"format": "mp3"}).
		WithOutput(output).
		Run()
}
