package user_validation_service

//go:generate moq -out ./$GOFILE.moq.go . Port
type Port interface {
	CheckDuplicatedInExtSource(dupModels []dupModel) ([]bool, error)
}
