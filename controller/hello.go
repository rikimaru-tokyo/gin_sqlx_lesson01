package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson01/model"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"Hello": "world"})
}

// RankGet rankテーブル 全体を出力
func RankGet(c *gin.Context) {

	// 参照：model/ranks.go
	result, err := model.GetRanksAll()
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

// GetMembers プレースフォルダを使った検索
func GetMembers(c *gin.Context) {

	var m int = 2
	var r string = "Gold"

	// 参照：model/members.go
	result, err := model.GetMembersWithRanks(m, r)
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

// PostRank プレースフォルダ(別形式)を使ったINSERT実行
func PostRank(c *gin.Context){

	// 参照：model/ranks.go
	result, err := model.InsertRank()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"Affected Rows": result})
}