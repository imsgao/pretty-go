package prettygo

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileEx(codingType string, v interface{}, filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err.Error())
	}

	if err := Decode(codingType, b, v); err != nil {
		return fmt.Errorf("error reading file:: %s", err.Error())
	}
	return nil
}

func ReadFile(v interface{}, filename string) error {
	codingType, err := CodingTypeFromExtension(filepath.Ext(filename))
	if err != nil {
		return err
	}
	return ReadFileEx(codingType, v, filename)
}

func WriteFileEx(codingType string, v interface{}, filename string) error {
	b, err := Encode(codingType, v)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, b, 0666); err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}
	return nil
}

func WriteFile(v interface{}, filename string) error {
	codingType, err := CodingTypeFromExtension(filepath.Ext(filename))
	if err != nil {
		return err
	}
	return WriteFileEx(codingType, v, filename)
}
