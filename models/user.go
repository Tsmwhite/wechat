package model

type User struct {
	Id           int64  `json:"id"`
	Uuid         string `json:"uuid"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Mail         string `json:"mail"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
	Passlook     string `json:"passlook"`
	Headimg      string `json:"headimg"`
	RoleId       int64  `json:"role_id"`
	IsHid        int64  `json:"is_hid"`
	IsDel        int64  `json:"is_del"`
	RegisterTime string `json:"register_time"`
	LoginTime    string `json:"login_time"`
	UpdateTime   string `json:"update_time"`
	RegisterIp   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
}

func NewUser() *User {
	return new(User)
}

func GetUserByAccount(account string) *User {
	user := NewUser()
	DB.Raw("SELECT * FROM `users` WHERE `mobile` = ?", account).Scan(user)
	return user
}

func (u *User) CheckMemberByUuid(uuid string) bool {
	DB.Raw("SELECT * FROM `users` WHERE `uuid` = ?", uuid).Scan(u)
	return u.Uuid == uuid
}

func (u *User) GetUuid() string {
	return u.Uuid
}

func (u *User) GetMemberKey() int {
	return int(u.Id)
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetInfo() map[string]string {
	return nil
}

func (u *User) Update(map[string]string) {

}

func (u *User) Exist() bool {
	return u.Id > 0
}
