package models_test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func TestCreateDB(t *testing.T) {
	// 用来创建数据库
	_, err := os.OpenFile("../data/myblog.db", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err := sql.Open("sqlite3", "../data/myblog.db")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	table_blog_owner := `
	CREATE TABLE blog_owner (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		image_icon_path varchar(100) NOT NULL,
		introduction varchar(200) NOT NULL
	)`

	table_category := `
	CREATE TABLE category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(100) NOT NULL
	)
	`

	table_article := `
	CREATE TABLE article(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title varchar(200) NOT NULL,
		post_date datetime NOT NULL,
		content text,
		category_id INTEGER,
		CONSTRAINT FK_category_id FOREIGN key(category_id) REFERENCES category (id)
	)
	`

	table_comment := `
	CREATE TABLE comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Content text NOT NULL,
		article_id INTEGER,
		CONSTRAINT FK_article_id FOREIGN key(article_id) REFERENCES article(id)
	)
	`
	_, err = db.Exec(table_blog_owner)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(table_category)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(table_article)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(table_comment)
	if err != nil {
		fmt.Println(err.Error())
	}
}
