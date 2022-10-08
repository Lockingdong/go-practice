package isp_demo

type UserProfileSrv struct {
	repo IUserProfileRepo
}

func NewUserProfileSrv(repo IUserProfileRepo) *UserProfileSrv {
	return &UserProfileSrv{
		repo: repo,
	}
}
