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

type LoginReply struct {
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

func (a *Auth) LoginWithCode(args *LoginArgs, reply *LoginReply) error {
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

	loginRes, err := do[LoginReply](a.Client, "POST", url, q)
	if err != nil {
		log.Error("unable to do LoginWithCode req", "error", err)
		return err
	}

	*reply = *loginRes

	return nil
}

// func (ac *AuthClient) RefreshAccessToken() (*LoginResponse, error) {
// 	q := map[string]string{
// 		"grant_type":    "refresh_token",
// 		"refresh_token": viper.GetString("refresh_token"),
// 	}
// 	url := viper.GetString("spotify_account_url") + "/api/token"
// 	refreshRes, err := do[LoginResponse](ac, "POST", url, q)
// 	if err != nil {
// 		log.Error("could not refresh access token", "error", err)
// 		return nil, err
// 	}
// 	return refreshRes, nil
// }
