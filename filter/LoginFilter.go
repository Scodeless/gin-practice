package filter

type LoginFilter struct {
	//g *gin.Context
}

func (f *LoginFilter) Login(user map[string]string) bool {
	return true
}