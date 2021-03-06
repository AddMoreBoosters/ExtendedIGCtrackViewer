package main 

type TrackDB struct {
	DatabaseURL string
	DatabaseName string
	CollectionName string
}

type ApiMetadata struct {
	Uptime  string
	Info    string
	Version string
}

type Track struct {
	Hdate       	time.Time 	`json:"H_date"`
	Pilot       	string    	`json:"pilot"`
	Glider      	string    	`json:"glider"`
	GliderID    	string    	`json:"glider_id"`
	TrackLength 	float64   	`json:"track_length"`
	TrackSourceURL	string		`json:"track_src_url"`
}