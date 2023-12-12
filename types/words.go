package types

import "time"

type RawWord struct {
	Word string `json:"word"`
	Mean string `json:"mean"`
}

type UserWord struct {
	WordID     int
	Word       string
	Mean       string
	CountMarks int
	LastMark   *time.Time
}
