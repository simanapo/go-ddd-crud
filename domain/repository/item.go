package repository

import (
	"github.com/morio-kitahara/go-crud/domain/model"
)

// ItemRepository : Item における Repository のインターフェース
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type ItemRepository interface {
	FindAll() (items []*model.Item, err error)
	Create(item *model.Item) (err error)
	Update(item *model.Item) (err error)
}
