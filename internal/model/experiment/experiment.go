package experiment

import (
	"database/sql"
)

type Experiment struct {
	UserId    int          `db:"user_id"`
	SegmentId int          `db:"segment_id"`
	ExpireAt  sql.NullTime `db:"expire_at"`
}
