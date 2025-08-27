package database

import (
	"database/sql"
	"fmt"
	"log"
)

// Post represents a blog post
type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreatePost creates a new post in the database
func CreatePost(title, content string, userID int) (int, error) {
	query := `
		INSERT INTO posts (title, content, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err := DB.QueryRow(query, title, content, userID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create post: %v", err)
	}

	log.Printf("Post created successfully with ID: %d", id)
	return id, nil
}

// GetPosts retrieves all posts from the database
func GetPosts() ([]Post, error) {
	query := `
		SELECT id, title, content, user_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %v", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %v", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating posts: %v", err)
	}

	return posts, nil
}

// GetPostByID retrieves a specific post by its ID
func GetPostByID(id int) (*Post, error) {
	query := `
		SELECT id, title, content, user_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post Post
	err := DB.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("failed to fetch post: %v", err)
	}

	return &post, nil
}

// UpdatePost updates an existing post
func UpdatePost(id int, title, content string) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
	`

	result, err := DB.Exec(query, title, content, id)
	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}

	log.Printf("Post updated successfully with ID: %d", id)
	return nil
}

// DeletePost deletes a post by its ID
func DeletePost(id int) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}

	log.Printf("Post deleted successfully with ID: %d", id)
	return nil
}
