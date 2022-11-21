package user

import (
	"encoding/json"
	"time"
)

const ()

func NewUser() *User {
	return &User{}
}

// patch JSON {a: 1, b： 2}， {b:20}  ===> {a:1, b:20}
func ObjectPatch(old, new interface{}) error {
	// {b: 20}
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	// {a:1, b:2}
	// {a:1, b: 20}
	return json.Unmarshal(newByte, old)
}

type User struct {
	user_id         int
	user_name       string
	user_password   string
	last_login_time time.Time
	authority_id    int
}

type UserSet struct {
	Items []*User `json:"items"`
	Total int     `json:"total"`
}

func (us *UserSet) Add(item *User) {
	us.Items = append(us.Items, item)
}
