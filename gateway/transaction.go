package gateway

type Transaction interface {
	Commit() error
	Rollback() error
}
