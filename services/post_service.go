package services

import (
	"go-gin-auth/models"
	"go-gin-auth/repositories"
)

// CreatePost handles business logic for creating a post
func CreatePost(post models.Post) (models.Post, error) {
	return repositories.CreatePost(post)
}

// DeletePost handles business logic for deleting a post

func DeletePost(id uint) error {
	return repositories.DeletePost(id)
}

// GetPostByID handles business logic for getting a post by ID
func GetPostByID(id uint) (models.Post, error) {
	return repositories.GetPostByID(id)
}

// UpdatePost handles business logic for updating a post
func UpdatePost(id uint, post models.Post) (models.Post, error) {
	return repositories.UpdatePost(id, post)
}
