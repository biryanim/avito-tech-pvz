package dto

type PVZCreateRequest struct {
	City string `json:"city" required:"true"`
}

type PVZResponse struct {
	ID               string `json:"id"`
	RegistrationDate string `json:"registrationDate"`
	City             string `json:"city"`
}

type PVZListResponse struct {
	PVZ        PVZResponse              `json:"pvz"`
	Receptions []ReceptionsWithProducts `json:"receptions"`
}

type ReceptionsWithProducts struct {
	Reception ReceptionResponse `json:"reception"`
	Products  []ProductResponse `json:"products"`
}

type ReceptionResponse struct {
	ID       string `json:"id"`
	DateTime string `json:"dateTime"`
	PvzID    string `json:"pvzId"`
	Status   string `json:"status"`
}

type CloseReceptionResponse struct {
	Reception ReceptionResponse `json:"reception"`
}

type ProductCreateRequest struct {
	Type  string `json:"type" required:"true"`
	PvzID string `json:"pvzId" required:"true"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	DateTime    string `json:"dateTime"`
	Type        string `json:"type"`
	ReceptionID string `json:"receptionId"`
}

type PaginationRequest struct {
	StartDate string
	EndDate   string
	Page      string
	Limit     string
}

type ReceptionsRequest struct {
	PvzID string `json:"pvzId"`
}
