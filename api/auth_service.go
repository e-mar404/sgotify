package api

import (
	"encoding/json"

	"github.com/charmbracelet/log"
)

func init() {
	server.Register(NewAuthService())
}

type LoginArgs struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Code         string
	State        string
}

type RefreshArgs struct {
	RefreshToken string
	ClientID     string
	ClientSecret string
}

type CredentialsReply struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct {
	Client *AuthClient
}

func NewAuthService() *Auth {
	return &Auth{
		Client: NewAuthClient(),
	}
}

func (a *Auth) LoginWithCode(args *LoginArgs, reply *CredentialsReply) error {
	log.Info("Auth.LoginWithCode called")

	q := map[string]string{
		"code":         args.Code,
		"redirect_uri": args.RedirectURI,
		"grant_type":   "authorization_code",
	}
	url := accountBaseURL + "/api/token"
	a.Client.prepArgs = authPrepArgs{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
	}

	loginRes, err := do[CredentialsReply](a.Client, "POST", url, q, nil)
	if err != nil {
		log.Error("unable to do LoginWithCode req", "error", err)
		return err
	}

	*reply = *loginRes

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "CredentialReply", string(jsonReply))
	log.Info("Auth.LoginWithCode sent reply")

	return nil
}

func (a *Auth) RefreshAccessToken(args *RefreshArgs, reply *CredentialsReply) error {
	log.Info("called Auth.RefreshAccessToken")

	q := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": args.RefreshToken,
	}
	url := accountBaseURL + "/api/token"
	a.Client.prepArgs = authPrepArgs{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
	}
	creds, err := do[CredentialsReply](a.Client, "POST", url, q, nil)
	if err != nil {
		return err
	}

	*reply = *creds

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "CredentialReply", string(jsonReply))
	log.Info("Auth.RefreshAccessToken sent reply")

	return nil
}
