package v1

import "github.com/muratovdias/my-project/internal/service"

type serviceI interface {
	service.Entry
	service.Account
}
