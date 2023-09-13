package convertor

// TODO: deprecated test, we still need it but in some other place that validates the file format.
// func TestConvertedFilePath(t *testing.T) {
// 	test_cases := []struct {
// 		name           string
// 		input          string
// 		expectedOutput string
// 		expectedErr    error
// 	}{
// 		{"simple conversion", "./files/test.mp3", "files/test.wav", nil},
// 		{"convert from another place", "another.mp3", "files/another.wav", nil},
// 		{"convert same input", "test.wav", "files/test.wav", nil},
// 		{"wrong file format", "test.blu", "", errors.New("entered file extension is not valid for conversion, use a .mp3 or .wav file")},
// 	}

// 	for _, test := range test_cases {
// 		t.Run(test.name, func(t *testing.T) {
// 			output, err := convertedFilePath(test.input)
// 			if err != nil && err.Error() != test.expectedErr.Error() {
// 				t.Error(err)
// 			}
// 			if output != test.expectedOutput {
// 				t.Error("unexpected output", output)
// 			}
// 		})
// 	}
// }
