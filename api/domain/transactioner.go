package domain

type Transactioner interface {
	// コールバック内での操作はトランザクション内で実行される
	Transaction(func(Transaction) error) error
	New() Transaction
}

type Transaction interface{}
