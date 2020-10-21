package db

import (
	"io/ioutil"
	"path"
	"strings"
)

type SQLManager struct {
	Data map[string]string
}

func NewSQLManager(path string) *SQLManager {
	data := make(map[string]string)
	manager := SQLManager{
		Data: data,
	}
	manager.LoadSQL(path)
	return &manager
}

func (s *SQLManager) LoadSQL(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("Error load sql: " + err.Error())
	}
	for _, f := range files {
		if !strings.Contains(f.Name(), ".sql") {
			continue
		}
		b, err := ioutil.ReadFile(path.Join(dir, f.Name()))
		if err != nil {
			panic("Error load sql: " + err.Error())
		}
		content := string(b)
		name := strings.Replace(f.Name(), ".sql", "", 1)
		s.Data[name] = content
	}
}
