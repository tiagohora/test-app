package xratedomain

type RepositoryInterface interface {
	GetFromDatabase() error
}
