package main

import (
	"fmt"
	"todoapp/app/controllers"
	"todoapp/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	/*
		fmt.Println(models.Db)

		u := &models.User{}
		u.Name = "test3"
		u.Email = "test3@example.com"
		u.PassWord = "test1234"
		fmt.Println(u)

		u.CreateUser()
	*/
	/*

		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("study")
	*/

	//t, _ := models.GetTodo(1)
	//fmt.Println(t)

	//user, _ := models.GetUser(3)
	//user.CreateTodo("swimming")

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
		t.Content = "hangout"
		t.UpdateTodo()
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
		t.DeleteTodo()
	*/

	// Server の設定

	fmt.Println(models.Db)
	controllers.StartMainServer()

	/*
		user, _ := models.GetUserByEmail("test@example.com")
		fmt.Println(user)

		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(session)

		valid, err := session.CheckSession()
		fmt.Println(valid)
	*/
}
