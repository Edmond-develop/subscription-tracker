package database

type Subscription struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserName    string `json:"user_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
