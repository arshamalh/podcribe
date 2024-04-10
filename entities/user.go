package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint      `bun:",pk,autoincrement"`
	ChatID      int64     `bun:",unique"`
	Password    []byte    `json:"-"`
	Email       string    `bun:",unique" json:"email"`
	PhoneNumber string    `bun:",unique" json:"phone"`
	CreatedAt   time.Time `bun:",default:current_timestamp"`
	UpdatedAt   time.Time `bun:",default:current_timestamp"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

// TODO: ChatID should not be unique because it can be empty in case of registering from WebUI
// Phone number and Email also can't be unique
// But their uniqueness should be checked in code if they were not empty. (maybe DB provide this feature)
/// https://www.sqlitetutorial.net/sqlite-unique-constraint/#:~:text=SQLite%20UNIQUE%20constraint%20and%20NULL,can%20have%20multiple%20NULL%20values.&text=As%20you%20can%20see%2C%20even,can%20accept%20multiple%20NULL%20values.
// It seems there is no problem using sqlite
