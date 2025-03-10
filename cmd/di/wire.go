//go:build wireinject
// +build wireinject

package di

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	"github.com/changchanghwang/wdwb_back/internal/server"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(db.Init, server.New, stocks.StockSet, secClient.New)
	return &server.Server{}, nil
}
