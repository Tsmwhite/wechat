package model

type User struct {
	Id           int64  `json:"id"`
	Uuid         string `json:"uuid"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Mail         string `json:"mail"`
	Avatar       string `json:"avatar"`
	RoleId       int64  `json:"role_id"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
	PassLook     string `json:"pass_look"`
	Status       int64  `json:"status"`
	RegisterTime int64  `json:"register_time"`
	LoginTime    int64  `json:"login_time"`
	UpdateTime   int64  `json:"update_time"`
	RegisterIp   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
	IsDel        int64  `json:"is_del"`
}

type ShowAppUser struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Mail   string `json:"mail"`
	Avatar string `json:"avatar"`
	RoleId int64  `json:"role_id"`
	Token  string `json:"token"`
}

func NewUser() *User {
	return new(User)
}

func GetUserByUuid(uuid string) *User {
	u := NewUser()
	DB.Raw("SELECT * FROM `users` WHERE `uuid` = ? AND `is_del` = 0 ", uuid).Scan(u)
	return u
}

func (u *User) ShowAppUser() *ShowAppUser {
	res := new(ShowAppUser)
	res.Name = u.Name
	res.Mobile = u.Mobile
	res.Avatar = u.Avatar
	res.Mail = u.Mail
	res.RoleId = u.RoleId
	return res
}

func (u *User) Create() error {
	res := DB.Create(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetUserByAccount(account string) *User {
	user := NewUser()
	DB.Raw("SELECT * FROM `users` WHERE (`mobile` = ? OR `mail` = ?) AND `is_del` = 0", account, account).Scan(user)
	return user
}

func (u *User) CheckMemberByUuid(uuid string) bool {
	DB.Raw("SELECT * FROM `users` WHERE `uuid` = ? AND `is_del` = 0 ", uuid).Scan(u)
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
	return u.Id > 0 && u.IsDel == IsNoDel
}
