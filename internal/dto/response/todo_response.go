package response

type TodoResponse struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
}