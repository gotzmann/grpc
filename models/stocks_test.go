// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testStocks(t *testing.T) {
	t.Parallel()

	query := Stocks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Stocks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StockExists(ctx, tx, o.StoreID, o.ProductID)
	if err != nil {
		t.Errorf("Unable to check if Stock exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockExists to return true, but got false.")
	}
}

func testStocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stockFound, err := FindStock(ctx, tx, o.StoreID, o.ProductID)
	if err != nil {
		t.Error(err)
	}

	if stockFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Stocks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Stocks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testStocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stockColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockToOneProductUsingProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Stock
	var foreign Product

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProductID = foreign.ProductID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Product().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ProductID != foreign.ProductID {
		t.Errorf("want: %v, got %v", foreign.ProductID, check.ProductID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadProduct(ctx, tx, false, (*[]*Stock)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Product = nil
	if err = local.L.LoadProduct(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockToOneStoreUsingStore(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Stock
	var foreign Store

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StoreID = foreign.StoreID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Store().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.StoreID != foreign.StoreID {
		t.Errorf("want: %v, got %v", foreign.StoreID, check.StoreID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadStore(ctx, tx, false, (*[]*Stock)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Store = nil
	if err = local.L.LoadStore(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockToOneSetOpProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Stock
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Product{&b, &c} {
		err = a.SetProduct(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Product != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProductID != x.ProductID {
			t.Error("foreign key was wrong value", a.ProductID)
		}

		if exists, err := StockExists(ctx, tx, a.StoreID, a.ProductID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testStockToOneSetOpStoreUsingStore(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Stock
	var b, c Store

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Store{&b, &c} {
		err = a.SetStore(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Store != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StoreID != x.StoreID {
			t.Error("foreign key was wrong value", a.StoreID)
		}

		if exists, err := StockExists(ctx, tx, a.StoreID, a.ProductID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testStocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockDBTypes = map[string]string{`StoreID`: `integer`, `ProductID`: `integer`, `Quantity`: `integer`}
	_            = bytes.MinRead
)

func testStocksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockDBTypes, true, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockDBTypes, true, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockAllColumns, stockPrimaryKeyColumns) {
		fields = stockAllColumns
	} else {
		fields = strmangle.SetComplement(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := StockSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStocksUpsert(t *testing.T) {
	t.Parallel()

	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Stock{}
	if err = randomize.Struct(seed, &o, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stockDBTypes, false, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err = Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
