package services

import (
	"io"
	"fmt"
	"encoding/json"

	config "github.com/siliconsagekz/govkz-go/config"
	"github.com/siliconsagekz/govkz-go/internal/models"
	request "github.com/siliconsagekz/govkz-go/pkg/request"
	// parser "github.com/siliconsagekz/govkz-go/internal/parser"
)

type stat struct {
	Conf *config.Config
}

type Stat interface {
	GetInfoByIinOrBin(bin, ln string) (*models.Company, error)
}

func (s *stat) GetInfoByIinOrBin(bin, ln string) (*models.Company, error) {
	control := models.Control{}
	
	uri := fmt.Sprintf("%s/api/juridical?bin=%s&lang=%s",s.Conf.Services.External.Stat_gov, bin, ln)
	
	response, err := request.Request(uri); if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body); if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &control); err != nil {
		return nil, err
	}

	return &control.Obj, nil 
}


func NewStat(conf *config.Config) Stat {
	return &stat{
		Conf: conf,
	}
}
