package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewBuyerMap(db map[int64]domain.Buyer, loader *loader.JSONLoader[domain.BuyerDoc, int64]) *BuyerMap {
	// default db
	defaultDb := make(map[int64]domain.Buyer)
	if db != nil {
		defaultDb = db
	}
	return &BuyerMap{
		db:     defaultDb,
		loader: loader,
	}
}

type BuyerMap struct {
	db     map[int64]domain.Buyer
	loader *loader.JSONLoader[domain.BuyerDoc, int64]
}

func (r *BuyerMap) FindAll() (buyers []domain.Buyer, err error) {
	buyers = utils.MapToSlice(r.db)
	return buyers, err
}

func (r *BuyerMap) Create(buyer domain.Buyer) (domain.Buyer, error) {
	id := time.Now().UnixNano()
	if _, ok := r.db[id]; ok {
		return domain.Buyer{}, utils.ErrInternalServer
	}
	for _, v := range r.db {
		if v.CardNumberId == buyer.CardNumberId {
			return domain.Buyer{}, utils.ErrBuyerCardNumberAlreadyExists
		}
	}
	buyer.Id = id
	r.db[buyer.Id] = buyer
	// Save to JSON
	if err := r.loader.SaveToJSON(r.getDBInBuyerDoc()); err != nil {
		return domain.Buyer{}, err
	}
	return buyer, nil
}

func (r *BuyerMap) getDBInBuyerDoc() map[int64]domain.BuyerDoc {
	dbDoc := make(map[int64]domain.BuyerDoc)
	for key, value := range r.db {
		dbDoc[key] = value.ParseToBuyerDoc()
	}
	return dbDoc
}
func (r *BuyerMap) FindById(id int64) (buyer domain.Buyer, err error) {
	buyer, ok := r.db[id]
	if !ok {
		return buyer, utils.ErrBuyerNotFound
	}
	return buyer, err
}

func (r *BuyerMap) Update(id int64, buyer domain.Buyer) (domain.Buyer, error) {
	if _, ok := r.db[id]; !ok {
		return domain.Buyer{}, utils.ErrBuyerNotFound
	}
	r.db[id] = buyer
	// Save to JSON
	if err := r.loader.SaveToJSON(r.getDBInBuyerDoc()); err != nil {
		return domain.Buyer{}, err
	}
	return buyer, nil
}

func (r *BuyerMap) ExistByCardId(id int) bool {
	for _, v := range r.db {
		if v.CardNumberId == id {
			return true
		}
	}
	return false
}

func (r *BuyerMap) Delete(id int64) (err error) {
	_, ok := r.db[id]
	if !ok {
		return utils.ErrBuyerNotFound
	}
	delete(r.db, id)
	if err := r.loader.SaveToJSON(r.getDBInBuyerDoc()); err != nil {
		return err
	}
	return nil
}
