package session

type Repository interface {
	BulkGet([]ID) ([]Session, error)
	BulkSave([]Session) error
	BulkDelete([]ID) error
}
