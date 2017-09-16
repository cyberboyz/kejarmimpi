package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Articles_20170916_133541 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Articles_20170916_133541{}
	m.Created = "20170916_133541"

	migration.Register("Articles_20170916_133541", m)
}

// Run the migrations
func (m *Articles_20170916_133541) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE articles(id SERIAL, title  TEXT NOT NULL, content  TEXT NOT NULL, id_user  SERIAL, id_category  SERIAL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Articles_20170916_133541) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE articles")
}
