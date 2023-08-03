All the steps:

1 - Finding the link

2 - Downloading the podcast

3 - Converting MP3 to wav

4 - Transcribing wav

5 - Translating

Possible combinations:

345 => User have a MP3 file to transcribe

45 => User have a WAV file to transcribe

12 => User have a link to just download (no transcription, conversion or translation needed)

1234 => User want's downloading and transcribing but don't want translations

12345 => full flow

34 => User have a MP3 file to transcribe (and 5 translate)

4 => User have a WAV file to transcribe (and 5 translate)

## How to use
```bash
podcribe transcribe go-time-330.mp3 # => automatically transcribes and and translates
podcribe transcribe https://podcast.google.com/... # => automatically downloads and transcribes and and translates
podcribe start --telegram-enabled true --telegram-token hgkhkgkk --ui-enabled true
```
Automatically detect whether to download a file or directly transcribe a file in a provided path

```
LIBRARY_PATH=whisper C_INCLUDE_PATH=whisper go run main.go transcribe "files/01 Into You.wav"
```
<!-- TODO: podcribe model command -->

resources: read about go bindings
https://github.com/ggerganov/whisper.cpp/tree/master/bindings/go
