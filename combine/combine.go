package combine

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"cyaml/types"

	log "github.com/sirupsen/logrus"
)

func Exec(path, prefix string) (*types.CyamlMeta, error) {
	files, err := traverse(path)
	if err != nil {
		return nil, err
	}
	re := &types.CyamlMeta{Type: "cyaml", Version: "1"}
	for _, file := range files {
		content, err := read(file)
		if err != nil {
			log.Errorf("Can't read file:%s error:%s", file, err)
			continue
		}
		re.Entries = append(re.Entries, types.CyamlEntry{Name: stripPrefix(file, prefix), Content: string(content)})
	}
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

func read(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
