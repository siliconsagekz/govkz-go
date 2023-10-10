package models

import (
	"time"
)

type Control struct {
	Success 	bool `json:"success"`
	Obj 		Company `json:"obj"`
}

type Company struct {
	Id 				int `json:"id"`
	Bin 			string `json:"bin"`
	Name 			string `json:"name"`
	RegisterDate 	time.Time `json:"registerDate"`
	OkedCode 		string `json:"okedCode"`
	OkedName 		string `json:"okedName"`
	SecondOkeds 	string `json:"secondOkeds"`
	KrpCode 		string `json:"krpCode"`
	KrpName 		string `json:"krpName"`
	KrpBfCode 		string `json:"krpBfCode"`
	KrpBfName 		string `json:"krpBfName"`
	KseCode 		string `json:"kseCode"`
	KseName 		string `json:"kseName"`
	KfsCode			string `json:"kfsCode"`
	KfsName			string `json:"kfsName"`
	KatoCode		string `json:"katoCode"`
	KatoId			int `json:"katoId"`
	KatoAddress		string `json:"katoAddress"`
	Fio				string `json:"fio"`
	Ip				bool `json:"ip"`
}