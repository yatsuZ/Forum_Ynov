package alert

//Here we're using enumerations to decouple the state of the game with web implementation
const (
	StatePasswordOk     = 100
	StatePasswordNotOk  = 101
	Stateuserexist      = 102
	StateverificationOk = 103

	StateSigneinOK   = 200
	StateSigninNotOK = 201
	StateSignUpOk    = 202
	StateSignUpNotOK = 203
	StatePostOK      = 205

	StateUserLoggedIn = 300
)

var gameStateText = map[int]string{
	StatePasswordOk:     "WRONG PASSWORD OR EMAIL NIQUE LES PD",
	StatePasswordNotOk:  "Password and Verification are not similar ",
	Stateuserexist:      "username or email already used",
	StateverificationOk: "verificationok",
	StatePostOK:         "POST CREATED SUCESSED",

	StateSignUpOk:     "You won, Want to play again ?",
	StateSignUpNotOK:  "You lost, Play again !",
	StateSigneinOK:    "Good luck !",
	StateSigninNotOK:  "You must provide a valide input",
	StateUserLoggedIn: "Hello ",
}

func WebSiteStateText(code int) string {
	return gameStateText[code]
}
