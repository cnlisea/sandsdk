package sandsdk

type Service struct {
	MerId             string // 商户号
	AppPrivateKeyPath string // 应用私钥路径
	SandPublicKeyPath string // 杉德公钥路径
}

func NewService(merId, appPrivateKeyPath, sandPublicKeyPath string) *Service {
	return &Service{
		MerId:             merId,
		AppPrivateKeyPath: appPrivateKeyPath,
		SandPublicKeyPath: sandPublicKeyPath,
	}
}
