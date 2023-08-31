package report

import (
	"strconv"
	"time"
)

const (
	AddAction    = "add"
	RemoveAction = "remove"
)

type Record struct {
	UserId      int       `db:"user_id"`
	SegmentSlug string    `db:"segment_slug"`
	Action      string    `db:"action"`
	Date        time.Time `db:"date"`
}

func GetCSVHeader() []string {
	return []string{"user_id", "segment_slug", "action", "date"}
}

func (r *Record) ToCSV() []string {
	return []string{
		strconv.Itoa(r.UserId),
		r.SegmentSlug,
		r.Action,
		r.Date.String(),
	}
}
