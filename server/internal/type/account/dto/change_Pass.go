package dto_account

type ChangePass struct {
	OldPass     string `p:"oldPass" v:"required#请输入旧密码"     dc:"旧密码"`
	NewPass     string `p:"newPass" v:"required|passport|different:oldPass#请输入旧密码|字母开头，只能包含字母、数字和下划线，长度在6~18之间|新密码不能与原密码相同" dc:"新密码"`
	ConfirmPass string `p:"confirmPass" v:"required|passport|same:newPass#请输入旧密码|字母开头，只能包含字母、数字和下划线，长度在6~18之间|两次输入的密码不一致" dc:"旧密码"`
}
