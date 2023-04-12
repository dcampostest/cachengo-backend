package models

type DbConfiguration struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
}
