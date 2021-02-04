package Models

type User struct {
	Id uint `json:"id"`
	Name string `json:"string"`
	Email string `json:"string"`
	Phone string `json:"string"`
	Address string `json:"string"`
}

func (b* User) TableName() string {
	return "user"
}