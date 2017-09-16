package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Articles_20170916_153235 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Articles_20170916_153235{}
	m.Created = "20170916_153235"

	migration.Register("Articles_20170916_153235", m)
}

// Run the migrations
func (m *Articles_20170916_153235) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE articles(id SERIAL PRIMARY KEY, title  TEXT NOT NULL, content  TEXT NOT NULL, id_user  SERIAL REFERENCES Users(id), id_category SERIAL REFERENCES category(id), created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Articles_20170916_153235) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE articles")
}
