package repository

import (
	"backend-ekkn/modules/admin/domain"
)

type AdminRepository interface {
	Create(admin domain.Admin) error
	FindByUsername(username string) (domain.Admin, error)
}
