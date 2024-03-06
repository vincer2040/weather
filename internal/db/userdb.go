package db

import "github.com/vincer2040/weather/internal/types"

func (db *DB) CreateUserTable() error {
	stmt := `
    CREATE TABLE IF NOT EXISTS
    users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        email TEXT NOT NULL,
        firstname TEXT NOT NULL,
        lastname TEXT NOT NULL,
        password TEXT NOT NULL
    )
    `
	_, err := db.exec(stmt)
	if err != nil {
		return err
	}

	usernameIndexStmt := `
    CREATE UNIQUE INDEX IF NOT EXISTS username_index ON users(username)
    `

	_, err = db.exec(usernameIndexStmt)
	if err != nil {
		return err
	}

	emailIndexStmt := `
    CREATE UNIQUE INDEX IF NOT EXISTS email_index ON users(email)
    `

	_, err = db.exec(emailIndexStmt)
	return err
}

func (db *DB) InsertUser(user *types.User) (int, error) {
	stmt := `
    INSERT INTO
    users(username, email, firstname, lastname, password)
    VALUES(?, ?, ?, ?, ?)
    RETURNING id
    `
	row := db.queryRow(stmt, user.Username, user.Email, user.FirstName, user.LastName, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) GetUserById(id int) (*types.User, error) {
	stmt := `
    SELECT * FROM users
    WHERE id = ?
    `
	row := db.queryRow(stmt, id)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password)
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
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUserByEmail(email string) (*types.User, error) {
	stmt := `
    SELECT * FROM users
    WHERE email = ?
    `
	row := db.queryRow(stmt, email)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUserByUsernameOrEmail(usernameOrEmail string) (*types.User, error) {
	stmt := `
    SELECT * FROM users
    WHERE username = ? OR email = ?
    `
	row := db.queryRow(stmt, usernameOrEmail)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password)
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
