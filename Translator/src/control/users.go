package control

import (
	"../model"
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dashboard")

	renderTemplate("home", w, r, func(data TemplateData) TemplateData {
		data.LanguageCompletion = model.GetLanguageCompletion()
		return data
	})
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		name := r.FormValue("name")
		user := &model.User{
			Email:    email,
			Name:     name,
			Password: "",
			Secret:   "",
		}
		user.Save()

		http.Redirect(w, r, "/users", 303)
	} else {
		renderTemplate("users", w, r, func(data TemplateData) TemplateData {
			data.Users = model.GetUsers()
			return data
		})
	}
}

func UsersAddHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate("users_add", w, r, nil)
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r)

	if r.Method == "POST" {
		user.Name = r.FormValue("name")
		language := r.FormValue("language")
		if language != "" {
			user.Language = language
		}
		user.Save()

		http.Redirect(w, r, "/home", 303)
	} else {
		renderTemplate("account", w, r, func(data TemplateData) TemplateData {
			return data
		})
	}
}

func AccountReclaimHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate("account_reclaim", w, r, nil)
}

func SetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r)

	if r.Method == "POST" {
		password := r.FormValue("password")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err == nil {
			user.Password = string(hash)
			user.Save()
		}
		http.Redirect(w, r, "/account", 303)
	} else {
		renderTemplate("account_set_password", w, r, func(data TemplateData) TemplateData {
			return data
		})
	}
}
