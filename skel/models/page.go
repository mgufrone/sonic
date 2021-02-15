package models

import (
	"errors"
	"log"

	"github.com/openware/pkg/database"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

func init() {
	Register("pages", &Page{}, func(rawYaml []byte) (interface{}, error) {
		list := []Page{}
		err := yaml.Unmarshal(rawYaml, &list)
		return list, err
	})
}

// Page : Table name is `Pages`
type Page struct {
	ID          uint   `gorm:"primarykey"`
	Path        string `gorm:"uniqueIndex;size:64;not null" yaml:"path"`
	Lang        string `yaml:"lang"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Body        string `yaml:"body"`
	database.Timestamps
}

// FIXME: page methods will not look nice. Rails has modules, and in Go
// it's better to create some service abstraction or transform to a regular function.

// FindPageByPath find and return a page by path
func FindPageByPath(path string) *Page {
	page := Page{}
	tx := app.GetDB().Where("path = ?", path).First(&page)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Fatalf("FindPageByPath failed: %s", tx.Error.Error())
		return nil
	}
	return &page
}

// ListPages returns all pages
func ListPages() []Page {
	pages := []Page{}
	tx := app.GetDB().Find(&pages)

	if tx.Error != nil {
		log.Fatalf("FindPageByPath failed: %s", tx.Error.Error())
	}
	return pages
}
