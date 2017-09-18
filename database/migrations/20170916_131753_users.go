package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20170916_131753 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20170916_131753{}
	m.Created = "20170916_131753"

	migration.Register("Users_20170916_131753", m)
}

// Run the migrations
func (m *Users_20170916_131753) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(id SERIAL PRIMARY KEY, name character varying(25), profession character varying(25), dream character varying(25), hobby character varying(25), website_url TEXT NOT NULL, phone character varying(15), avatar_url TEXT NOT NULL, email character varying(35), password TEXT NOT NULL, token TEXT NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Users_20170916_131753) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
