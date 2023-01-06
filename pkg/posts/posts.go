package posts //here will be creating posts functions

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/categorie"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/users"
)

type Comments struct {
	Id             int
	Content        string
	Auteur         string
	Post_id        int
	Number_Like    int
	Number_Dislike int
}

type Post struct {
	Id             int    //clé primaire du post
	Created_At     string //Date de la publication //Moi mr METEHRI
	Content        string //Conteneue du post
	User_id        int    //Clé etranger qui est un id d'un user existant
	User_name      string
	CategoriePost  []string //tableau de categorie
	Number_Like    int      //Number    //Nombre de like
	Number_Dislike int      //Nombre de dislike
	AllComment     []Comments
}

// ########################## Post ################################
//creation d'un type post
func createPost(content string, User_Id int) Post {
	var p Post
	p.Content = content
	p.User_id = User_Id
	p.User_name = users.SelectId(data.BDD, p.User_id).Name
	p.Created_At = time.Now().Format("01-02-2006 15:04:05 Mon")
	return p
}

//incrementation d'un post dans le tableau posts
func insertIntoPosts(db *sql.DB, p Post) (int64, error) {
	result, err := db.Exec("INSERT INTO 'posts' (content, user_id, date_post) VALUES (?,?,?)", p.Content, p.User_id, p.Created_At)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("~~ Post Cree ~~")
	return result.LastInsertId()
}

func CreateNewPost(db *sql.DB, content string, id_user int, categoryId []int) {
	id, err := AddPost(db, content, id_user)
	if err == nil {
		AssignCategoryToPost(db, int(id), categoryId)
	}
}

// like pour post
func AssignlikeToPost(db *sql.DB, id int) {

	result, err := db.Exec("UPDATE posts SET like=like+1 WHERE id=?", id)
	if err != nil {
		log.Println(err)
		return
	}
	result.RowsAffected()
	fmt.Println("3HERETHE TEST----- LIKE")
	fmt.Println(result.RowsAffected())

}

//creation d'un post et son ajout dans le tableau posts
func AddPost(db *sql.DB, content string, id_user int) (int64, error) {
	Newp := createPost(content, id_user)
	return insertIntoPosts(db, Newp)
}

//Todo (Yassine) : A completer
//func AddComment(db *sql.DB, content string, id_user int) (int64, error) {
//	Newp := createPost(content, id_user)
//	return insertIntoPosts(db, Newp)
//}

func AssignCategoryToPost(db *sql.DB, post_id int, categoryId []int) {
	if categoryId == nil {
		result, err := db.Exec("INSERT INTO 'post_categorie' (post_id, cat_id) VALUES (?,?)", post_id, 6) //6 == other
		if err != nil {
			log.Fatal(err)
		}
		result.LastInsertId()
		return
	}
	for _, cat_id := range categoryId {
		result, err := db.Exec("INSERT INTO 'post_categorie' (post_id, cat_id) VALUES (?,?)", post_id, cat_id)
		if err != nil {
			log.Fatal(err)
		}
		result.LastInsertId()
	}

}

func SelectId(db *sql.DB, id int) Post {
	var p Post
	err := db.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&p.Id, &p.Created_At, &p.Content, &p.User_id, &p.Number_Like, &p.Number_Dislike)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(p)
	return p
}

func ArrayPost() []Post {
	var L []Post
	rows := data.SelectAllFromTable(data.BDD, "posts")
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Created_At, &p.Content, &p.User_id, &p.Number_Like, &p.Number_Dislike)
		// fmt.Println("\n\n\n\n\n ICI ===== ")
		// fmt.Println(p.Created_At)
		p.User_name = users.SelectId(data.BDD, p.User_id).Name
		p.CategoriePost = categorie.GetCategorieByIdPost(data.BDD, p.Id)
		if err != nil {
			log.Fatal(err)
		}
		L = append(L, p)
	}
	return L
}

func ArrayInt(StrInt string) []int {
	var L []int
	for i := range StrInt {
		L = append(L, int(StrInt[i]-48))
	}
	return L
}

type AllStruct struct {
	Pst      Post
	Category categorie.Categories
}

// type DisplayPost struct {

// 	content: string,
// 	date
// 	username: string,
// 	comments []Comment
// 	]
// }

func creePosttest() int {
	fmt.Println("\n### Creation des posts Test ###")
	if len(ArrayPost()) > 1 {
		fmt.Println("### Post Test Deja Cree###")
		return 1
	}
	CreateNewPost(data.BDD, "Premier Post test", 1, nil)
	CreateNewPost(data.BDD, "Deuxieme Post test", 1, nil)
	fmt.Println("### Post Test Cree###")
	return 0
}

var _ = creePosttest()
