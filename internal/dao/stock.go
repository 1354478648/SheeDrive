// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"SheeDrive/internal/dao/internal"
)

// internalStockDao is internal type for wrapping internal DAO implements.
type internalStockDao = *internal.StockDao

// stockDao is the data access object for table stock.
// You can define custom methods on it to extend its functionality as you wish.
type stockDao struct {
	internalStockDao
}

var (
	// Stock is globally public accessible object for table stock operations.
	Stock = stockDao{
		internal.NewStockDao(),
	}
)

// Fill with you ideas below.
