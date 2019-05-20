package githubauth

import (
	"bytes"
	"github.com/parnurzeal/gorequest"
	"fmt"
)

//github  url 
var githubBase = "https://github.com/login/oauth/authorize"
var githubOrg = "https://api.github.com/orgs/"
var githubToken = "https://github.com/login/oauth/access_token/"
var githubUser = "https://api.github.com/user"

//ContentProvider struct 
type ContentProvider struct {
	Code string
	ClientID string
	SecretKey string
	OrgName	string
	Token string
	Name string
}


//GetGithubAuth func
func GetGithubAuth(c *ContentProvider) string {
	var buffer bytes.Buffer
	buffer.WriteString(githubBase)
	buffer.WriteString("?client_id=")
	buffer.WriteString(c.ClientID)
	fmt.Println("=======================================")
	fmt.Println(c.ClientID)
	fmt.Println("=======================================")
	buffer.WriteString("&scope=user%20admin:org%20repo&allow_singup=false")
	return buffer.String()
}

//GetOrg func
func GetOrg(c *ContentProvider) int {
	var tokenbuffer bytes.Buffer
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(c.Token)

	var urlbuffer bytes.Buffer
	urlbuffer.WriteString(githubOrg)
	urlbuffer.WriteString(c.OrgName)
	urlbuffer.WriteString("/memberships/")
	urlbuffer.WriteString(c.Name)
	fmt.Println("=======================================")
	fmt.Println(urlbuffer.String())
	fmt.Println(tokenbuffer.String())
	fmt.Println("=======================================")
	resp,_,_ :=gorequest.New().Get(urlbuffer.String()).Set("Authorization", tokenbuffer.String()).End()
	
	return resp.StatusCode
	
}

//GetToken func
func GetToken(c *ContentProvider) string {
	type jsonAccessToken struct {
		AccessToken string `json:"access_token"`
		TokenType string
		Scope string
	}
	type Date struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}
	var jsonToken jsonAccessToken
	dataSend := Date{Code: c.Code, ClientID: c.ClientID, ClientSecret: c.SecretKey}
	gorequest.New().Post(githubToken).Set("Accept","application/json").Send(dataSend).EndStruct(&jsonToken)
	fmt.Println(jsonToken)
	return jsonToken.AccessToken
}


//GetUsername func
func GetUsername(token string) string {
	// 定義接收回傳欄位
	type jsonusername struct {
		Login string
	}
	var tokenbuffer bytes.Buffer
	var jsonUser jsonusername
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(token)
	gorequest.New().Get(githubUser).Set("Authorization", tokenbuffer.String()).EndStruct(&jsonUser)
	return jsonUser.Login
}