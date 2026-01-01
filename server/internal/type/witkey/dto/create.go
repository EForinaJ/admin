package dto_witkey

type Create struct {
	Name    string   `p:"name" v:"required|max-length:8#请输入威客昵称|长度最长为8" dc:"昵称"`
	UserId  int64    `p:"userId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"用户id"`
	TitleId int64    `p:"titleId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"头衔id"`
	GameId  int64    `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属分类id"`
	Album   []string `p:"album"  dc:"相册"`
	Rate    int      `p:"rate" v:"between:1,5#评分数值已超出" dc:"评分"`
}
