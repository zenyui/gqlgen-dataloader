package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	// modified this struct to hold a UserID instead of a User so
	// that we can lazy-load the nested Users with the dataloader
	UserID string `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
	// optional UserID supplied during creation to make some
	// testing easier
	UserID string `json:"userId"`
}
