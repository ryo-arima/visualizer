package response

type RelationsResponse struct {
    Code string
    Message string
    Relationss []Relations
}

type Relations struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}