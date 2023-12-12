package database

import (
	"GRE3000/types"
	"database/sql"
)

func (db *Database) InsertWord(word, mean string) {
	_, _ = db.conn.Exec(`insert into vocabulary(word, mean) values ($1, $2)`, word, mean)
}

func (db *Database) LoadRawWords(random bool) []*types.RawWord {
	var (
		rows *sql.Rows
		ret  = make([]*types.RawWord, 0, 3072)
	)
	if random {
		rows, _ = db.conn.Query(`select word, mean from vocabulary order by random()`)
	} else {
		rows, _ = db.conn.Query(`select word, mean from vocabulary order by word_id`)
	}
	defer rows.Close()

	for rows.Next() {
		var word, mean string
		if err := rows.Scan(&word, &mean); err != nil {
			panic(err)
		}
		ret = append(ret, &types.RawWord{Word: word, Mean: mean})
	}
	return ret
}

func (db *Database) LoadRawWordsJson(random bool) map[string]string {
	var (
		rows *sql.Rows
		book = make(map[string]string)
	)
	if random {
		rows, _ = db.conn.Query(`select word, mean from vocabulary order by random()`)
	} else {
		rows, _ = db.conn.Query(`SELECT word, mean FROM vocabulary ORDER BY id OFFSET 30`)
	}
	defer rows.Close()

	for rows.Next() {
		var word, mean string
		if err := rows.Scan(&word, &mean); err != nil {
			panic(err)
		}
		book[word] = mean
	}
	return book
}

func (db *Database) LoadUserWords(userID int, random bool) []*types.UserWord {
	var (
		rows *sql.Rows
		ret  = make([]*types.UserWord, 0, 3072)
	)
	if random {
		rows, _ = db.conn.Query(`select word_id, word, mean, count_mark, last_mark from user_words left join vocabulary using (word_id) where user_id=$1 order by random()`, userID)
	} else {
		rows, _ = db.conn.Query(`select word_id, word, mean, count_mark, last_mark from user_words left join vocabulary using (word_id) where user_id=$1 order by count_mark desc,last_mark desc,word_id`, userID)
	}
	defer rows.Close()

	for rows.Next() {
		var word types.UserWord
		if err := rows.Scan(&word.WordID, &word.Word, &word.Mean, &word.CountMarks, &word.LastMark); err != nil {
			panic(err)
		}
		ret = append(ret, &word)
	}
	return ret
}

func (db *Database) GenerateUserWord(userID int) {
	_, _ = db.conn.Exec(`delete from user_words where user_id=$1`, userID)
	_, _ = db.conn.Exec(`insert into user_words (user_id, word_id) select $1, word_id from vocabulary`, userID)
}

func (db *Database) MarkUserWord(userID, wordID int) error {
	_, err := db.conn.Exec(`update user_words set count_mark=count_mark+1, last_mark=current_timestamp where user_id=$1 and word_id=$2`, userID, wordID)
	return err
}

func (db *Database) DeleteUserWord(userID, wordID int) error {
	_, err := db.conn.Exec(`delete from user_words where user_id=$1 and word_id=$2`, userID, wordID)
	return err
}
