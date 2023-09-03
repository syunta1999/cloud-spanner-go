package controllers

import (
	"log"

	"cloud-spanner-go/usecases"
)

type UserController struct {
	Interactor *usecases.UserInteractor
}

func NewUserController(interactor *usecases.UserInteractor) *UserController {
	return &UserController{Interactor: interactor}
}

func (c *UserController) GetUsers() error {
	users, err := c.Interactor.GetUsers()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
		return err
	}
	for _, user := range users {
		log.Printf("User ID: %s, Name: %s,Email: %s", int(user.ID), user.Name, user.Email)
	}
	return nil
}

func (c *UserController) CreateUsers() error {
	err := c.Interactor.CreateUsers()
	if err != nil {
		log.Fatalf("Failed to insert users: %v", err)
		return err
	}
	return nil
}
