package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/jpirolla/CRUD_golang/src/configuration/rest_err"
)

// regras de negocio do dominio de usuario
// retirar as tags pois o domanin nao pode ser exportavel pois tem
// Esse arquivo define como um usuário é representado dentro da nossa aplicação

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	// md5 com hash

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil)) //trocar o valor da senha pela senha encriptada
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr              // string é o id do usuario e o dominio a ser alterado
	FindUser(string) (*UserDomain, *rest_err.RestErr) // caso de erro não posso retornar um obj vazio
	DeleteUser(string) *rest_err.RestErr
}
