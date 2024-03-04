package user

type Repository interface {
	BulkGet([]ID) ([]User, error)
	BulkSave([]User) error
	BulkDelete([]ID) error
}
