package entity

import "shop/internal/sell-service/domain"

type Customer struct {
	ID              uint32
	Name            string
	Age             int
	WorldPerception domain.WorldPerceptionType
}
