package api

import "github.com/charmbracelet/log"

func init() {
	server.Register(NewAuthService())
}

type LoginArgs struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	BaseURL      string
	Code         string
	State        string
}

type RefreshArgs struct {
	RefreshToken string
	BaseURL string
}

type CredentialsReply struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService struct {
	Client *AuthClient
}

func NewAuthService() *AuthService {
	return &AuthService{
		Client: NewAuthClient(),
	}
}

func (a *AuthService) LoginWithCode(args *LoginArgs, reply *CredentialsReply) error {
	q := map[string]string{
		"code":         args.Code,
		"redirect_uri": args.RedirectURI,
		"grant_type":   "authorization_code",
	}
	url := args.BaseURL + "/api/token"
	a.Client.prepArgs = authPrepArgs{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
	}

	loginRes, err := do[CredentialsReply](a.Client, "POST", url, q)
	if err != nil {
		log.Error("unable to do LoginWithCode req", "error", err)
		return err
	}

	*reply = *loginRes

	return nil
}

func (a *AuthService) RefreshAccessToken(args *RefreshArgs, reply *CredentialsReply) error {
	q := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": args.RefreshToken, 
	}
	url := args.BaseURL + "/api/token"
	creds, err := do[CredentialsReply](a.Client, "POST", url, q)
	if err != nil {
		return err
	}

	*reply = *creds

	return  nil
}
