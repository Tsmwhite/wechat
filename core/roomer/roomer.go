package roomer

import (
	"wechat/core/member"
)

type Roomer interface {
	GetMembers() []string
	Join(member.Member) error
	Quit(member.Member) error
}
