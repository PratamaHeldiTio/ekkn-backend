package service

import (
	"backend-ekkn/modules/admin/domain"
	"backend-ekkn/pkg/shareddomain"
)

type AdminService interface {
	CreateAdmin(request shareddomain.AdminRequest) error
	LoginAdmin(request shareddomain.AdminRequest) (domain.Admin, error)
}
