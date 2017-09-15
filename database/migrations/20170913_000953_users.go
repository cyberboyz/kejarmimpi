package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20170913_000953 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20170913_000953{}
	m.Created = "20170913_000953"

	migration.Register("Users_20170913_000953", m)
}

// Run the migrations
func (m *Users_20170913_000953) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(id SERIAL, name TEXT NOT NULL, profession TEXT NOT NULL, hobby TEXT NOT NULL, website_url TEXT NOT NULL, phone TEXT DEFAULT NULL, avatar_url TEXT NOT NULL, email TEXT NOT NULL, password TEXT NOT NULL, dream TEXT NOT NULL,  token TEXT NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Users_20170913_000953) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
