package gocash

import (
	"log"
)

type Money struct {
	Amount   int
	Currency *Currency
}

func New(a int, c string) *Money {
	return &Money{
		Amount:   a,
		Currency: new(Currency).Get(c),
	}
}

func (m *Money) SameCurrency(om *Money) bool {
	return m.Currency.Equals(om.Currency)
}

func (m *Money) assertSameCurrency(om *Money) {
	if !m.SameCurrency(om) {
		log.Fatalf("Currencies %s and %s don't match", m.Currency.Code, om.Currency.Code)
	}
}

func (m *Money) compare(om *Money) int {

	m.assertSameCurrency(om)

	switch {
	case m.Amount > om.Amount:
		return 1
	case m.Amount < om.Amount:
		return -1
	}

	return 0
}

func (m *Money) Equals(om *Money) bool {
	return m.compare(om) == 0
}

func (m *Money) GreaterThan(om *Money) bool {
	return m.compare(om) == 1
}

func (m *Money) GreaterThanOrEqual(om *Money) bool {
	return m.compare(om) >= 0
}

func (m *Money) LessThan(om *Money) bool {
	return m.compare(om) == -1
}

func (m *Money) LessThanOrEqual(om *Money) bool {
	return m.compare(om) <= 0
}

func (m *Money) IsZero() bool {
	return m.Amount == 0
}

func (m *Money) IsPositive() bool {
	return m.Amount > 0
}

func (m *Money) IsNegative() bool {
	return m.Amount < 0
}

func (m *Money) Absolute() *Money {
	if m.Amount < 0 {
		m.Amount = -m.Amount
	}

	return m
}

func (m *Money) Negative() *Money {
	if m.Amount > 0 {
		m.Amount = -m.Amount
	}

	return m
}

func (m *Money) Round() {}

//func (m *Money) Add(om *Money) *Money      {}
//func (m *Money) Subtract(om *Money) *Money {}
//func (m *Money) Multiply(om *Money) *Money {}
//func (m *Money) Divide(om *Money) *Money   {}
//func (m *Money) Allocate(r []int) []*Money {}
//func (m *Money) Split(n int) []*Money      {}
