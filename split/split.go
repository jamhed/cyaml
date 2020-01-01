package split

import (
	"cyaml/types"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

func Exec(file, prefix string) error {
	meta, err := read(file)
	if err != nil {
		return err
	}
	for _, entry := range meta.Entries {
		targetFile := path.Join(prefix, entry.Name)
		targetPath := path.Dir(targetFile)
		err := os.MkdirAll(targetPath, os.ModePerm)
		if err != nil {
			log.Errorf("Can't create target path %s, error:%s", targetPath, err)
			continue
		}
		buf, err := yaml.Marshal(entry.Content)
		if err != nil {
			log.Errorf("Can't unmarshal entry %s, error:%s", entry.Name, err)
			continue
		}
		ioutil.WriteFile(targetFile, buf, os.ModePerm)
	}
	return nil
}

func read(file string) (*types.CyamlMeta, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	m := &types.CyamlMeta{}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
