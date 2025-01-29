package models

import (
	"errors"
	"fmt"

	"example.com/api/db"
	"example.com/api/utils"
)
 

type User struct {
	ID       int64
	Email    string
	Password string
}

func (u *User) Save() error {
query := `INSERT INTO users(email, password) VALUES (?, ?)`
stmt,err := db.DB.Prepare(query)
if err != nil {
	fmt.Println(err)
	return err
}
defer stmt.Close()
 hashedPassword,err:= utils.HashPassword(u.Password)
 if err != nil {
	 fmt.Println(err)
	 return err
 }
result,err:= stmt.Exec(u.Email, hashedPassword)
if err != nil {
	fmt.Println(err)
	return err
}
id,err := result.LastInsertId()
if err != nil {
	fmt.Println(err)
	return err
}
u.ID = id
return err
}

func (u * User) Authenticate() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	
	row := db.DB.QueryRow(query ,u.Email) 
	var retrivedPassword  string
	err := row.Scan(&u.ID,&retrivedPassword) 
	if err != nil { 
		fmt.Println(err)
		return err
	} 
	isValid := utils.ComparePasswords(u.Password,retrivedPassword)
 
	if !isValid {
		fmt.Println("Invalid password")
		return errors.New("invalid password")
	}
	
	return nil

}