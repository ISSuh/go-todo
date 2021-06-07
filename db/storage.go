package db

// Storage for user account
type AccountStorage interface {
	InitLocalAccountStorage() error

	CreateAccount(user Account) error
	DeleteAccount(user Account) error
	AccessAccount(user User) Account
}

// Storage for active session
type SessionStorage interface {
	InitLocalSessiontStorage() error

	CreateSession(user User) error
	DeleteSession(session Session) error
	FindSession(session Session) error
}

// Storage for todo contents
type ContentStorage interface {
	InitLocalContentStorage() error

	GetItem(itemId int) (*TodoItem, error)
	GetItemList() (TodoItemList, error)
	AddItem(item TodoItem) (int, error)
	DeleteItem(itemId int) error
}
