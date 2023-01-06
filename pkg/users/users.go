package users // here will be users functions (create users ...etc)

import (
	"database/sql"
	"fmt"
	"log"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/utils"
)

type User struct {
	Id      int    //clé primaire
	Name    string //Pseudo de l'utilisateur
	Email   string //email de l'utilisateur (unique)
	Pasword []byte //Mot de passe de lutilisateur
	Alert   string
}

func creatUser(name string, email string, password string) User {
	var u User
	u.Name = name
	u.Email = email
	u.Pasword, _ = utils.GetPasswordHash([]byte(password))
	return u
}

func CheckSignIn(userName string, password string) bool { // Une fonction que premet si le user et pass exists dans la base de données pour signIN

	return true
}

func CheckUnicName(db *sql.DB, userName string) {

}

// func SelectedNameCategorie(db *sql.DB, name string) Categories {
// 	var selection Categories
// 	db.QueryRow("SELECT * FROM categories WHERE name = ?", name).Scan(&selection.Id, &selection.Name)
// 	return selection
// }

//slection tout les name ayant c'ette paterne dans user return un tableau d'user
func SelectName(db *sql.DB, name string) User {
	var U User
	db.QueryRow("SELECT * FROM users WHERE name = ?", name).Scan(&U.Id, &U.Name, &U.Email, &U.Pasword)
	// fmt.Println(U)
	return U
}

func SelectId(db *sql.DB, id int) User {
	var U User
	db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&U.Id, &U.Name, &U.Email, &U.Pasword)
	// fmt.Println(U)
	return U
}

func SelectEmail(db *sql.DB, email string) User {
	var U User
	db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&U.Id, &U.Name, &U.Email, &U.Pasword)
	// fmt.Println(U)
	return U
}

//incrementation d'un user dans le tableau users
func insertIntoUsers(db *sql.DB, u User) (int64, error) {
	result, err := db.Exec("INSERT INTO 'users' (name,email,pasword) VALUES (?,?,?)", u.Name, u.Email, u.Pasword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nUn Nouvelle Utilisateur\n-------------------------------------------------------\n")
	return result.LastInsertId()
}

func AddUser(db *sql.DB, name string, email string, password string) (int64, error) {
	Newu := creatUser(name, email, password)
	if UniqueEmail(db, email) && UniqueName(db, name) {
		return insertIntoUsers(db, Newu)
	} else {
		fmt.Println("Le Pseudo ou l'adresse Mail a deja etait utiliser. ")
		return 1, nil
	}
}

func RowsNil(rows *sql.Rows) bool {
	var i = 0
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Pasword)
		if err != nil {
			log.Fatal(err)
		}
		i++
		// fmt.Println(u)
	}
	if i == 0 {
		// fmt.Println("VIDE\ntrue")
		return true
	}
	// fmt.Println("false")
	return false
}

//fonction qui verifie qu'il n'y pas pas 2 fois l'email dans la bdd
func UniqueEmail(db *sql.DB, email string) bool {
	var vide User
	// fmt.Println(vide)
	fmt.Println("---Verification si l'Email est unique---")
	// fmt.Println("l'email = ", email)
	// fmt.Println(SelectEmail(db, email) != vide)
	if SelectEmail(db, email).Email != vide.Email {
		// fmt.Println("\nIl y a deja un utilisateur avec l'email :\n --  \"" + email + "\".")
		fmt.Println("---Fin REDONDONCE!!---")
		return false
	}
	fmt.Println("---Fin l'Email est Unique ---")
	// fmt.Println("\nCette email n'a pas etait utiliser il est unique")
	return true
}

//fonction qui verifie qu'il n'y pas pas 2 fois le meme pseudo dans la bdd
func UniqueName(db *sql.DB, name string) bool {
	var vide User
	// fmt.Println(vide)
	fmt.Println("---Verification si le Pseudo est unique---")
	// fmt.Println("Pseudo = ", name)
	// fmt.Println(SelectName(db, name)!=vide)
	if SelectName(db, name).Name != vide.Name {
		// fmt.Println("\nIl y a deja un utilisateur avec le pseudo :\n - \"" + name + "\".")
		fmt.Println("---Fin REDONDONCE!!---")
		return false
	}
	// fmt.Println("\nCe pseudo n'a pas etait crée tu es unique")
	fmt.Println("---Fin le Pseudo est Unique---")
	return true
}

//afficher le tableau USER
func DisplayUserRows(rows *sql.Rows) {
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Pasword)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}
}

func SelectUserByName(db *sql.DB, name string) User {
	var selection User
	db.QueryRow("SELECT * FROM users WHERE name = ?", name).Scan(&selection.Id, &selection.Name, &selection.Email, &selection.Pasword)
	//<--ICI string ?? password sting ?? d'accord
	return selection //
}

var _, _ = AddUser(data.BDD, "admin", "admin@amin.com", "admin")
