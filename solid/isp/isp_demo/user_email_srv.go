package isp_demo

type UserEmailSrv struct {
	repo IUserEmailRepo
}

func NewUserEmailSrv(repo IUserEmailRepo) *UserEmailSrv {
	return &UserEmailSrv{
		repo: repo,
	}
}
