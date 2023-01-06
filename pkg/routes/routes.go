package routes // here will handleFunc

import (
	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/controllers"

	"net/http"
)

var ForumRoutes = func() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../img"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("../htmlAssets"))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("../cssAssets"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("../jsAssets"))))

	http.HandleFunc("/", controllers.Acceuil)
	http.HandleFunc("/signin", controllers.SignIn)       //connexion
	http.HandleFunc("/signup", controllers.SignUp)       //inscruption
	http.HandleFunc("/dashboard", controllers.Dashboard) //inscruption
	http.HandleFunc("/user", controllers.UserPage)
	http.HandleFunc("/post", controllers.NewPost)                   //inscruption
	http.HandleFunc("/verifsignin", controllers.VerificationSignIn) //inscruption
	http.HandleFunc("/verifsignup", controllers.VerificationSignUp) //inscruption
	http.HandleFunc("/error", controllers.Error)                    //inscruption
	http.HandleFunc("/redirection", controllers.Redirection)        //inscruption

	http.HandleFunc("/comment/", controllers.Comment)       //inscruption
	http.HandleFunc("/NEWcomment/", controllers.NEWComment) //inscruption
	http.HandleFunc("/homeuser", controllers.HomeUser)

}
