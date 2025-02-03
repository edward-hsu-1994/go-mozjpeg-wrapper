package go_mozjpeg_wapper

import "testing"

func TestMozJpegEncoder_Encode(t *testing.T) {
	encoderBuilder := NewMozJpegEncoderBuilder()
	encoderBuilder = encoderBuilder.SetCliExecutablePath("echo")
	encoderBuilder, err := encoderBuilder.SetQuality(90)

	encoder := encoderBuilder.Build()

	err = encoder.Encode("/path/to/input", "/path/to/output")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
