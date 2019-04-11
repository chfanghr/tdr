package core

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/chfanghr/tdr/spotify/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type OAuth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Error        string
}

var PrintOAuthMessage = func(msg string) {
	fmt.Println(msg)
}

func GetOauthAccessToken(code string, redirectUri string, clientId string, clientSecret string) (*OAuth, error) {
	if data, err := UnwrapResultFromJob(func() {
		val := url.Values{}
		val.Set("grant_type", "authorization_code")
		val.Set("code", code)
		val.Set("redirect_uri", redirectUri)
		val.Set("client_id", clientId)
		val.Set("client_secret", clientSecret)

		resp, err := http.PostForm("https://accounts.spotify.com/api/token", val)
		if err != nil {
			// Retry since there is an nginx bug that causes http2 streams to get
			// an initial REFUSED_STREAM response
			// https://github.com/curl/curl/issues/804
			resp, err = http.PostForm("https://accounts.spotify.com/api/token", val)
			ThrowIfError(err)
		}
		defer ThrowIfError(resp.Body.Close())
		auth := OAuth{}
		body, err := ioutil.ReadAll(resp.Body)
		ThrowIfError(err)
		ThrowIfError(json.Unmarshal(body, &auth))
		if auth.Error != "" {
			WrapAndThrowError("error getting token", errors.New(auth.Error))
		}
		ThrowData(&auth)
	}); err != nil {
		return nil, err
	} else {
		return data.(*OAuth), nil
	}
}

func getOAuthToken(clientId string, clientSecret string) OAuth {
	ch := make(chan OAuth)

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	ThrowIfError(err)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		auth, err := GetOauthAccessToken(params.Get("code"), "http://localhost:8888/callback", clientId, clientSecret)
		if err != nil {
			IgnoreError(fmt.Fprintf(w, "failed to get token %q", err))
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		IgnoreError(fmt.Fprintf(w, "got token, loggin in"))
		ch <- *auth
	})

	defer func() { CrashProgramIfError("close listener", listener.Close()) }()
	go func() { _ = http.Serve(listener, http.DefaultServeMux) }()

	PrintOAuthMessage("go to this url")
	urlPath := "https://accounts.spotify.com/authorize?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&redirect_uri=http://localhost:8888/callback" +
		"&scope=streaming"
	PrintOAuthMessage(urlPath)

	return <-ch
}
