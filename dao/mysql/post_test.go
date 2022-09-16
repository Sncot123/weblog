package mysql

import (
	"testing"
	"web_app/models"
	"web_app/settings"
)

func init() {
	conf := settings.MysqlConf{
		Host:     "43.143.76.231",
		Port:     3306,
		UserName: "root",
		Password: "950629",
		Dbname:   "student",
		MaxConn:  5,
		MaxIdle:  3,
	}
	err := Init(&conf)
	if err != nil {
		panic(err)
	}
}
func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          1009,
		AuthorId:    2,
		CommunityId: 2,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf(" createpost insert record into mysql failed,err:%v\n", err)
	}
	t.Logf(" createpost insert record success")
}
