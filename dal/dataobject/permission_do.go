package dataobject

type Menu struct {
	Id      int32  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Title   string `db:"title" json:"title"`
	AddTime int64  `db:"add_time" json:"add_time"`
}

type Permissions struct {
	Id        int32  `db:"id" json:"id"`
	MenuId    int32  `db:"menu_id" json:"menu_id"`
	FuncName  string `db:"func_name" json:"func_name"`
	FuncTitle string `db:"func_title" json:"func_title"`
	Name      string `db:"name" json:"name"`
	Title     string `db:"title" json:"title"`
	AddTime   int64  `db:"add_time" json:"add_time"`
}

type MenuFunc struct {
	MenuId    int32  `db:"menu_id" json:"menu_id"`
	FuncName  string `db:"func_name" json:"func_name"`
	FuncTitle string `db:"func_title" json:"func_title"`
}

type PermissionsUrl struct {
	Id            int32  `db:"id" json:"id"`
	PermissionsId int32  `db:"permissions_id" json:"permissions_id"`
	Url           string `db:"url" json:"url"`
	Method        string `db:"method" json:"method"`
	MethodName    string `db:"method_name" json:"method_name"`
	IsEffect      bool   `db:"is_effect" json:"is_effect"`
	AddTime       int64  `db:"add_time" json:"add_time"`
}
