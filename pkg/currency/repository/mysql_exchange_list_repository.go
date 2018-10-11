package repository

import (
	"context"
	model "currency/models"
	"currency/pkg/currency"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type mysqlExchangeListRepo struct {
	DB *sqlx.DB
}

func NewMysqlExchangeListRepository(db *sqlx.DB) currency.ExchangeListRepository {
	return &mysqlExchangeListRepo{
		DB: db,
	}
}

func (m *mysqlExchangeListRepo) GetExchangeLists(ctx context.Context) ([]model.ExchangeList, error) {
	exl := []model.ExchangeList{}

	err := m.DB.Select(&exl, "SELECT el.id, el.from_currency, el.to_currency FROM exchange_list el")
	if err != nil {
		log.Println("func GetExchangeLists", err)
		return exl, err
	}

	return exl, nil
}

func (m *mysqlExchangeListRepo) Store(ctx context.Context, el model.ExchangeList) (int64, error) {
	var lastID int64
	stmtIns, err := m.DB.Prepare(fmt.Sprintf(
		`INSERT INTO exchange_list (from_currency, to_currency) VALUES (?,?)			
		`))

	if err != nil {
		log.Println("func Store", err)
		return 0, err
	}

	defer stmtIns.Close()
	res, err := stmtIns.Exec(el.From, el.To)

	if err != nil {
		log.Println("func Store", err)
		return 0, err
	}

	lastID, _ = res.LastInsertId()

	return lastID, nil
}

func (m *mysqlExchangeListRepo) Delete(ctx context.Context, id int64) (bool, error) {
	stmtIns, err := m.DB.Prepare(fmt.Sprintf(
		`DELETE FROM exchange_list WHERE id = ?			
		`))

	if err != nil {
		log.Println("func Store", err)
		return false, err
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(id)

	if err != nil {
		log.Println("func Delete", err)
		return false, err
	}

	return true, nil
}
