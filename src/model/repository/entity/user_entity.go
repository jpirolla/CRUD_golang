package entity

type UserEntity struct {
	ID       string `bson:" id.omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Age      int8   `bson:"age"`
}
