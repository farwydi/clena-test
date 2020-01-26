package gateway

type ContentTranslator interface {
	ToServiceGateway() ServiceGateway
	ToCategoryGateway() CategoryGateway
	ToContentGateway() ContentGateway
	BeginTx() (ContentTranslator, error)
}
