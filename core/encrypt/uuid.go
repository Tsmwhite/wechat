package encrypt

import (
	uuid "github.com/satori/go.uuid"
	"sort"
	"strings"
)

func CreateUuid() string {
	uid := uuid.NewV4().String()
	return strings.ReplaceAll(uid, "-", "")
}

func CreteUuidsKey(uuids []string) string {
	sortSlice := sort.StringSlice(uuids)
	sort.Stable(sortSlice)
	sortStr := strings.Join(sortSlice, ",")
	return Md5(sortStr)
}
