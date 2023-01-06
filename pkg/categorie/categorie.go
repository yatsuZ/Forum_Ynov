package categorie //here will be creating posts functions

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	_ "time"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
	_ "git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
)

type Categories struct {
	Id   int
	Name string
}

func createCategorie(name string) Categories {
	var C Categories
	C.Name = name
	return C
}

func insertIntoCategorie(db *sql.DB, C Categories) (int64, error) {
	result, err := db.Exec("INSERT INTO 'categories' (name) VALUES (?)", C.Name)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Creation de Categorie " + C.Name)
	return result.LastInsertId()
}

func AddCategorie(db *sql.DB, name string) (int64, error) {
	var vide Categories
	NewC := createCategorie(name)
	if SelectedNameCategorie(db, name) == vide {
		return insertIntoCategorie(db, NewC)
	} else {
		// fmt.Println("Categorie " + name + " deja cree ")
		return 1, nil
	}
}

func SelectedNameCategorie(db *sql.DB, name string) Categories {
	var selection Categories
	db.QueryRow("SELECT * FROM categories WHERE name = ?", name).Scan(&selection.Id, &selection.Name)
	return selection
}

func SelectedIdCategorie(db *sql.DB, id int) Categories {
	var selection Categories
	db.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(&selection.Id, &selection.Name)
	return selection
}

//afficher le tableau USER
func DisplayCatgorieRows() int {
	rows := data.SelectAllFromTable(data.BDD, "categories")
	for rows.Next() {
		var C Categories
		err := rows.Scan(&C.Id, &C.Name)
		if err != nil {
			log.Fatal(err)
			return 1
		}
		// fmt.Println(C)
	}
	return 0
}

//slection tout les name ayant c'ette paterne dans user return un tableau d'user
func RowsAllIdPost(db *sql.DB, idPost int) *sql.Rows {
	// fmt.Println("Itoa result: ", strconv.Itoa(idPost))
	query := "SELECT * FROM post_categorie WHERE post_id LIKE '" + strconv.Itoa(idPost) + "'"
	result, err := db.Query(query)
	if err != nil {
		// fmt.Println("l'erreur est ICI !!!!!!")
		log.Fatal(err)
	}
	return result
}

func GetCategorieByIdPost(db *sql.DB, idPost int) []string {
	// fmt.Println("\n////J'essaye de trouver toute les categorie a partir d'un ID post //////")
	var AllCatgorie []string
	rows := RowsAllIdPost(db, idPost)

	for rows.Next() {
		var C Categories
		var Id_Post int
		var cle_Post_Cat int
		var idCatgorie int
		err := rows.Scan(&cle_Post_Cat, &Id_Post, &idCatgorie)
		if err != nil {
			log.Fatal(err)
		}
		C = SelectedIdCategorie(db, idCatgorie)
		// fmt.Println(C.Name)
		AllCatgorie = append(AllCatgorie, C.Name)
	}
	// fmt.Println("\n////FIN //////")
	return AllCatgorie
}

//netoyer les fmt.Println
// var _ = DisplayCatgorieRows()
var _, _ = fmt.Println("\n--- Initialisation des Categorie ---")
var _, _ = AddCategorie(data.BDD, "tech")
var _, _ = AddCategorie(data.BDD, "film")
var _, _ = AddCategorie(data.BDD, "sport")
var _, _ = AddCategorie(data.BDD, "anime")
var _, _ = AddCategorie(data.BDD, "music")
var _, _ = AddCategorie(data.BDD, "other")
var _, _ = fmt.Println("--- Fin de l'Initialisation ---\n")

// <option name="1" value="1">tech</option>
// <option name="2" value="2">film</option>
// <option name="3" value="3">sport</option>
// <option name="4" value="4">anime</option>
// <option name="5"value="5">music</option><!-- form ? -->

// var CategorieTab = `
// CREATE TABLE IF NOT EXISTS categories (
// 	name	TEXT NOT NULL UNIQUE,
// 	id	INTEGER,
// 	PRIMARY KEY(id AUTOINCREMENT)
// );
//  `
