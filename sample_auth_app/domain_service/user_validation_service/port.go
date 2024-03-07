package user_validation_service

type Port interface {
	CheckDuplicatedInExtSource(dupModels []dupModel) ([]bool, error)
}
