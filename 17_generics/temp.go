package main

type PostController interface {
	CreatePost() error
	GetAllPosts() error
}

type UserController interface {
	CreateUser() error
	GetAllUsers() error
}

type Controllers interface {
	PostController | UserController
}
