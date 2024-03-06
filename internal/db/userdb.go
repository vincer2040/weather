package db

import "github.com/vincer2040/weather/internal/types"

func (db *DB) CreateUserTable() error {
	stmt := `
    CREATE TABLE IF NOT EXISTS
    users(
        ID INTEGER AUTO INCREMENT,
        username TEXT NOT NULL,
        firstname TEXT NOT NULL,
        lastname TEXT NOT NULL,
        password TEXT NOT NULL,
        PRIMARY KEY (ID, username)
    )
    `
	_, err := db.exec(stmt)
	return err
}

func (db *DB) InsertUser(user *types.User) (int, error) {
	stmt := `
    INSERT INTO
    users(username, firstname, lastname, password)
    VALUES(?, ?, ?, ?)
    RETURNING ID
    `
	row := db.queryRow(stmt, user.Username, user.FirstName, user.LastName, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func (db *DB) GetUserById(id int) (*types.User, error) {
	stmt := `
    SELECT * FROM users
    WHERE ID = ?
    `
	row := db.queryRow(stmt, id)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUserByUsername(username string) (*types.User, error) {
	stmt := `
    SELECT * FROM users
    WHERE username = ?
    `
	row := db.queryRow(stmt, username)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (db *DB) DropUserTable() error {
	stmt := `
    DROP TABLE IF EXISTS users
    `
	_, err := db.exec(stmt)
	return err
}
