package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

const APEndpoint = "https://APResolve.spotify.com"

func getAP() {
	resp, err := http.Get(APEndpoint)
	ThrowIfError(err)
	defer ThrowIfError(resp.Body.Close())
	type apList struct {
		ApList []string `json:"ap_list"`
	}
	var data apList
	ThrowIfError(json.NewDecoder(resp.Body).Decode(&data))
	idx := rand.Int31n(int32(len(data.ApList)))
	ThrowData(data.ApList[idx])
}

func GetAP() (string, error) {
	data, err := UnwrapResultFromJob(getAP)
	return data.(string), err
}
