package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewBuyerSQL(db *sql.DB) *BuyerSQL {
	return &BuyerSQL{
		db: db,
	}
}

type BuyerSQL struct {
	db *sql.DB
}

func (r *BuyerSQL) FindAll() (buyers []domain.Buyer, err error) {
	rows, err := r.db.Query(queries.FindAllBuyers)
	if err != nil {

		return []domain.Buyer{}, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	for rows.Next() {
		var buyer domain.Buyer
		err := rows.Scan(&buyer.Id, &buyer.FirstName, &buyer.LastName, &buyer.CardNumberId)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		buyers = append(buyers, buyer)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.ErrInternalServer
	}

	return
}

func (r *BuyerSQL) Create(buyer domain.Buyer) (domain.Buyer, error) {
	result, err := r.db.Exec(queries.CreateBuyer,
		buyer.FirstName, buyer.LastName, buyer.CardNumberId)
	if err != nil {
		err = utils.ValidateErrorTypeSQL(err)
		if errors.Is(err, utils.ErrDBDuplicateEntry) {
			return domain.Buyer{}, utils.ErrBuyerCardNumberAlreadyExists
		}
		return domain.Buyer{}, utils.ErrInternalServer

	}
	if result == nil {
		return domain.Buyer{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Buyer{}, utils.ErrInternalServer
	}

	buyer.Id = id
	return buyer, nil
}

func (r *BuyerSQL) FindById(id int64) (buyer domain.Buyer, err error) {
	row := r.db.QueryRow(queries.FindBuyerByID, id)

	if err := row.Scan(&buyer.Id, &buyer.FirstName, &buyer.LastName, &buyer.CardNumberId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Buyer{}, utils.ErrBuyerNotFound
		}
		return domain.Buyer{}, utils.ErrInternalServer
	}

	return buyer, nil
}

func (r *BuyerSQL) Update(id int64, buyer domain.Buyer) (domain.Buyer, error) {
	result, err := r.db.Exec(
		queries.UpdateBuyer,
		buyer.FirstName, buyer.LastName, buyer.CardNumberId, id,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBDuplicateEntry) {
			return domain.Buyer{}, utils.ErrBuyerCardNumberAlreadyExists
		}
		return domain.Buyer{}, utils.ErrInternalServer
	}

	if result == nil {
		return domain.Buyer{}, utils.ErrInternalServer
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Buyer{}, utils.ErrInternalServer
	}

	return buyer, nil
}

func (r *BuyerSQL) Delete(id int64) (err error) {
	result, err := r.db.Exec(queries.DeleteBuyer, id)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBCannotDeleteOrUpdateParent) {
			return utils.ErrNoDeleteByForeignKey
		}
		return utils.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrInternalServer
	}

	if rowsAffected == 0 {
		return utils.ErrBuyerNotFound
	}

	return nil
}

func (r *BuyerSQL) ReportPurchaseOrdersByBuyerID(buyerID int64) ([]domain.BuyerPurchaseOrdersReport, error) {
	row := r.db.QueryRow(queries.ReportPurchaseOrdersByBuyerID, buyerID)
	var report domain.BuyerPurchaseOrdersReport
	err := row.Scan(&report.Id, &report.CardNumberId, &report.FirstName, &report.LastName, &report.PurchaseOrders)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrBuyerNotFound
		}
		return nil, utils.ErrInternalServer
	}
	return []domain.BuyerPurchaseOrdersReport{report}, nil
}

func (r *BuyerSQL) ReportPurchaseOrdersAll() ([]domain.BuyerPurchaseOrdersReport, error) {
	rows, err := r.db.Query(queries.ReportPurchaseOrdersByBuyer)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	reports := make([]domain.BuyerPurchaseOrdersReport, 0)
	for rows.Next() {
		var report domain.BuyerPurchaseOrdersReport
		err := rows.Scan(&report.Id, &report.CardNumberId, &report.FirstName, &report.LastName, &report.PurchaseOrders)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		reports = append(reports, report)
	}
	return reports, nil
}
