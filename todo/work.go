package todo

type Work struct {
	Content string `json: "content"`
	Done    bool   `json: "done`
}

type Item struct {
	Time     string `json: "time`
	Title    string `json: "title"`
	WorkList []Work `json: "work`
}

type StoredItem struct {
	Id   int32 `json: "id"`
	Item Item  `json: "item"`
}

type StoredItemList struct {
	ItemList []StoredItem `json: "item"`
}

type Method interface {
	AddItem(item Item) (int32, error)
	DeleteItem(itemId int32) error
	GetItem(itemId int32) (Item, error)
	GetItemList() (StoredItemList, error)
}
