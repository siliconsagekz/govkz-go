package govkzgo

import (
	config "github.com/siliconsagekz/govkz-go/config"
	// request "github.com/siliconsagekz/govkz-go/pkg/request"
)

type gov struct {
	
}

func Init() *config.Config {
	cnf := config.GetConfig()
	return cnf
}



// EGOV KZ  https://stat.gov.kz/api/juridical?bin=110440015004&lang=kz
//res, _ := request.Request("https://stat.gov.kz/api/juridical?bin=110440015004&lang=kz")
//fmt.Print(res.Status)
// 			https://stat.gov.kz/api/juridical?bin=010123000150&lang=kk