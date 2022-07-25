package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	schema, _ := c.readDB()
	if _, ok := schema.Users[userEmail]; ok {
		id := uuid.New().String()
		newPost := Post{
			ID:        id,
			CreatedAt: time.Now().UTC(),
			UserEmail: userEmail,
			Text:      text,
		}
		schema.Posts[id] = newPost
		c.updateDB(schema)
		return newPost, nil
	}
	return Post{}, errors.New("There was an error creating the post")
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	return []Post{}, nil
}

func (c Client) DeletePost(id string) error {
	schema, _ := c.readDB()
	if _, ok := schema.Posts[id]; ok {
		delete(schema.Posts, id)
		c.updateDB(schema)
		return nil
	}
	return errors.New("Error when deleting, ID doesnt exist")
}
