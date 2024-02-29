package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

// data
type Data struct {
	ID       int
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var data = []Data{
	{ID: 1, Email: "john.doe@example.com", Password: "john", Address: "123 Main St, Cityville, USA", Job: "Software Engineer", Reason: "Sunshine and warmth"},
	{ID: 2, Email: "jane.smith@example.com", Password: "jane", Address: "321 South St, Sydney, Australia", Job: "Mobile Engineer", Reason: "Test Reason Jane"},
	{ID: 3, Email: "michael.johnson@example.com", Password: "michael", Address: "782 West St, Jakarta, Indonesia", Job: "Backend Engineer", Reason: "Test Michael"},
}

var PORT = ":8080"

func main() {
	// routing static
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// routing
	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profileAccount)

	fmt.Println("App is listening on port: ", PORT)
	http.ListenAndServe(PORT, nil)
}

// routing /
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filePath = path.Join("views", "index.html")
		// parsing file
		tpl, err := template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		publicData := getDataList()
		tpl.Execute(w, publicData)
		return
	}

	http.Error(w, "Invalid Method", http.StatusInternalServerError)
}

// routing /profile/
func profileAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusFound)

	case "POST":
		var email = r.FormValue("email_input")
		var password = r.FormValue("password_input")

		var logAccount Data
		var filePath string = path.Join("views", "invalid.html")
		for _, value := range data {
			if email == value.Email && password == value.Password {
				logAccount.Email = value.Email
				logAccount.Address = value.Address
				logAccount.Job = value.Job
				logAccount.Reason = value.Reason

				filePath = path.Join("views", "profile.html")
				break
			}
		}

		// parsing file
		tpl, err := template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, logAccount)

	default:
		http.Error(w, "Invalid Method", http.StatusInternalServerError)
	}
}

func getDataList() []Data {
	var publicData []Data
	var infoData Data
	for _, value := range data {
		infoData.Email = value.Email
		infoData.Password = value.Password
		publicData = append(publicData, infoData)
	}

	return publicData
}
