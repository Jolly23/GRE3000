package database

import (
	"GRE3000/types"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

func (db *Database) FindUserByToken(token string) *types.User {
	row := db.conn.QueryRow(`select user_id, username, avatar, email, url, signature from users where token=$1`, token)

	var user types.User
	if err := row.Scan(&user.ID, &user.Username, &user.Avatar, &user.Email, &user.URL, &user.Signature); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	user.Token = token
	return &user
}

func (db *Database) AuthUser(username, password string) (string, bool) {
	row := db.conn.QueryRow(`select token from users where username=$1 and password=$2`, username, password)

	var token string
	if err := row.Scan(&token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		panic(err)
	}
	return token, true
}

func (db *Database) HasUsername(username string) bool {
	row := db.conn.QueryRow(`select username from users where username=$1`, username)

	if err := row.Scan(&username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		panic(err)
	}
	return true
}

func (db *Database) SaveUsername(username, password string) (int, string, error) {
	token := uuid.New().String()
	row := db.conn.QueryRow(`insert into users(username,password,token) values ($1,$2,$3) returning user_id`, username, password, token)
	var id int
	if err := row.Scan(&id); err != nil {
		return -1, "", err
	}
	return id, token, nil
}

//func (c *StgAutoTransfer) isExistBNBProcessingFromContractToCex() (id int64) {
//	row := c.db.GetDatabaseConn().QueryRow(`select record_id from record_bank_transfer where token=$1 and status=$2 and sender_bank=$3 and recipient_bank=$4`, "BNB", types.TransferStatusProcessing, TransHolderBscContract, TransHolderBINANCE)
//
//	if err := row.Scan(&id); err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return -1
//		}
//		panic(err)
//	}
//	return id
//}
//func (c *StgAutoTransfer) loadTokenInfo() (ret []*types.TokenContract) {
//	rows, err := c.db.GetDatabaseConn().Query(`select address, decimals, token from support_swap_token where chain_id=$1`, c.config.BscNode.ChainID.Uint64())
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var _addr string
//		var _decimals uint8
//		var _token types.Token
//		if err = rows.Scan(&_addr, &_decimals, &_token); err != nil {
//			panic(err)
//		}
//		addr := common.HexToAddress(_addr)
//		ret = append(ret, &types.TokenContract{Token: _token, Contract: addr, Decimals: _decimals})
//	}
//	return ret
//}
//
//func (c *StgAutoTransfer) loadBossAccount() *eoa.EOA {
//	row := c.db.GetDatabaseConn().QueryRow(`select ks_data from ethereum_accounts where role=100`)
//
//	var _ks string
//	if err := row.Scan(&_ks); err != nil {
//		panic(err)
//	}
//
//	return eoa.KSDecrypt(_ks, c.config.BscNode.ChainID)
//}
