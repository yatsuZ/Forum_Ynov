package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/alert"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/comment"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/data"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/posts"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/users"
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/utils"
)

type alldata struct {
	Posts []posts.Post
	Alert string
}

//Initializing the Data
var (
	U users.User
	d = alldata{
		Posts: posts.ArrayPost(),
		Alert: "",
	}
)

func Redirection(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
	}
	if U.Name == "" {
		http.Redirect(rw, r, "/", http.StatusSeeOther)

	} else {
		http.Redirect(rw, r, "/homeuser", http.StatusSeeOther)

	}
}

func Error(rw http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	err2 := temp.ExecuteTemplate(rw, "Error", nil)
	if err != nil {
		log.Fatal(err2)
	}
}

func Acceuil(rw http.ResponseWriter, r *http.Request) {
	if U.Name != "" {
		http.Redirect(rw, r, "/homeuser", http.StatusSeeOther)
	}
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	d.Alert = alert.WebSiteStateText(300)
	// fmt.Println("\n\n\nce que j'envoie dans le html == ")
	// fmt.Println(d)
	d.Posts = posts.ArrayPost()
	err2 := temp.ExecuteTemplate(rw, "home", d)
	if err != nil {
		log.Fatal(err2)
	}
}

func VerificationSignUp(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	if r.Method == "POST" {
		password := r.FormValue("password")
		if password != r.FormValue("verification") {
			//Alert le mot de passe et la verification est incorecte
			d.Alert = alert.WebSiteStateText(101)
			fmt.Println("\nErreur le password et different de la verification !!")
		} else {
			d.Alert = alert.WebSiteStateText(103)
			fmt.Println("\nLe mot de passe est la verification est similaire c'est bon :) ")
			email := r.FormValue("email")
			/////////////////////////////////////////////////////////////////////////////
			pseudo := r.FormValue("name")
			nbr, err3 := users.AddUser(data.BDD, pseudo, email, password)
			if err3 != nil {
				log.Fatal(err3)
			}
			if nbr != 1 {
				U = users.SelectName(data.BDD, pseudo)
				http.Redirect(rw, r, "/homeuser", http.StatusSeeOther) //<--quelqun sait changer le redirect
			} else {
				//Le pseudo ou l'email est deja utiliser
				d.Alert = alert.WebSiteStateText(102)

			}
		}
	}
	http.Redirect(rw, r, "/signup", http.StatusSeeOther) //<--quelqun sait changer le redirect
}

func SignUp(rw http.ResponseWriter, r *http.Request) {
	if U.Name != "" {
		http.Redirect(rw, r, "/homeuser", http.StatusSeeOther)
	}
	//inscription
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	err2 := temp.ExecuteTemplate(rw, "signUp", d)
	if err != nil {
		log.Fatal(err2)
	}
}

///////////////////////////////////////////////////////////////////////////////

func VerificationSignIn(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		var vide users.User
		userName := r.FormValue("name")
		UserBDD := users.SelectUserByName(data.BDD, userName)
		password := r.FormValue("password")

		c, errC := r.Cookie("session") // dans 1mn
		sID := uuid.NewV4()
		if errC != nil {
			c = &http.Cookie{
				Name:  "session",
				Value: sID.String(),
			}
		}
		http.SetCookie(rw, c)

		// fmt.Printf("UserName %v , Password: %v\n", userName, password)

		if UserBDD.Name != vide.Name && utils.ComparePassowrds(UserBDD.Pasword, []byte(password)) { //NON pas ici pardon
			fmt.Println("\nTout es bon user trouver et bon mot de passe")
			U = UserBDD
			http.Redirect(rw, r, "/homeuser", http.StatusSeeOther)
		} else {
			// llaaaaalaalalalala
			d.Alert = alert.WebSiteStateText(100)

			fmt.Println("\nL'user n'est pas bon ou le mot de passe")
			http.Redirect(rw, r, "/signin", http.StatusSeeOther)

		}

	}

}
func SignIn(rw http.ResponseWriter, r *http.Request) { //connexion
	if U.Name != "" {
		http.Redirect(rw, r, "/homeuser", http.StatusSeeOther)
	}
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	err = temp.ExecuteTemplate(rw, "signIn", d)
	if err != nil {
		log.Fatal(err)
	}

}

func Dashboard(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	dataPost := posts.ArrayPost()
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	err2 := temp.ExecuteTemplate(rw, "dashboard", dataPost)
	if err != nil {
		log.Fatal(err2)
	}

}

//Ici on suppose que le formulaire a /post dans la propriete action
func NewPost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	if r.Method == "POST" {
		text := r.FormValue("postText")

		choix := posts.ArrayInt(r.FormValue("categories"))

		posts.CreateNewPost(data.BDD, text, 1, choix)
		//Votre Post a etais crée
		d.Alert = alert.WebSiteStateText(205)
		// fmt.Println(choix)
		http.Redirect(rw, r, "/dashboard", http.StatusSeeOther)
	}
}

func UserPage(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}

	err2 := temp.ExecuteTemplate(rw, "userpage", U)
	if err != nil {
		log.Fatal(err2)
	}
}

func Comment(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}

	aStringPostId := r.URL.RequestURI()
	fmt.Println("url = ", aStringPostId)
	long := len("/comment/")
	ID_Post_Present, err := strconv.Atoi(aStringPostId[long:])
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	}
	fmt.Println(ID_Post_Present)

	d.Posts[ID_Post_Present-1].AllComment = comment.ArrayCommentToPost(d.Posts[ID_Post_Present-1])
	err2 := temp.ExecuteTemplate(rw, "comment", d.Posts[ID_Post_Present-1])
	if err != nil {
		log.Fatal(err2)
	}
}

func NEWComment(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	if r.Method == "POST" {
		aStringPostId := r.URL.RequestURI()
		fmt.Println("url = ", aStringPostId)
		long := len("/NEWcomment/")
		ID_Post_Present, err := strconv.Atoi(aStringPostId[long:])
		if err != nil {
			log.Fatal(err)
		}
		text := r.FormValue("commentText")
		comment.CreateNewComment(data.BDD, text, d.Posts[ID_Post_Present-1])
		//Votre Comment a etais crée
		// d.Alert = alert.WebSiteStateText(205)
		// fmt.Println(choix)
		d := strconv.Itoa(ID_Post_Present)
		fmt.Println("http://localhost:5555/comment/" + d)
		http.Redirect(rw, r, "/comment/"+d, http.StatusSeeOther)
	}
}

func HomeUser(rw http.ResponseWriter, r *http.Request) {
	if U.Name == "" {
		http.Redirect(rw, r, "/error", http.StatusSeeOther) //<--quelqun sait changer le redirect
		return
	}
	temp, err := template.ParseGlob("../htmlAssets/*.html")
	if err != nil {
		log.Fatal(err)
	}

	err2 := temp.ExecuteTemplate(rw, "homeUser", d)
	if err != nil {
		log.Fatal(err2)
	}
}
