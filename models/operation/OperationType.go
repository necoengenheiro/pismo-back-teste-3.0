package operation

const (
	COMPRA_A_VISTA   = 1
	COMPRA_PARCELADA = 2
	SAQUE            = 3
	PAGAMENTO        = 4
)

type OperationType struct {
	Id           int
	Description0 string
}

func NewOperationType(operationTypeId int) *OperationType {
	return &OperationType{
		Id: operationTypeId,
	}
}

func (op *OperationType) IsValid() bool {
	if op.Id == COMPRA_A_VISTA ||
		op.Id == COMPRA_PARCELADA ||
		op.Id == SAQUE ||
		op.Id == PAGAMENTO {
		return true
	}

	return false
}
