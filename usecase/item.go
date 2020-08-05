package usecase

import (
	"errors"

	"github.com/morio-kitahara/go-crud/domain/model"
	"github.com/morio-kitahara/go-crud/domain/repository"
)

// ItemUseCase : Item における UseCase のインターフェース
type ItemUseCase interface {
	FindAll() (items []*model.Item, err error)
	Create(status int, name string) (err error)
	Update(id, status int, name string) (err error)
}

type itemUseCase struct {
	itemRepository repository.ItemRepository
}

// NewItemUseCase : Item データに関する UseCase を生成
func NewItemUseCase(ir repository.ItemRepository) ItemUseCase {
	return &itemUseCase{
		itemRepository: ir,
	}
}

// Search : Item データを全件取得するためのユースケース
func (iu itemUseCase) FindAll() (items []*model.Item, err error) {
	// Persistence（Repository）を呼び出し
	items, err = iu.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Create : Item データを新規登録するためのユースケース
func (iu itemUseCase) Create(status int, name string) (err error) {
	// Item 引数から構造体生成
	item := &model.Item{
		Name:   name,
		Status: status,
	}

	// バリデーションを呼び出し
	if err := validate(item); err != nil {
		return err
	}

	// Persistence（Repository）を呼び出し
	err = iu.itemRepository.Create(item)
	if err != nil {
		return err
	}
	return nil
}

// Update : Item データを更新するためのユースケース
func (iu itemUseCase) Update(id, status int, name string) (err error) {
	// Item 引数から構造体生成
	item := &model.Item{
		Id:     id,
		Name:   name,
		Status: status,
	}

	// バリデーションを呼び出し
	if err := validate(item); err != nil {
		return err
	}

	// Persistence（Repository）を呼び出し
	err = iu.itemRepository.Update(item)
	if err != nil {
		return err
	}
	return nil
}

// Validate : Item データをバリデーション
func validate(item *model.Item) error {
	if len(item.Name) >= 200 {
		return errors.New("タスク名は200文字未満で書いてください。")
	}
	return nil
}
