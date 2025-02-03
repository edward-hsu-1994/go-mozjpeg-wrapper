package go_mozjpeg_wapper

import (
	"errors"
	"fmt"
)

type MozJpegEncoderBuilder struct {
	cliExecutablePath string
	cliArgs           map[string]interface{}
}

const WellKnownMozJpegEncoderPath = "cjpeg"

func NewMozJpegEncoderBuilder() *MozJpegEncoderBuilder {
	return &MozJpegEncoderBuilder{
		cliExecutablePath: WellKnownMozJpegEncoderPath,
		cliArgs:           make(map[string]interface{}),
	}
}

func (builder *MozJpegEncoderBuilder) Clone() *MozJpegEncoderBuilder {
	newInstance := NewMozJpegEncoderBuilder()
	newInstance.cliExecutablePath = builder.cliExecutablePath
	newInstance.cliArgs = make(map[string]interface{})

	for key, value := range builder.cliArgs {
		newInstance.cliArgs[key] = value
	}

	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetCliExecutablePath(path string) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliExecutablePath = path
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetQuality(quality int) (*MozJpegEncoderBuilder, error) {
	if quality < 0 || quality > 100 {
		return builder, errors.New("quality must be in range 0-100")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["-quality"] = fmt.Sprintf("%d", quality)
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) SetDCScanOpt(mode int) (*MozJpegEncoderBuilder, error) {
	if mode < 0 || mode > 2 {
		return builder, errors.New("dc-scan-opt must be in range 0-2")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["-dc-scan-opt"] = fmt.Sprintf("%d", mode)
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) SetTuneMode(mode TuneMode) (*MozJpegEncoderBuilder, error) {
	if !validTuneModes[mode] {
		return builder, errors.New("invalid tune mode")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["-tune-"+string(mode)] = ""
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) SetDCTMethod(method DCTMethod) (*MozJpegEncoderBuilder, error) {
	if !validDCTMethods[method] {
		return builder, errors.New("invalid DCT method")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["-dct"] = string(method)
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) SetQuantTable(table QuantTable) (*MozJpegEncoderBuilder, error) {
	if !validQuantTables[table] {
		return builder, errors.New("invalid quantization table, must be in range 0-8")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["quant-table"] = fmt.Sprintf("%d", table)
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) SetQuantBaseline(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	if enabled {
		newInstance.cliArgs["-quant-baseline"] = ""
	} else {
		delete(newInstance.cliArgs, "quant-baseline")
	}
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetBaseline(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliArgs["-baseline"] = ""
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetProgressive(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliArgs["-progressive"] = ""
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetGrayscale(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliArgs["-grayscale"] = ""
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetSmooth(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliArgs["-smooth"] = ""
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetOptimize(enabled bool) *MozJpegEncoderBuilder {
	newInstance := builder.Clone()
	newInstance.cliArgs["-optimize"] = ""
	return newInstance
}

func (builder *MozJpegEncoderBuilder) SetMaxMemory(maxMemory int) (*MozJpegEncoderBuilder, error) {
	if maxMemory < 0 {
		return builder, errors.New("max memory must be a positive integer")
	}
	newInstance := builder.Clone()
	newInstance.cliArgs["-maxmemory"] = fmt.Sprintf("%d", maxMemory)
	return newInstance, nil
}

func (builder *MozJpegEncoderBuilder) Build() *MozJpegEncoder {
	return &MozJpegEncoder{
		cliExecutablePath: builder.cliExecutablePath,
		cliArgs:           builder.cliArgs,
	}
}
