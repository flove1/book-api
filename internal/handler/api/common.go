package api

type ID struct {
	Value int64 `uri:"id,min=0" binding:"required"`
}

type Username struct {
	Value string `uri:"username" binding:"required,min=5,max=50" example:"username"`
}

type Filter struct {
	Page int `form:"page,default=1" default:"1"`

	//Number of books inside of one page. Can range between 1-100
	PageSize int `form:"page_size,default=50" binding:"min=1,max=100" default:"50"`

	//The field that is used for sorting. Add prefix "-" to change direction
	Sort string `form:"sort,default=created_at" default:"created_at"`
}

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DefaultResponseWithBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
