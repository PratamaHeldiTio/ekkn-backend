package repository

import (
	"backend-ekkn/modules/admin/domain"
)

type AdminRepository interface {
	Save(admin domain.Admin) (domain.Admin, error)
}
