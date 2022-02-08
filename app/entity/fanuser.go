package entity

type Fanuser struct {
	Id        int    `json:"-"`
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
