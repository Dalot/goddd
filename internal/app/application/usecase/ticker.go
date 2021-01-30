package usecase

import (
	"github.com/Dalot/goddd/internal/app/application/service"
	"github.com/Dalot/goddd/internal/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}
