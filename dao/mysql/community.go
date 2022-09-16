package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlSTR := "SELECT community_id,community_name FROM community"
	err = db.Select(&communityList, sqlSTR)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}
func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	sqlStr := "SELECT community_id,community_name,introduction,create_time FROM community WHERE community_id=?"
	err = db.Get(data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no objective community in db")
			return nil, err
		}
	}
	return
}
