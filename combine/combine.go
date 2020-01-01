package combine

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Cyaml map[string]interface{}

func Exec(path, prefix string) (Cyaml, error) {
	files, err := traverse(path)
	if err != nil {
		return nil, err
	}
	re := make(Cyaml)
	re["type"] = "cyaml"
	re["version"] = "1"
	entries := []Cyaml{}
	for _, file := range files {
		content, err := read(file)
		if err != nil {
			log.Errorf("Can't read file:%s error:%s", file, err)
			continue
		}
		entry := make(Cyaml)
		entry["name"] = stripPrefix(file, prefix)
		entry["content"] = content
		entries = append(entries, entry)
	}
	re["files"] = entries
	return re, nil
}

func stripPrefix(file, prefix string) string {
	return strings.TrimPrefix(file, prefix)
}

func traverse(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}

func read(file string) (Cyaml, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	m := make(Cyaml)
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
