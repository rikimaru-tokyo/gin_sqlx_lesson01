package model

import (
	"time"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson01/myDatabase"
)

//データの受け手となる構造体を宣言
//`db`タグは必須。
//タグの値はDBテーブルのカラム名に合わせること。
type RankResult struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

// GetRanksAll ranksテーブルデータ全出力
func GetRanksAll() ([]RankResult, error) {
	//データベース接続
	db := myDatabase.DbInit()

	//SQL文字列定義
	sql := `SELECT * FROM ranks;`

	//データの受け手となる構造体をポインタにしてSelect()の第1引数に渡す。
	//Select()の第2引数はSQL文字列をそのまま渡す。
	//第3引数以降は、プレースフォルダのパラメータなどを入れる。
	var ranks []RankResult
	err := db.Select(&ranks, sql)

	if err != nil {
	 	return nil, err
	}

	return ranks, nil
}

// InsertRank() プレースフォルダ「:変数名」として使いたい場合
// ranksテーブルに{name:"Platinum"}を新規登録します。
func InsertRank() (int, error) {
	db := myDatabase.DbInit()
	now := time.Now()

	sql := `INSERT INTO ranks 
				(id, name, created_at, updated_at)
			VALUES
				(:id, :name, :cr, :up);`

	// 「:変数名」という形式でプレースフォルダを使いたい場合、NamedExecというメソッドを使います。
	//  第1引数はSQL文字列、第2引数は「map[string]interface(連想配列)」
	//  連想配列インデックス名は、SQL文中の「:変数名」に合わせます。
	//      [参考資料: https://github.com/jmoiron/sqlx の READMEより]
	rows, err := db.NamedExec(sql,  map[string]interface{}{
		"id":4, "name": "Platinum", "cr": now, "up": now,
	})
	if err != nil {
		return 0, err
    }

    // INSERT件数カウント
    r, err := rows.RowsAffected()
	if err != nil {
		return 0, err
    }

   return int(r), nil

}