package session

type Repository interface {
	BulkGet([]ID) ([]Session, error)
	BulkSave([]Session) error
	BulkDelete([]ID) error

	Get(ID) (Session, error)
	Save(Session) error
	Delete(ID) error
}
