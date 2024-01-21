package entity

type Usr struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Role     string `json:"role"`
	Reg_dt   string `json:"reg_dt"`
	Updt_dt  string `json:"updt_dt"`
	Sts      string `json:"sts"`
	IsLogin  bool   `json:"isLogin"`
}
