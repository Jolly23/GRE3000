package types

import "time"

type RawWord struct {
	Word string `json:"w"`
	Mean string `json:"m"`
}

type UserWord struct {
	WordID      int        `json:"i"`
	Word        string     `json:"w"`
	Mean        string     `json:"m"`
	CountMarks  int        `json:"c"`
	LastMarkAge string     `json:"t"`
	LastMark    *time.Time `json:"-"`
}
