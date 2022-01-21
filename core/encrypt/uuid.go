package encrypt

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func CreateUuid() string {
	uid := uuid.NewV4().String()
	return strings.ReplaceAll(uid, "-", "")
}
