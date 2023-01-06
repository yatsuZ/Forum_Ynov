package comment

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/posts"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/users"
)

// 0.5. refaire la bdd      ✅
// 1. fair une page pour les commentaire/ pour les poste✅
// 2. faire la redirection✅
// 2.25.netoyer le terminal parceque sinon jpp avancer ✅
// 2.5. cree des fonction qui cree des commentaire et l'implement dans la BDD voir modifier la structure du post✅
// 3. afficher les poste avec les commentaire✅
// 4.Faire une boite pour crée de nouveau commentaire

//ajouter 2 commentaire automatiquement au premier poste

//creation d'un type Comment
func createComment(db *sql.DB, content string, post_id int, autor string) posts.Comments {
	var C posts.Comments
	C.Content = content
	C.Post_id = post_id //fonction qui prent un id int et return un string
	C.Auteur = autor
	C.Number_Dislike = 0
	C.Number_Like = 0
	return C
}

//creation d'un comment et son ajout dans le tableau coment
func AddComment(db *sql.DB, content string, post_id int, autor string) (int64, error) {
	Newc := createComment(data.BDD, content, post_id, autor)
	fmt.Println("la ->",Newc)
	return insertIntoComment(db, Newc)
}

//incrementation d'un post dans le tableau posts
func insertIntoComment(db *sql.DB, c posts.Comments) (int64, error) {
	// c.Post_id =  users.SelectName(data.BDD, c.Auteur).Id
	fmt.Println("ici -> ",c)
	result, err := db.Exec("INSERT INTO 'comments' (content,Autor , post_id, like, dislike) VALUES (?,?,?,?,?)", c.Content, c.Auteur, c.Post_id, c.Number_Like, c.Number_Dislike)
	if err != nil {
		fmt.Println("l'erreur est ici")
		log.Fatal(err)
	}
	// fmt.Println("++ Comentaire Crée ++")
	return result.LastInsertId()
}

// var CommentTab = `
// CREATE TABLE IF NOT EXISTS comments (
// 	id	INTEGER NOT NULL,
// 	content	TEXT NOT NULL,
// 	Autor 	TEXT NOT NULL,
// 	post_id	INTEGER NOT NULL,
// 	like	INTEGER NOT NULL DEFAULT 0 ,
// 	dislike	INTEGER NOT NULL DEFAULT 0 ,
// 	FOREIGN KEY(post_id) REFERENCES posts(id),
// 	PRIMARY KEY(id AUTOINCREMENT)
// );

func ArrayComment() []posts.Comments {
	var L []posts.Comments
	rows := data.SelectAllFromTable(data.BDD, "comments")
	for rows.Next() {
		var c posts.Comments
		err := rows.Scan(&c.Id, &c.Auteur, &c.Content, &c.Post_id, &c.Number_Like, &c.Number_Dislike)

		if err != nil {
			log.Fatal(err)
		}
		L = append(L, c)
	}
	return L
}
func RowsAllIdPost(db *sql.DB, idPost int) *sql.Rows {
	// fmt.Println("Itoa result: ", strconv.Itoa(idPost))
	query := "SELECT * FROM comments WHERE post_id LIKE '" + strconv.Itoa(idPost) + "'"
	result, err := db.Query(query)
	if err != nil {
		// fmt.Println("l'erreur est ICI !!!!!!")
		log.Fatal(err)
	}
	return result
}
func ArrayCommentToPost(p posts.Post) []posts.Comments {
	var L []posts.Comments
	rows := RowsAllIdPost(data.BDD,p.Id)
	for rows.Next() {
		var c posts.Comments
		err := rows.Scan(&c.Id, &c.Auteur, &c.Content, &c.Post_id, &c.Number_Like, &c.Number_Dislike)

		if err != nil {
			log.Fatal(err)
		}
		L = append(L, c)
	}
	return L
}

func CreateNewComment(db *sql.DB, content string, p posts.Post) {
	_, err := AddComment(db, content, p.Id, users.SelectId(data.BDD, p.User_id).Name)
	if err == nil {
		p.AllComment = ArrayCommentToPost(p)
	}
}

func creeCommentTest() int {
	fmt.Println("\n+++ Creation des Comment Test +++")
	if len(ArrayComment()) > 1 {
		fmt.Println("+++ Comment Test Deja Cree +++")
		return 1
	}
	CreateNewComment(data.BDD, "Premier Commentaire au Premier Post", posts.SelectId(data.BDD, 1))
	CreateNewComment(data.BDD, "Deuxieme Commentaire au Premier Postt", posts.SelectId(data.BDD, 1))
	CreateNewComment(data.BDD, "Premier Commentaire au Deuxieme Postt", posts.SelectId(data.BDD, 2))

	fmt.Println("+++ Comment Test Cree +++")
	return 0
}

var T = creeCommentTest()
