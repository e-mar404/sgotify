package api

import "github.com/charmbracelet/log"

func init() {
	server.Register(NewUserService())
}

type ProfileArgs struct {
}

type User struct {
	Client *UserClient
}

func NewUserService() *User {
	return &User{
		Client: NewUserClient(),
	}
}

func (u *User) Profile(args *int, reply *string) error {
	log.Info("calling user.profile")
	*reply = "profile"

	return nil
}
