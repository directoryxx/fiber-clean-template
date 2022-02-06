package session

import (
	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var sessionGet *session.Session

var dataAuth *SessionData

type SessionData struct {
	Auth *domain.User
	Role *domain.Role
}

func InitSession(c *fiber.Ctx, user *service.UserService, user_id int) {
	store := session.New()
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	auth := user.CurrentUser(uint64(user_id))

	dataAuth = &SessionData{
		Auth: auth,
		Role: &auth.Role,
	}

	sess.Set("user_id", user_id)

	sessionGet = sess

}

func GetSession() *session.Session {
	return sessionGet
}

func GetAuth() *SessionData {
	return dataAuth
}
