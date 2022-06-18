package model

import (
	"github.com/rikimaru-tokyo/gin_sqlx_lesson01/myDatabase"
)

//データの受け手となる構造体を宣言
type MembersResult struct {
	MemberID   int     `db:"member_id"`
	MemberName string  `db:"member_name"`
	Rank       string  `db:"rank_name"`
}


// GetMembersWithRanks テーブル結合かつプレースフォルダを使った結果を取得
func GetMembersWithRanks(memberId int, rankName string)([]MembersResult, error){
	//データベース接続
	db := myDatabase.DbInit()

	//SQL文字列定義
	sql := `SELECT 
				m.id as member_id, m.name as member_name, r.name as rank_name
			FROM
				members as m 
			INNER JOIN
				ranks as r ON m.rank_id = r.id
			WHERE
				m.id = ? OR r.name = ?;
	`
	var member []MembersResult

	// 第3引数以降は、プレースフォルダ「?」がついた順番にあわせてパラメータ変数を入れておく。
	err := db.Select(&member, sql, memberId, rankName)

	if err != nil {
	 	return nil, err
	}

	return member, nil

}