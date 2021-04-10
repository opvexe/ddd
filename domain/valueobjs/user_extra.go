package valueobjs

// 用于描述额外信息
type UserExtra struct {
	UserPhone string
	UserQQ    string
	UserCity  string
}

func NewUserExtra(attrs ...UserExtraAttrFunc) *UserExtra {
	userExtra := &UserExtra{}
	UserExtraAttrFuncs(attrs).apply(userExtra)
	return userExtra
}

func (this *UserExtra) Equal(other *UserExtra) bool {
	return this.UserPhone == other.UserPhone && this.UserQQ == other.UserQQ && this.UserCity == other.UserCity
}
