package session

import (
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var sessionGet *session.Session

var dataAuth *SessionData

type SessionData struct {
	Auth *gen.User
	Role *gen.Role
}

func InitSession(c *fiber.Ctx, user *service.UserService, user_id int) {
	store := session.New()
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	auth, role, _ := user.CurrentUser(uint64(user_id))

	dataAuth = &SessionData{
		Auth: auth,
		Role: role,
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
