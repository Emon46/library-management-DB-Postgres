package pkg
type User struct {
	ID int
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"password"`
	PhoneNo string `json:"phone_no"`
}

var Users []User
func (user *User) CreateUser() bool {

	Users = append(Users, User{
		ID:       user.ID,
		Name:     user.Name,
		Mail:     user.Mail,
		Password: user.Password,
		PhoneNo:  user.PhoneNo,
	})
	return true
}