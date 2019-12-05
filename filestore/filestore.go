package filestore

import (
	"gowiki2/domain"
	"io/ioutil"
)

type FileStore struct {
}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (f FileStore) Save(p *domain.Page) error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func (f FileStore) LoadPage(title string) (*domain.Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &domain.Page{Title: title, Body: body}, nil
}
