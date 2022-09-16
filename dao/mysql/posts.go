package mysql

import (
	"strings"
	"web_app/models"

	"github.com/jmoiron/sqlx"
)

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := "SELECT post_id,title,content,author_id,community_id,create_time FROM post LIMIT ?,?"
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	return
}
func GetPostList2(ids []string) (data []*models.Post, err error) {
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time
		FROM post
		WHERE id IN(?)
		ORDER BY FIND_IN_SET(post_id,?)
		`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&data, query, args...)
	return
}
