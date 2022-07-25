package database

import (
	"errors"
	"time"
)

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	//check map if email already exists, if so return error?
	schema, _ := c.readDB()
	if val, ok := schema.Users[email]; !ok {
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
