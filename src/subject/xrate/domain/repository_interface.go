package domain

type RepositoryInterface interface {
	GetFromDatabase() error
}
