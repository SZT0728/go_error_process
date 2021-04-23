package service

import (
	"fmt"
	"go_error_process/dao"
)

type ScoreLevel string

const (
	SCORE_LEVEL_GOOD ScoreLevel = "优"
	SCORE_LEVEL_FIND ScoreLevel = "良"
)

type UserInfo struct {
	Name    string
	Math    float64
	Chinese float64
	English float64
	Level   ScoreLevel
}

func (this *UserInfo) ToString()  {
	fmt.Printf("name=%s   level=%s  chinese=%.2f english=%.2f math=%.2f\n",this.Name,this.Level,this.Chinese,this.English,this.Math)
}

func GetUserInfo(userId int32) (*UserInfo, error) {
	user, err := dao.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	score, err := dao.GetScoreByUserId(userId)
	if err != nil {
		return nil, err
	}

	info := &UserInfo{
		Name:user.Name,
		Math:score.Math,
		Chinese:score.Chinese,
		English:score.English,
	}

	if score.Math > 90 && score.English > 90 && score.Chinese > 90{
		info.Level = SCORE_LEVEL_GOOD
	}else {
		info.Level = SCORE_LEVEL_FIND
	}
	return info, nil
}
