package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20170911_232159 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20170911_232159{}
	m.Created = "20170911_232159"

	migration.Register("User_20170911_232159", m)
}

// Run the migrations
func (m *User_20170911_232159) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20170911_232159) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
