package models


type SummarizeRequest struct {
	Text string `json:"text"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
}