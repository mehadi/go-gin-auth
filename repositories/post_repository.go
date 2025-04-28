package repositories

import (
	"go-gin-auth/config"
	"go-gin-auth/models"
)

// CreatePost saves a new post to the database
func CreatePost(post models.Post) (models.Post, error) {
	err := config.DB.Create(&post).Error
	return post, err
}

// GetPostByID finds a post by its ID
func GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	err := config.DB.First(&post, id).Error
	return post, err
}

// Delete post
func DeletePost(id uint) error {
	// First check if post exists
	_, err := GetPostByID(id)
	if err != nil {
		return err
	}

	// If post exists, proceed with deletion
	err = config.DB.Delete(&models.Post{}, id).Error
	return err
}

// UpdatePost updates an existing post
func UpdatePost(id uint, post models.Post) (models.Post, error) {
	// First check if post exists
	_, err := GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}

	// Update the post
	err = config.DB.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error
	if err != nil {
		return models.Post{}, err
	}

	// Fetch the updated post
	updatedPost, err := GetPostByID(id)
	return updatedPost, err
}
