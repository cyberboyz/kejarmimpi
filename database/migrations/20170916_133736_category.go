package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Category_20170916_133736 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Category_20170916_133736{}
	m.Created = "20170916_133736"

	migration.Register("Category_20170916_133736", m)
}

// Run the migrations
func (m *Category_20170916_133736) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE category(id SERIAL PRIMARY KEY, category  TEXT NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Category_20170916_133736) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE comments")
}
