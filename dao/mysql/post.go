package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

func CreatePost(c *models.Post) (err error) {
	sqlStr := "INSERT INTO post(post_id,title,content,author_id,community_id)VALUES(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, c.ID, c.Title, c.Content, c.AuthorId, c.CommunityId)
	return
}

func GetPostById(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := "SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE post_id=?"
	err = db.Get(data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no objective data in db")
			return
		}
	}
	return
}
func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := "SELECT user_id,username FROM user WHERE user_id=?"
	err = db.Get(user, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no author in db", zap.Int64("id", id), zap.Error(err))
			return
		}
	}
	return
}
