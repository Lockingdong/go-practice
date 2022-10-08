package isp_demo

// interface for user email srv
type IUserEmailRepo interface {
	GetUserEmailByID(userID string) string
}

// interface for user profile srv
type IUserProfileRepo interface {
	GetUserProfileByID(userID string) any
}

// 以下為 UserRepo 實作
type UserRepo struct {
}

func (*UserRepo) GetUserEmailByID(userID string) string {
	panic("implement me")
}

func (*UserRepo) GetUserProfileByID(userID string) any {
	panic("implement me")
}

var _ IUserEmailRepo = (*UserRepo)(nil)
var _ IUserProfileRepo = (*UserRepo)(nil)
