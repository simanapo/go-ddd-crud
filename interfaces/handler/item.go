package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/morio-kitahara/go-crud/usecase"
)

// ItemHandler : Item における Handler のインターフェース
type ItemHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	Update(http.ResponseWriter, *http.Request, httprouter.Params)
}

type itemHandler struct {
	itemUseCase usecase.ItemUseCase
}

// NewItemUseCase : Item データに関する Handler を生成
func NewItemHandler(iu usecase.ItemUseCase) ItemHandler {
	return &itemHandler{
		itemUseCase: iu,
	}
}

// ItemIndex : GET /items -> Item データの全件取得結果を返す
func (ih itemHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// ユースケースの呼び出し
	items, err := ih.itemUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// クライアントにレスポンスを返却
	if err = json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// ItemCreate : POST /items -> Item データの新規登録結果を返す
func (ih itemHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// パラメータ
	name := r.FormValue("name")
	status, _ := strconv.Atoi(r.FormValue("status"))

	// ユースケースの呼び出し
	err := ih.itemUseCase.Create(status, name)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// ItemUpdate : PUT /items -> Item データの更新結果を返す
func (ih itemHandler) Update(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// パラメータ
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	status, _ := strconv.Atoi(r.FormValue("status"))

	// ユースケースの呼び出し
	err := ih.itemUseCase.Update(id, status, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
