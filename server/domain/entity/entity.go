package entity

/*basic usr login information*/
type Usr struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Class    string `json:"class"` //can be seen in com_code type:I3
	Role     string `json:"role"`  //can be seen in com_code type:I2
	Reg_dt   string `json:"reg_dt"`
	Updt_dt  string `json:"updt_dt"`
	Sts      string `json:"sts"`     //can be seen in com_code type:I1
	IsLogin  bool   `json:"isLogin"` //true: can keep login, false: cannot
}

type ComCode struct {
	Code     string `json:"code"`
	Val      string `json:"val"`
	Remark   string `json:"remark"`
	CodeType string `json:"code_type"`
}

type ScoreInf struct {
	Username string `json:"username"`
	Total    int    `json:"total"`
	Reg_dt   string `json:"reg_dt"`
	Updt_dt  string `json:"updt_dt"`
	His_tot  string `json:"his_tot"`
}

type ScoreHis struct {
	Username string `json:"username"`
	Trx_tot  int    `json:"trx_tot"`
	Total    int    `json:"total"`
	Reg_dt   string `json:"reg_dt"`
	His_no   int    `json:"his_no"`
	Remark   string `json:"remark"`
	ApprvUsr string `json:"apprv_usr"`
}
