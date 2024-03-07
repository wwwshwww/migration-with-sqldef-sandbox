package user

type Repository interface {
	BulkGet([]ID) ([]User, error)
	BulkSave([]User) error
	BulkDelete([]ID) error

	Get(ID) (User, error)
	Save(User) error
}
