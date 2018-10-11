package repository

import (
	"context"
	"currency/models"
	"currency/pkg/currency"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type mysqlExchangeRateRepo struct {
	DB *sqlx.DB
}

func NewMysqlExchangeRateRepository(db *sqlx.DB) currency.ExchangeRepository {
	return &mysqlExchangeRateRepo{
		DB: db,
	}
}

func (m *mysqlExchangeRateRepo) GetExchangeRate(ctx context.Context, from, to string, date time.Time) (*model.ExchangeRate, error) {

	c := &model.ExchangeRate{}

	err := m.DB.Get(c, "SELECT c.id, c.rate_date, c.to_currency,c.from_currency, c.rate FROM exchange c WHERE c.rate_date = date(?) and c.to_currency = ? and c.from_currency= ?", date.Format("2006-01-02"), to, from)
	if err != nil {
		log.Println("func GetExchangeRate", err)
		return c, err
	}

	eWeekAgo, err := m.GetExchangeRateWeekAverage(ctx, to, from, date)
	c.OneWeekRate = eWeekAgo.OneWeekRate
	c.Variance = eWeekAgo.Variance

	return c, nil
}

func (m *mysqlExchangeRateRepo) GetExchangeRateWeekAverage(ctx context.Context, to, from string, date time.Time) (*model.ExchangeRate, error) {
	const weekBefore = -6
	minDate := date.AddDate(0, 0, weekBefore)
	var total int
	err := m.DB.Get(&total, `SELECT count(*) as total FROM exchange c 
	WHERE (c.rate_date BETWEEN ? AND ?)
	 and c.to_currency = ? and c.from_currency =? `, minDate.Format("2006-01-02"), date.Format("2006-01-02"), to, from)
	if err != nil || total != 7 {
		log.Println(total)
		log.Println("func GetExchangeRateWeekAverage")
		return &model.ExchangeRate{}, errors.New("Lack of data")
	}

	var average float64
	err = m.DB.Get(&average, `SELECT AVG(c.rate) as total FROM exchange c 
	WHERE (c.rate_date BETWEEN ? AND ?)
	 and c.to_currency = ? and c.from_currency = ? `, minDate.Format("2006-01-02"), date.Format("2006-01-02"), to, from)

	if err != nil {
		log.Println("func GetExchangeRateWeekAverage", err)
		return &model.ExchangeRate{}, nil
	}

	var max float64
	var min float64

	err = m.DB.Get(&max, `SELECT MAX(c.rate) as total FROM exchange c 
	WHERE (c.rate_date BETWEEN ? AND ?)
	 and c.to_currency = ? and c.from_currency = ? `, minDate.Format("2006-01-02"), date.Format("2006-01-02"), to, from)

	if err != nil {
		log.Println("func GetExchangeRateWeekAverage", err)
		return &model.ExchangeRate{}, nil
	}

	err = m.DB.Get(&min, `SELECT MIN(c.rate) as total FROM exchange c 
	WHERE (c.rate_date BETWEEN ? AND ?)
	 and c.to_currency = ? and c.from_currency = ? `, minDate.Format("2006-01-02"), date.Format("2006-01-02"), to, from)

	if err != nil {
		log.Println("func GetExchangeRateWeekAverage", err)
		return &model.ExchangeRate{}, nil
	}

	return &model.ExchangeRate{
		Variance:    max - min,
		OneWeekRate: average,
		From:        from,
		To:          to,
		Date:        date,
	}, nil
}

func (m *mysqlExchangeRateRepo) Store(ctx context.Context, c model.ExchangeRate) (int64, error) {
	var lastID int64
	stmtIns, err := m.DB.Prepare(fmt.Sprintf(
		`INSERT INTO exchange (from_currency, to_currency, rate, rate_date) VALUES (?,?,?,?)			
		`))

	if err != nil {
		log.Println("func Store", err)
		return 0, err
	}

	defer stmtIns.Close()
	res, err := stmtIns.Exec(c.From, c.To, c.Rate, c.Date.Format("2006-01-02"))

	if err != nil {
		log.Println("func Store", err)
		return 0, err
	}

	lastID, _ = res.LastInsertId()

	return lastID, nil

}
