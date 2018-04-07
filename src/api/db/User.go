package db

// User структура пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SelectUser info
func SelectUser(id string) (*User, error) {
	var u User

	row := db.QueryRow("SELECT id,name FROM `user` WHERE id=?", id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

// Insert user info
func (u *User) Insert() error {
	r, err := db.Exec("INSERT INTO `user`(`name`) VALUES (?)", u.Name)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	return nil
}

// Update user info
func (u *User) Update() error {
	_, err := db.Exec("UPDATE `user` SET `name`=? WHERE `id`=?", u.Name, u.ID)
	return err
}

// Delete user info
func (u *User) Delete() error {
	_, err := db.Exec("DELETE FROM `user` WHERE `id`=?", u.ID)
	return err
}
