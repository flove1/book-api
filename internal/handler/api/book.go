package api

type CreateBookRequest struct {
	Title       string   `json:"title" binding:"required,min=5,max=100" example:"The King in Yellow"`
	Author      string   `json:"author" binding:"required,min=5,max=100" example:"Robert W. Chambers"`
	Description string   `json:"description" binding:"required" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
	Tags        []string `json:"tags" binding:"required,min=1,max=5" example:"horror,mystery"`
	Year        int64    `json:"year" binding:"required,min=1" example:"1994"`
}

type UpdateBookRequest struct {
	Title       *string   `json:"title" binding:"omitempty,min=5,max=100" example:"The King in Yellow"`
	Author      *string   `json:"author" binding:"omitempty,min=5,max=100" example:"Robert W. Chambers"`
	Description *string   `json:"description" binding:"omitempty" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
	Tags        *[]string `json:"tags" binding:"omitempty,min=1,max=5" example:"horror,mystery"`
	Year        int64     `json:"year" binding:"omitempty,min=1" example:"1994"`
}

type GetBooksRequest struct {
	Title  *string   `form:"title" binding:"omitempty" example:"The King in Yellow"`
	Author *string   `form:"author" binding:"omitempty" example:"Robert W. Chambers"`
	Tags   *[]string `form:"tags" binding:"omitempty,min=1" example:"horror,mystery"`
	Filter
}
