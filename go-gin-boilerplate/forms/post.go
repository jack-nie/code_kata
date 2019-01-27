package forms

type PostContent struct {
	Content string `json:"content" binding:"required"`
	Title   string `json:"title" binding:"required"`
}
