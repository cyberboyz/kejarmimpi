package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comments_20170916_133552 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comments_20170916_133552{}
	m.Created = "20170916_133552"

	migration.Register("Comments_20170916_133552", m)
}

// Run the migrations
func (m *Comments_20170916_133552) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comments(id SERIAL, comment  TEXT NOT NULL, id_user SERIAL, id_article SERIAL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Comments_20170916_133552) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE comments")
}