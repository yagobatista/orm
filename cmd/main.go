package main

import (
	"context"
	"fmt"
	"orm/db"
)

type User struct {
	db.ModelDB

	Name string `orm:"name"`
}

func main() {
	db.Init()
	// db.AutoMiograte()

	var err error
	ctx := context.TODO()
	userModel := db.GetModel(User{})

	// user := User{
	// 	Name: "yaya",
	// }
	// err = userModel.Insert(ctx, user)
	// if err != nil {
	// 	panic(err)
	// }

	var newUser User

	err = userModel.
		Select(
			{Atribute: "name", Operator: db.Eq, Value: "yaya"},
		).Find(ctx, &newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println("User ----> ", newUser)
}
