package member

type Member interface {
	CheckMemberByUuid(string) bool
	GetUuid() string
	GetMemberKey() int
	GetName() string
	GetInfo() map[string]string
	Update(map[string]string)
}
