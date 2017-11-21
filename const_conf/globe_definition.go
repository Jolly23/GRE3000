package const_conf

import "time"

const (
	LogMarkWord        = "IncrWordIdMark"
	LogMarkWordFmt     = "IncrWordIdMark: %d"
	CookieExpire       = 60 * 60 * 24 * 365
)

const (
	Ok int = 0
)

type UserWordsJson struct {
	Id         int
	Word       string
	Means      string
	CountMarks int
	LastMark   string
}

type RawWord struct {
	Id         int
	Word       string
	Means      string
	CountMarks int
	LastMark   time.Time
}