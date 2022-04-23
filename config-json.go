package configjson

import (
	jsoniter "github.com/json-iterator/go"
	"bytes"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func readFile(path string) ([]byte, error) {
	configFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil { return []byte{}, err }
	defer configFile.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(configFile)
	return buf.Bytes(), nil
}


func unmarshalConfig(data []byte, config any) error {
	err := json.Unmarshal(data, config)
	return err
}

func ReadConfigFile(path string, config any) error {
	data, err := readFile(path)
	if err != nil { return err }
	err = unmarshalConfig(data, config)
	return err
}