package manualwire

import (
	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"github.com/JerryJeager/amor-rendezvous-backend/http"
	"github.com/JerryJeager/amor-rendezvous-backend/service/invitees"
	"github.com/JerryJeager/amor-rendezvous-backend/service/users"
	"github.com/JerryJeager/amor-rendezvous-backend/service/wedding"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}
func GetWeddingRepository() *wedding.WeddingRepo {
	repo := config.GetSession()
	return wedding.NewWeddingRepo(repo)
}

func GetWeddingService(repo wedding.WeddingStore) *wedding.WeddingServ {
	return wedding.NewWeddingService(repo)
}

func GetWeddingController() *http.WeddingController {
	repo := GetWeddingRepository()
	service := GetWeddingService(repo)
	return http.NewWeddingController(service)
}
func GetInviteeRepository() *invitees.InviteeRepo {
	repo := config.GetSession()
	return invitees.NewInviteeRepo(repo)
}

func GetInviteeService(repo invitees.InviteeStore) *invitees.InviteeServ {
	return invitees.NewInviteeService(repo)
}

func GetInviteeController() *http.InviteeController {
	repo := GetInviteeRepository()
	service := GetInviteeService(repo)
	return http.NewInviteeController(service)
}
