package persistence

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/morio-kitahara/go-crud/domain/model"
	"github.com/morio-kitahara/go-crud/domain/repository"
)

type itemPersistence struct {
	Conn *gorm.DB
}

// NewItemPersistence : Item データに関する Persistence を生成
func NewItemPersistence(conn *gorm.DB) repository.ItemRepository {
	return &itemPersistence{Conn: conn}
}

// Search : DB から Item データの全件取得（ItemRepository インターフェースの Search() を実装したもの）
func (ip *itemPersistence) FindAll() (items []*model.Item, err error) {
	// DB接続
	db := ip.Conn

	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

// Create : DB から Item データの新規登録（ItemRepository インターフェースの Create() を実装したもの）
func (ip *itemPersistence) Create(item *model.Item) (err error) {
	// DB接続
	db := ip.Conn
	now := time.Now()

	// 新規登録用Item生成
	createItem := &model.Item{
		Name:      item.Name,
		Status:    item.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 新規登録
	if err := db.Create(&createItem).Error; err != nil {
		return err
	}

	return nil
}

// Update : DB から Item データの更新（ItemRepository インターフェースの Update() を実装したもの）
func (ip *itemPersistence) Update(item *model.Item) (err error) {
	// DB接続
	db := ip.Conn
	now := time.Now()

	// 更新対象のItem取得
	updateItem := &model.Item{}
	if err := db.First(&updateItem, item.Id).Error; err != nil {
		return err
	}

	updateItem.Name = item.Name
	updateItem.Status = item.Status
	updateItem.UpdatedAt = now

	// 更新
	if err := db.Save(&updateItem).Error; err != nil {
		return err
	}

	return nil
}
