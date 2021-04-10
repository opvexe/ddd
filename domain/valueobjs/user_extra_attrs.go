package valueobjs

type UserExtraAttrFunc func(model *UserExtra)
type UserExtraAttrFuncs []UserExtraAttrFunc

//apply 方法
func (this UserExtraAttrFuncs) apply(u *UserExtra) {
	for _, f := range this {
		f(u)
	}
}

func WithUserPhone(phone string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserPhone = phone
	}
}

func WithUserQQ(qq string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserQQ = qq
	}
}

func WithUserCity(city string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserCity = city
	}
}
