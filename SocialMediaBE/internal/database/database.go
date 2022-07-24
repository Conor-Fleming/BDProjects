package database

import (
	"encoding/json"
	"errors"
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

//unexported function updateDB
//saves data in given databaseSchema to filepath specified in Client
//will overwrite what existed previously
func (c Client) updateDB(db databaseSchema) error {
	file, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.filepath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

//unexported function readDB
//returns new databaseSchema populated with latest data from disc
func (c Client) readDB() (databaseSchema, error) {
	var result databaseSchema
	data, err := os.ReadFile(c.filepath)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	//check map if email already exists, if so return error?
	schema, _ := c.readDB()
	if val, ok := schema.Users[email]; ok {
		return val, errors.New("user already exists")
	}
	//if the email is unique then add to map and call updateDB()
	//set created time with 'time.Now().UTC()'
	newUser := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	schema.Users[email] = newUser
	c.updateDB(schema)
	return newUser, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	schema, _ := c.readDB()
	if val, ok := schema.Users[email]; !ok {
		return val, errors.New("user doesn't exist")
	}
	updateUser := User{
		CreatedAt: schema.Users[email].CreatedAt,
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	schema.Users[email] = updateUser
	c.updateDB(schema)
	return updateUser, nil

}

func (c Client) GetUser(email string) (User, error) {
	schema, _ := c.readDB()
	if val, ok := schema.Users[email]; !ok {
		return val, errors.New("user doesn't exist")
	}
	result := schema.Users[email]
	return result, nil
}

func (c Client) DeleteUser(email string) error {
	schema, _ := c.readDB()
	if _, ok := schema.Users[email]; !ok {
		return errors.New("user doesn't exist")
	}
	delete(schema.Users, email)
	c.updateDB(schema)
	return nil
}

func NewClient(path string) Client {
	newCli := Client{
		filepath: path,
	}
	return newCli
}
