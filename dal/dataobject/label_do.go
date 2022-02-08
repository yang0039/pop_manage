package dataobject

type Label struct {
	Id         int32  `db:"id"`
	LabelName string `db:"label_name"`
	Operator   string `db:"operator"`
	AddTime    int64  `db:"add_time"`
}
