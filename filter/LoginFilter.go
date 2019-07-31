package filter

type LoginFilter struct {
	//g *gin.Context
}

func (f *LoginFilter) Login() bool {
	return true
}