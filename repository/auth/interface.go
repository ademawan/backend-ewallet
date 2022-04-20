package auth

import "backend-ewallet/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
