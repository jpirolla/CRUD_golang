package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil

}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
