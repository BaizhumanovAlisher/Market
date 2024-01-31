package repo

import (
	"context"
	"database/sql"
	"platform/config"
	"platform/logging"
)

type SqlRepository struct {
	config.Configuration
	logging.Logger
	Commands SqlCommands
	*sql.DB
	context.Context
}

type SqlCommands struct {
	Init,
	Seed,

	GetProduct,
	GetProducts,

	GetCategories,
	GetCategoryPage,
	GetCategoryPageCount,

	GetPage,
	GetPageCount,

	GetOrder,
	GetOrderLines,
	GetOrders,
	GetOrdersLines,

	SaveOrder,
	SaveOrderLine,
	UpdateOrder,

	SaveProduct,
	UpdateProduct,

	SaveCategory,
	UpdateCategory *sql.Stmt
}
