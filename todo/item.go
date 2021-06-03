package todo

type Work struct {
	Content string `json: "content"`
	Done    bool   `json: "done`
}

type Item struct {
	Time  string `json: "time`
	Title string `json: "title"`
	Works []Work `json: "works, omitempty"`
}

type TodoItem struct {
	Id   int  `json: "id, omitempty"`
	Item Item `json: "item"`
}

type TodoItemList struct {
	List []TodoItem `json: "list"`
}
