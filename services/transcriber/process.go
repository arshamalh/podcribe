package transcriber

import (
	"fmt"
	"io"
	"os"
	"time"

	// Package imports
	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	wav "github.com/go-audio/wav"
)

func Process(model whisper.Model, podcast_path string, informer io.Writer, output io.Writer) error {
	var data []float32

	// Create processing context
	context, err := model.NewContext()
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", context.SystemInfo())

	// Open the file
	fmt.Fprintf(informer, "Loading %q\n", podcast_path)
	fh, err := os.Open(podcast_path)
	if err != nil {
		return err
	}
	defer fh.Close()

	// Decode the WAV file - load the full buffer
	dec := wav.NewDecoder(fh)
	if buf, err := dec.FullPCMBuffer(); err != nil {
		return err
	} else if dec.SampleRate != whisper.SampleRate {
		return fmt.Errorf("unsupported sample rate: %d", dec.SampleRate)
	} else if dec.NumChans != 1 {
		return fmt.Errorf("unsupported number of channels: %d", dec.NumChans)
	} else {
		data = buf.AsFloat32Buffer().Data
	}

	// Segment callback when -tokens is specified
	var cb whisper.SegmentCallback

	// Process the data
	fmt.Fprintf(informer, "processing %q\n", podcast_path)
	context.ResetTimings()
	if err := context.Process(data, cb, nil); err != nil {
		return err
	}

	context.PrintTimings()

	return OutputSRT(output, context)
}

// Output text as SRT file
func OutputSRT(w io.Writer, context whisper.Context) error {
	n := 1
	for {
		segment, err := context.NextSegment()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Fprintln(w, n)
		fmt.Fprintln(w, srtTimestamp(segment.Start), " --> ", srtTimestamp(segment.End))
		fmt.Fprintln(w, segment.Text)
		fmt.Fprintln(w, "")
		n++
	}
}

// Return srtTimestamp
func srtTimestamp(t time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d", t/time.Hour, (t%time.Hour)/time.Minute, (t%time.Minute)/time.Second, (t%time.Second)/time.Millisecond)
}
