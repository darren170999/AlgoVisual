package requests

type EditAQuestionRequest struct {
	// Id        primitive.ObjectID `json:"id,omitempty"`
	QnId        string `json:"qnid,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Status      string `json:"status,omitempty" validate:"required"`
	Tags        string `json:"tags,omitempty" validate:"required"`
}
