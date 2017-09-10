package main

import (
	//"fmt"
	//"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL: "https://webrepo.nal.ie.u-ryukyu.ac.jp/",
	ClientID: "370442566774-osi0bgsn710brv1v3uc1s7hk24blhdq2.apps.googleusercontent.com",
	ClientSecret: "E46tGSdcop7sU9L8pF30Nz_u",
	Scopes: []string{
		"openid"},
	Endpoint: google.Endpoint,
}

