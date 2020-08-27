package Models



type Request struct {
	ID            int		`json:"id"`
	StatusCode    int       `json:"status_code"     gorm:"type:int(1)"`
	LatencyTime   string    `json:"latency_time"    gorm:"type:varchar(30)"`
	ClientIP      string    `json:"client_ip"       gorm:"type:varchar(20)"`
	ReqMethod     string    `json:"req_method"      gorm:"type:varchar(10)"`
	ReqUri        string    `json:"req_uri"         gorm:"type:varchar(50)"`
	Body          string    `json:"body"            gorm:"type:text"`
}