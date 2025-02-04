package go_mozjpeg_wrapper

import "testing"

func TestMozJpegEncoderBuilder_SetQuality(t *testing.T) {
	builder := NewMozJpegEncoderBuilder()
	newBuilder, err := builder.SetQuality(90)
	if err != nil || newBuilder.cliArgs["-quality"] != "90" {
		t.Errorf("Expected quality to be set to 90, got %v", newBuilder.cliArgs["-quality"])
	}
	_, err = builder.SetQuality(110)
	if err == nil {
		t.Error("Expected error for quality out of range, got nil")
	}
}

func TestMozJpegEncoderBuilder_SetDCScanOpt(t *testing.T) {
	builder := NewMozJpegEncoderBuilder()
	_, err := builder.SetDCScanOpt(3)
	if err == nil {
		t.Error("Expected error for dc-scan-opt out of range, got nil")
	}
}
