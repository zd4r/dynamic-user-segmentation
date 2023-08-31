package experiment

type Experiment struct {
	UserId    int `db:"user_id"`
	SegmentId int `db:"segment_id"`
}
