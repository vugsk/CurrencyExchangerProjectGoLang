package models

type InterfaceDataBase interface {
	Close() error
	Connect(driverName, userName, dataBaseName string) error

	Insert(tableName string, obj any) error
	Read(tableName string, idProfile string) (error, any)
	ReadAll(tableName string) ([]any, error)
	Update(tableName string, obj any) error
	Delete(tableName string, idProfile string) error

	Count(tableName string) (uint, error)
	IsEmpty(tableName string) bool
	Ping() error
	Migrate() error
	IsConnected() bool
}
