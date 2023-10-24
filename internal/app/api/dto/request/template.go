package request

type (
	TemplateRequest struct {
		ShouldReturnSuccess uint `form:"shouldReturnSuccess" binding:"required"`
	}
)
