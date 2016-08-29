package trakitapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	Id          int      `json:"uuid"`
	Username    string   `json:"username"`
	Tenant      string   `json:"tenant"`
	Password    string   `json:"password"`
	Permissions []string `json:"permissions"`
	Data        UserData `json:"data"`
}

type UserData struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Alias            string `json:"alias"`
	Description      string `json:"description"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Group            string `json:"group"`
	RoleId           int    `json:"role_id"`
	FailedLoginCount []int  `json:"failedLoginCount"`
	LastUpdate       string `json:"last_update"`
}

func GetUser(id int) (User, error) {
	var user User

	res := Get("user/" + strconv.Itoa(id))

	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}

	fmt.Println(user)

	return user, nil
}
