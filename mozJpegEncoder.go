package go_mozjpeg_wrapper

import (
	"os"
	"os/exec"
)

type MozJpegEncoder struct {
	cliExecutablePath string
	cliArgs           map[string]interface{}
}

func (encoder *MozJpegEncoder) Encode(inputFilePath string, outputFilePath string) error {
	return runCommand(encoder.cliExecutablePath, encoder.cliArgs, inputFilePath, outputFilePath)
}

func (encoder *MozJpegEncoder) EncodeAsync(inputFilePath string, outputFilePath string) chan error {
	var errChan = make(chan error)

	go func() {
		errChan <- encoder.Encode(inputFilePath, outputFilePath)

		defer close(errChan)
	}()

	return errChan
}

func runCommand(executablePath string, args map[string]interface{}, inputFilePath string, outputFilePath string) error {
	var cliArgs []string

	for key, value := range args {
		cliArgs = append(cliArgs, key)
		cliArgs = append(cliArgs, value.(string))
	}

	cliArgs = append(cliArgs, inputFilePath)

	cmd := exec.Command(executablePath, cliArgs...)

	outputBinary, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	err = writeBinaryToFile(outputFilePath, outputBinary)

	return nil
}

func writeBinaryToFile(filePath string, binary []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(binary)
	if err != nil {
		return err
	}

	return nil
}
