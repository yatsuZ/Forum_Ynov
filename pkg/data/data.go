package data // here will be database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var BDD = InitDatabase()

var CategorieTab = `
CREATE TABLE IF NOT EXISTS categories (
	id	INTEGER,
	name TEXT NOT NULL UNIQUE,
	PRIMARY KEY(id AUTOINCREMENT)
);
`

var UserTab = `
CREATE TABLE IF NOT EXISTS users (
	id	INTEGER NOT NULL,
	name	TEXT NOT NULL UNIQUE,
	email	TEXT NOT NULL UNIQUE,
	pasword	BLOB NOT NULL, 
	PRIMARY KEY(id AUTOINCREMENT)
);
`

//on ressaye comme ça bien on redemare // Sa c'est moi qui a ecris sa hier ok

var PostTab = `
CREATE TABLE IF NOT EXISTS posts (
	id	INTEGER NOT NULL,
	date_post	TEXT,
	content	TEXT NOT NULL,
	user_id	INTEGER NOT NULL,
	like	INTEGER NOT NULL DEFAULT 0, 
	dislike	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY(user_id) REFERENCES users(id),
	PRIMARY KEY(id AUTOINCREMENT)
);
`
var CommentTab = `
CREATE TABLE IF NOT EXISTS comments (
	id	INTEGER NOT NULL,
	content	TEXT NOT NULL,
	Autor 	TEXT NOT NULL,
	post_id	INTEGER NOT NULL,
	like	INTEGER NOT NULL DEFAULT 0 ,
	dislike	INTEGER NOT NULL DEFAULT 0 ,
	FOREIGN KEY(post_id) REFERENCES posts(id),
	PRIMARY KEY(id AUTOINCREMENT)
);
`
var LikeTab = `
CREATE TABLE IF NOT EXISTS like (
	id	INTEGER NOT NULL,
	content	TEXT NOT NULL,
	user_id	INTEGER NOT NULL,
	post_id	INTEGER,
	comment_id	INTEGER,
	type	INTEGER,
	FOREIGN KEY(comment_id) REFERENCES comments(id),
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(post_id) REFERENCES posts(id),
	PRIMARY KEY(id AUTOINCREMENT)
);
`

var Bridge_Post_CategorieTab = `
CREATE TABLE IF NOT EXISTS post_categorie (
	id	INTEGER NOT NULL,
	post_id	integer NOT NULL,
	cat_id	integer DEFAULT 6,
	PRIMARY KEY(id AUTOINCREMENT)
);
`

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "Forum.db")

	if err != nil {
		log.Fatal(err)
	}

	sqltStmt := `
    PRAGMA foreign_keys = ON;` + UserTab + CommentTab + PostTab + LikeTab + CategorieTab + Bridge_Post_CategorieTab
	_, err = db.Exec(sqltStmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creation de la Base de Donner Forum")
	return db
}

//selectione la table de la bdd et return un tableau
func SelectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// var CategorieTab = `
// 	CREATE TABLE IF NOT EXISTS catgories (
// 	name TEXT NOT NULL PRIMARY KEY
// 	);`

// var UserTab = `
// 	CREATE TABLE IF NOT EXISTS users (
// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	name TEXT NOT NULL UNIQUE,
// 	email TEXT NOT NULL UNIQUE,
// 	pasword TEXT NOT NULL
// 	);
// `

// var PostTab = `
// 	CREATE TABLE IF NOT EXISTS posts (
// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	catgegorie TEXT,
// 	title TEXT NOT NULL,
// 	date_post  DATETIME ,
// 	content TEXT NOT NULL,
// 	user_id INTEGER NOT NULL,
// 	like INTEGER,
// 	dislike INTEGER,
// 	FOREIGN KEY (catgegorie) REFERENCES catgories(name),
// 	FOREIGN KEY (user_id) REFERENCES users(id)
// 	);
// `
// var CommentTab = `
// 	CREATE TABLE IF NOT EXISTS comments (
// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	content TEXT NOT NULL,
// 	user_id INTEGER NOT NULL,
// 	post_id INTEGER NOT NULL,
// 	like INTEGER,
// 	dislike INTEGER,
// 	FOREIGN KEY (user_id) REFERENCES users(id),
// 	FOREIGN KEY (post_id) REFERENCES posts(id)
// 	);
// `
// var LikeTab = `
// 	CREATE TABLE IF NOT EXISTS like (
// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	content TEXT NOT NULL,
// 	user_id INTEGER NOT NULL,
// 	post_id INTEGER ,
// 	comment_id INTEGER,
// 	type INTEGER,
// 	FOREIGN KEY (user_id) REFERENCES users(id),
// 	FOREIGN KEY (post_id) REFERENCES posts(id),
// 	FOREIGN KEY (comment_id) REFERENCES comments(id)
// 	);
// `

// func InitDataBase() *sql.DB {
// 	db, err := sql.Open("sqlite3", "Forum.db")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sqlStmt := `PRAGMA foreign_keys = ON;` + CategorieTab + UserTab + PostTab + CommentTab + LikeTab
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Creation de la Base de Donner Forum")
// 	return db
// }
