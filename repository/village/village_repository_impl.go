package village

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/erikrios/reog-apps-apis/utils/logging"
)

type villageRepositoryImpl struct {
	logger logging.Logging
}

func NewVillageRepositoryImpl(logger logging.Logging) *villageRepositoryImpl {
	return &villageRepositoryImpl{logger: logger}
}

func (v *villageRepositoryImpl) FindByID(id string) (village entity.Village, err error) {
	baseUrl := os.Getenv("PONOROGO_ADMINISTRATIVE_AREA_BASE_URL")
	client := &http.Client{}

	request, reqErr := http.NewRequest(http.MethodGet, baseUrl+"/villages/"+id, nil)
	if reqErr != nil {
		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(v.logger, reqErr.Error())

		log.Print(reqErr)
		err = repository.ErrDatabase
		return
	}

	response, reqErr := client.Do(request)
	if reqErr != nil {
		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(v.logger, reqErr.Error())

		log.Print(reqErr)
		err = repository.ErrDatabase
		return
	}
	defer func(body io.ReadCloser) {
		if reqErr := body.Close(); reqErr != nil {
			log.Println(reqErr)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		var villageResponse serviceResponse[villageResponse]
		if decodeErr := json.NewDecoder(response.Body).Decode(&villageResponse); decodeErr != nil {
			go func(logger logging.Logging, message string) {
				logger.Error(message)
			}(v.logger, decodeErr.Error())

			log.Println(decodeErr)
			err = repository.ErrDatabase
		} else {
			villageResponse := villageResponse.Data
			village = toVillage(villageResponse)
		}
	} else if response.StatusCode == http.StatusNotFound {
		err = repository.ErrRecordNotFound
	} else {
		err = repository.ErrDatabase
	}
	return
}

type serviceResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type villageResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	District districtResponse `json:"district"`
}

type districtResponse struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Regency regencyResponse `json:"regency"`
}

type regencyResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Province provinceResponse `json:"province"`
}

type provinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toVillage(response villageResponse) (village entity.Village) {
	village.ID = response.ID
	village.Name = response.Name
	village.District.ID = response.District.ID
	village.District.Name = response.District.Name
	village.District.Regency.ID = response.District.Regency.ID
	village.District.Regency.Name = response.District.Regency.Name
	village.District.Regency.Province.ID = response.District.Regency.Province.ID
	village.District.Regency.Province.Name = response.District.Regency.Province.Name
	return
}
