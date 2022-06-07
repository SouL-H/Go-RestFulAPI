package handlers

import (
	. "api/dataloaders"
	. "api/models"
	"encoding/json"
	"fmt"

	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//Page Object
	page := Page{ID: 7, Name: "Useers", Description: "Users List", URI: "/users"}
	//Data Loaders
	users := LoadUsers()
	interests := LoadInterests()
	InterestMappings := LoadInterestMappings()

	//Process

	var newUsers []User

	for _, user := range users {
		for _, interestMapping := range InterestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}

		}
		newUsers = append(newUsers, user)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
	fmt.Println(data)

}
