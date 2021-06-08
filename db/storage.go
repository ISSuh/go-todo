package db

// Storage for user account
type AccountStorage interface {
	InitAccountStorage() error

	CreateAccount(user Account) error
	DeleteAccount(user Account) error
	AccessAccount(user User) Account
}

// Storage for active session
type SessionStorage interface {
	InitSessionStorage() error

	AddSession(session Session) error
	DeleteSession(user User) error
	GetSession(user User) (*Session, error)
}

// Storage for todo contents
type ContentStorage interface {
	InitContentStorage() error

	GetItem(itemId int) (*TodoItem, error)
	GetItemList() (TodoItemList, error)
	AddItem(item TodoItem) (int, error)
	DeleteItem(itemId int) error
}
