package database

import (
	"encoding/json"
	"os"
	"time"
)

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

type Client struct {
	filepath string
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.filepath, dat, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.filepath)
	if err != nil {
		c.createDB()
		return err
	}
	return nil
}

func NewClient(path string) Client {
	newCli := Client{
		filepath: path,
	}
	return newCli
}
