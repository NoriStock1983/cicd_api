package dto

/* ログイン画面から入力された内容 */
type LoginUser struct {
	Usercd   string `json:"usercd"`
	Password string `json:"password"`
}

/* 認証されたユーザ情報 */
type AuthUser struct {
	Usercd      string `json:"usercd"`
	User_f_name string `json:"user_f_name"`
	User_l_name string `json:"user_l_name"`
	Password    string `json:"password"`
}
