package datastruct

type MR struct {
	Id     int    `json:"id"`
	Iid    int    `json:"iid"`
	Title  string `json:"title"`
	Link   string `json:"web_url"`
	Author User `json:"author"`
}
