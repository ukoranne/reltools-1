package genmodels

import (
	"database/sql"
	"fmt"
)

func (obj BgpGlobalConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalConfig (RouterId, As) VALUES (%v, %v);",
		obj.RouterId, obj.As)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalState " +
		"( " +
		" RouterId TEXT " +
		" TotalPaths INTEGER " +
		" As INTEGER " +
		" TotalPrefixes INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalState (RouterId, TotalPaths, As, TotalPrefixes) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.TotalPaths, obj.As, obj.TotalPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalRouteSelectionOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalRouteSelectionOptionsConfig " +
		"( " +
		" RouterId TEXT " +
		" EnableAigp bool " +
		" As INTEGER " +
		" IgnoreAsPathLength bool " +
		" AlwaysCompareMed bool " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" AdvertiseInactiveRoutes bool " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalRouteSelectionOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalRouteSelectionOptionsConfig (RouterId, EnableAigp, As, IgnoreAsPathLength, AlwaysCompareMed, IgnoreNextHopIgpMetric, ExternalCompareRouterId, AdvertiseInactiveRoutes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.EnableAigp, obj.As, obj.IgnoreAsPathLength, obj.AlwaysCompareMed, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.AdvertiseInactiveRoutes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalRouteSelectionOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalRouteSelectionOptionsState " +
		"( " +
		" RouterId TEXT " +
		" EnableAigp bool " +
		" As INTEGER " +
		" IgnoreAsPathLength bool " +
		" AlwaysCompareMed bool " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" AdvertiseInactiveRoutes bool " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalRouteSelectionOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalRouteSelectionOptionsState (RouterId, EnableAigp, As, IgnoreAsPathLength, AlwaysCompareMed, IgnoreNextHopIgpMetric, ExternalCompareRouterId, AdvertiseInactiveRoutes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.EnableAigp, obj.As, obj.IgnoreAsPathLength, obj.AlwaysCompareMed, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.AdvertiseInactiveRoutes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalDefaultRouteDistanceConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalDefaultRouteDistanceConfig " +
		"( " +
		" RouterId TEXT " +
		" ExternalRouteDistance INTEGER " +
		" As INTEGER " +
		" InternalRouteDistance INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalDefaultRouteDistanceConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalDefaultRouteDistanceConfig (RouterId, ExternalRouteDistance, As, InternalRouteDistance) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.ExternalRouteDistance, obj.As, obj.InternalRouteDistance)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalDefaultRouteDistanceState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalDefaultRouteDistanceState " +
		"( " +
		" RouterId TEXT " +
		" ExternalRouteDistance INTEGER " +
		" As INTEGER " +
		" InternalRouteDistance INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalDefaultRouteDistanceState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalDefaultRouteDistanceState (RouterId, ExternalRouteDistance, As, InternalRouteDistance) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.ExternalRouteDistance, obj.As, obj.InternalRouteDistance)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalConfederationConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalConfederationConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Identifier INTEGER " +
		" Enabled bool " +
		" MemberAs TEXT " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalConfederationConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalConfederationConfig (RouterId, As, Identifier, Enabled, MemberAs) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.Identifier, obj.Enabled, obj.MemberAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalConfederationState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalConfederationState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Identifier INTEGER " +
		" Enabled bool " +
		" MemberAs TEXT " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalConfederationState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalConfederationState (RouterId, As, Identifier, Enabled, MemberAs) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.Identifier, obj.Enabled, obj.MemberAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Enabled bool " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsConfig (RouterId, As, Enabled) VALUES (%v, %v, %v);",
		obj.RouterId, obj.As, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Enabled bool " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsState (RouterId, As, Enabled) VALUES (%v, %v, %v);",
		obj.RouterId, obj.As, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsEbgpConfig " +
		"( " +
		" RouterId TEXT " +
		" AllowMultipleAs bool " +
		" As INTEGER " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsEbgpConfig (RouterId, AllowMultipleAs, As, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.AllowMultipleAs, obj.As, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsEbgpState " +
		"( " +
		" RouterId TEXT " +
		" AllowMultipleAs bool " +
		" As INTEGER " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsEbgpState (RouterId, AllowMultipleAs, As, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.AllowMultipleAs, obj.As, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsIbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsIbgpConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsIbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsIbgpConfig (RouterId, As, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalUseMultiplePathsIbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalUseMultiplePathsIbgpState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalUseMultiplePathsIbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalUseMultiplePathsIbgpState (RouterId, As, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalGracefulRestartConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" StaleRoutesTime REAL " +
		" HelperOnly bool " +
		" Enabled bool " +
		" RestartTime INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalGracefulRestartConfig (RouterId, As, StaleRoutesTime, HelperOnly, Enabled, RestartTime) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.StaleRoutesTime, obj.HelperOnly, obj.Enabled, obj.RestartTime)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalGracefulRestartState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" StaleRoutesTime REAL " +
		" HelperOnly bool " +
		" Enabled bool " +
		" RestartTime INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalGracefulRestartState (RouterId, As, StaleRoutesTime, HelperOnly, Enabled, RestartTime) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.StaleRoutesTime, obj.HelperOnly, obj.Enabled, obj.RestartTime)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiGracefulRestartConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiGracefulRestartConfig (RouterId, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiGracefulRestartState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiGracefulRestartState (RouterId, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiConfig (RouterId, As, AfiSafiName, Enabled) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiName, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiState " +
		"( " +
		" RouterId TEXT " +
		" TotalPaths INTEGER " +
		" As INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		" TotalPrefixes INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiState (RouterId, TotalPaths, As, AfiSafiName, Enabled, TotalPrefixes) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.TotalPaths, obj.As, obj.AfiSafiName, obj.Enabled, obj.TotalPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiApplyPolicyConfig " +
		"( " +
		" RouterId TEXT " +
		" ImportPolicy TEXT " +
		" As INTEGER " +
		" ExportPolicy TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiApplyPolicyConfig (RouterId, ImportPolicy, As, ExportPolicy, DefaultImportPolicy, DefaultExportPolicy, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.ImportPolicy, obj.As, obj.ExportPolicy, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiApplyPolicyState " +
		"( " +
		" RouterId TEXT " +
		" ImportPolicy TEXT " +
		" As INTEGER " +
		" ExportPolicy TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiApplyPolicyState (RouterId, ImportPolicy, As, ExportPolicy, DefaultImportPolicy, DefaultExportPolicy, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.ImportPolicy, obj.As, obj.ExportPolicy, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4UnicastPrefixLimitConfig (SendDefaultRoute, RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4UnicastPrefixLimitState (SendDefaultRoute, RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4UnicastConfig " +
		"( " +
		" RouterId TEXT " +
		" SendDefaultRoute bool " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4UnicastConfig (RouterId, SendDefaultRoute, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.SendDefaultRoute, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4UnicastState " +
		"( " +
		" RouterId TEXT " +
		" SendDefaultRoute bool " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4UnicastState (RouterId, SendDefaultRoute, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.SendDefaultRoute, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6UnicastPrefixLimitConfig (SendDefaultRoute, RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6UnicastPrefixLimitState (SendDefaultRoute, RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6UnicastConfig " +
		"( " +
		" RouterId TEXT " +
		" SendDefaultRoute bool " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6UnicastConfig (RouterId, SendDefaultRoute, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.SendDefaultRoute, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6UnicastState " +
		"( " +
		" RouterId TEXT " +
		" SendDefaultRoute bool " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6UnicastState (RouterId, SendDefaultRoute, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.SendDefaultRoute, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv4LabelledUnicastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiIpv6LabelledUnicastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv4UnicastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv6UnicastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv4MulticastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL3vpnIpv6MulticastPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL2vpnVplsPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL2vpnVplsPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL2vpnVplsPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL2vpnVplsPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL2vpnVplsPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL2vpnVplsPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL2vpnVplsPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL2vpnVplsPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL2vpnEvpnPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL2vpnEvpnPrefixLimitConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL2vpnEvpnPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL2vpnEvpnPrefixLimitConfig (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiL2vpnEvpnPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiL2vpnEvpnPrefixLimitState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RestartTimer REAL " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaxPrefixes INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiL2vpnEvpnPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiL2vpnEvpnPrefixLimitState (RouterId, As, ShutdownThresholdPct, RestartTimer, AfiSafiNameKey, Enabled, MaxPrefixes) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.ShutdownThresholdPct, obj.RestartTimer, obj.AfiSafiNameKey, obj.Enabled, obj.MaxPrefixes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiRouteSelectionOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiRouteSelectionOptionsConfig " +
		"( " +
		" RouterId TEXT " +
		" EnableAigp bool " +
		" As INTEGER " +
		" IgnoreAsPathLength bool " +
		" AlwaysCompareMed bool " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" AdvertiseInactiveRoutes bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiRouteSelectionOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiRouteSelectionOptionsConfig (RouterId, EnableAigp, As, IgnoreAsPathLength, AlwaysCompareMed, IgnoreNextHopIgpMetric, ExternalCompareRouterId, AfiSafiNameKey, Enabled, AdvertiseInactiveRoutes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.EnableAigp, obj.As, obj.IgnoreAsPathLength, obj.AlwaysCompareMed, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.AfiSafiNameKey, obj.Enabled, obj.AdvertiseInactiveRoutes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiRouteSelectionOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiRouteSelectionOptionsState " +
		"( " +
		" RouterId TEXT " +
		" EnableAigp bool " +
		" As INTEGER " +
		" IgnoreAsPathLength bool " +
		" AlwaysCompareMed bool " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" AdvertiseInactiveRoutes bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiRouteSelectionOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiRouteSelectionOptionsState (RouterId, EnableAigp, As, IgnoreAsPathLength, AlwaysCompareMed, IgnoreNextHopIgpMetric, ExternalCompareRouterId, AfiSafiNameKey, Enabled, AdvertiseInactiveRoutes) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.EnableAigp, obj.As, obj.IgnoreAsPathLength, obj.AlwaysCompareMed, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.AfiSafiNameKey, obj.Enabled, obj.AdvertiseInactiveRoutes)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsConfig (RouterId, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsState (RouterId, As, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsEbgpConfig " +
		"( " +
		" RouterId TEXT " +
		" AllowMultipleAs bool " +
		" As INTEGER " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsEbgpConfig (RouterId, AllowMultipleAs, As, MaximumPaths, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.AllowMultipleAs, obj.As, obj.MaximumPaths, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsEbgpState " +
		"( " +
		" RouterId TEXT " +
		" AllowMultipleAs bool " +
		" As INTEGER " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsEbgpState (RouterId, AllowMultipleAs, As, MaximumPaths, AfiSafiNameKey, Enabled) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.AllowMultipleAs, obj.As, obj.MaximumPaths, obj.AfiSafiNameKey, obj.Enabled)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsIbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsIbgpConfig " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsIbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsIbgpConfig (RouterId, As, AfiSafiNameKey, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalAfiSafiUseMultiplePathsIbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalAfiSafiUseMultiplePathsIbgpState " +
		"( " +
		" RouterId TEXT " +
		" As INTEGER " +
		" AfiSafiNameKey TEXT " +
		" Enabled bool " +
		" MaximumPaths INTEGER " +
		"PRIMARY KEY(AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalAfiSafiUseMultiplePathsIbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalAfiSafiUseMultiplePathsIbgpState (RouterId, As, AfiSafiNameKey, Enabled, MaximumPaths) VALUES (%v, %v, %v, %v, %v);",
		obj.RouterId, obj.As, obj.AfiSafiNameKey, obj.Enabled, obj.MaximumPaths)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalApplyPolicyConfig " +
		"( " +
		" RouterId TEXT " +
		" ImportPolicy TEXT " +
		" As INTEGER " +
		" ExportPolicy TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalApplyPolicyConfig (RouterId, ImportPolicy, As, ExportPolicy, DefaultImportPolicy, DefaultExportPolicy) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.ImportPolicy, obj.As, obj.ExportPolicy, obj.DefaultImportPolicy, obj.DefaultExportPolicy)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpGlobalApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpGlobalApplyPolicyState " +
		"( " +
		" RouterId TEXT " +
		" ImportPolicy TEXT " +
		" As INTEGER " +
		" ExportPolicy TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpGlobalApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpGlobalApplyPolicyState (RouterId, ImportPolicy, As, ExportPolicy, DefaultImportPolicy, DefaultExportPolicy) VALUES (%v, %v, %v, %v, %v, %v);",
		obj.RouterId, obj.ImportPolicy, obj.As, obj.ExportPolicy, obj.DefaultImportPolicy, obj.DefaultExportPolicy)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborConfig " +
		"( " +
		" RouteFlapDamping bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" NeighborAddress TEXT " +
		" PeerGroup TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerType INTEGER " +
		" PeerAs INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborConfig (RouteFlapDamping, SendCommunity, Description, AuthPassword, NeighborAddress, PeerGroup, RemovePrivateAs, PeerType, PeerAs, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.NeighborAddress, obj.PeerGroup, obj.RemovePrivateAs, obj.PeerType, obj.PeerAs, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborStateMessagesSent) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborStateMessagesSent " +
		"( " +
		" UPDATE INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" NeighborAddress TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SupportedCapabilities TEXT " +
		" RouteFlapDamping bool " +
		" NOTIFICATION INTEGER " +
		" SessionState INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborStateMessagesSent) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborStateMessagesSent (UPDATE, SendCommunity, Description, AuthPassword, NeighborAddress, RemovePrivateAs, PeerAs, SupportedCapabilities, RouteFlapDamping, NOTIFICATION, SessionState, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.UPDATE, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.NeighborAddress, obj.RemovePrivateAs, obj.PeerAs, obj.SupportedCapabilities, obj.RouteFlapDamping, obj.NOTIFICATION, obj.SessionState, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborStateMessagesReceived) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborStateMessagesReceived " +
		"( " +
		" UPDATE INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" NeighborAddress TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SupportedCapabilities TEXT " +
		" RouteFlapDamping bool " +
		" NOTIFICATION INTEGER " +
		" SessionState INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborStateMessagesReceived) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborStateMessagesReceived (UPDATE, SendCommunity, Description, AuthPassword, NeighborAddress, RemovePrivateAs, PeerAs, SupportedCapabilities, RouteFlapDamping, NOTIFICATION, SessionState, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.UPDATE, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.NeighborAddress, obj.RemovePrivateAs, obj.PeerAs, obj.SupportedCapabilities, obj.RouteFlapDamping, obj.NOTIFICATION, obj.SessionState, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborStateQueues) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborStateQueues " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" NeighborAddress TEXT " +
		" RemovePrivateAs INTEGER " +
		" Input INTEGER " +
		" PeerAs INTEGER " +
		" SupportedCapabilities TEXT " +
		" RouteFlapDamping bool " +
		" Output INTEGER " +
		" SessionState INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborStateQueues) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborStateQueues (SendCommunity, Description, AuthPassword, NeighborAddress, RemovePrivateAs, Input, PeerAs, SupportedCapabilities, RouteFlapDamping, Output, SessionState, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.AuthPassword, obj.NeighborAddress, obj.RemovePrivateAs, obj.Input, obj.PeerAs, obj.SupportedCapabilities, obj.RouteFlapDamping, obj.Output, obj.SessionState, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborTimersConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborTimersConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" ConnectRetry REAL " +
		" HoldTime REAL " +
		" RemovePrivateAs INTEGER " +
		" KeepaliveInterval REAL " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" MinimumAdvertisementInterval REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborTimersConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborTimersConfig (SendCommunity, Description, ConnectRetry, HoldTime, RemovePrivateAs, KeepaliveInterval, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, MinimumAdvertisementInterval, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.ConnectRetry, obj.HoldTime, obj.RemovePrivateAs, obj.KeepaliveInterval, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.MinimumAdvertisementInterval, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborTimersState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborTimersState " +
		"( " +
		" Uptime INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" ConnectRetry REAL " +
		" HoldTime REAL " +
		" RemovePrivateAs INTEGER " +
		" KeepaliveInterval REAL " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" MinimumAdvertisementInterval REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		" NegotiatedHoldTime REAL " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborTimersState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborTimersState (Uptime, SendCommunity, Description, ConnectRetry, HoldTime, RemovePrivateAs, KeepaliveInterval, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, MinimumAdvertisementInterval, PeerType, LocalAs, NegotiatedHoldTime) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Uptime, obj.SendCommunity, obj.Description, obj.ConnectRetry, obj.HoldTime, obj.RemovePrivateAs, obj.KeepaliveInterval, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.MinimumAdvertisementInterval, obj.PeerType, obj.LocalAs, obj.NegotiatedHoldTime)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborTransportConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborTransportConfig " +
		"( " +
		" RemovePrivateAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MtuDiscovery bool " +
		" PeerAs INTEGER " +
		" LocalAddress_InetIpAddress TEXT " +
		" PassiveMode bool " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" TcpMss INTEGER " +
		" LocalAddress_InetIpAddress_String TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborTransportConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborTransportConfig (RemovePrivateAs, SendCommunity, Description, MtuDiscovery, PeerAs, LocalAddress_InetIpAddress, PassiveMode, AuthPassword, RouteFlapDamping, NeighborAddressKey, TcpMss, LocalAddress_InetIpAddress_String, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RemovePrivateAs, obj.SendCommunity, obj.Description, obj.MtuDiscovery, obj.PeerAs, obj.LocalAddress_InetIpAddress, obj.PassiveMode, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.TcpMss, obj.LocalAddress_InetIpAddress_String, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborTransportState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborTransportState " +
		"( " +
		" RemotePort INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MtuDiscovery bool " +
		" RemoteAddress TEXT " +
		" PassiveMode bool " +
		" RemovePrivateAs INTEGER " +
		" LocalAddress_InetIpAddress TEXT " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" TcpMss INTEGER " +
		" LocalAddress_InetIpAddress_String TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		" LocalPort INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborTransportState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborTransportState (RemotePort, SendCommunity, Description, MtuDiscovery, RemoteAddress, PassiveMode, RemovePrivateAs, LocalAddress_InetIpAddress, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, TcpMss, LocalAddress_InetIpAddress_String, PeerGroup, PeerType, LocalAs, LocalPort) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RemotePort, obj.SendCommunity, obj.Description, obj.MtuDiscovery, obj.RemoteAddress, obj.PassiveMode, obj.RemovePrivateAs, obj.LocalAddress_InetIpAddress, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.TcpMss, obj.LocalAddress_InetIpAddress_String, obj.PeerGroup, obj.PeerType, obj.LocalAs, obj.LocalPort)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborErrorHandlingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborErrorHandlingConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" TreatAsWithdraw bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborErrorHandlingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborErrorHandlingConfig (SendCommunity, Description, TreatAsWithdraw, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.TreatAsWithdraw, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborErrorHandlingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborErrorHandlingState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" TreatAsWithdraw bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ErroneousUpdateMessages INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborErrorHandlingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborErrorHandlingState (SendCommunity, Description, TreatAsWithdraw, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ErroneousUpdateMessages, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.TreatAsWithdraw, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ErroneousUpdateMessages, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborLoggingOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborLoggingOptionsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" LogNeighborStateChanges bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborLoggingOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborLoggingOptionsConfig (SendCommunity, Description, LogNeighborStateChanges, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.LogNeighborStateChanges, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborLoggingOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborLoggingOptionsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" LogNeighborStateChanges bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborLoggingOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborLoggingOptionsState (SendCommunity, Description, LogNeighborStateChanges, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.LogNeighborStateChanges, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborEbgpMultihopConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborEbgpMultihopConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MultihopTtl INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborEbgpMultihopConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborEbgpMultihopConfig (SendCommunity, Description, MultihopTtl, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MultihopTtl, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborEbgpMultihopState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborEbgpMultihopState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MultihopTtl INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborEbgpMultihopState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborEbgpMultihopState (SendCommunity, Description, MultihopTtl, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MultihopTtl, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborRouteReflectorConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborRouteReflectorConfig " +
		"( " +
		" RouteReflectorClusterId TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" RouteReflectorClient bool " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborRouteReflectorConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborRouteReflectorConfig (RouteReflectorClusterId, SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, RouteReflectorClient, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteReflectorClusterId, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.RouteReflectorClient, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborRouteReflectorState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborRouteReflectorState " +
		"( " +
		" RouteReflectorClusterId TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" RouteReflectorClient bool " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborRouteReflectorState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborRouteReflectorState (RouteReflectorClusterId, SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, RouteReflectorClient, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteReflectorClusterId, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.RouteReflectorClient, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAsPathOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAsPathOptionsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ReplacePeerAs bool " +
		" PeerGroup TEXT " +
		" AllowOwnAs INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAsPathOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAsPathOptionsConfig (SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ReplacePeerAs, PeerGroup, AllowOwnAs, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ReplacePeerAs, obj.PeerGroup, obj.AllowOwnAs, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAsPathOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAsPathOptionsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ReplacePeerAs bool " +
		" PeerGroup TEXT " +
		" AllowOwnAs INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAsPathOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAsPathOptionsState (SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ReplacePeerAs, PeerGroup, AllowOwnAs, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ReplacePeerAs, obj.PeerGroup, obj.AllowOwnAs, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAddPathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAddPathsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Receive bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendMax INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAddPathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAddPathsConfig (SendCommunity, Description, Receive, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendMax, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.Receive, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendMax, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAddPathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAddPathsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Receive bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendMax INTEGER " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAddPathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAddPathsState (SendCommunity, Description, Receive, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendMax, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.Receive, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendMax, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiGracefulRestartConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiGracefulRestartConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiGracefulRestartState " +
		"( " +
		" Received bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" Advertised bool " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiGracefulRestartState (Received, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, Advertised, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Received, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.Advertised, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiConfig (SendCommunity, Description, RemovePrivateAs, AfiSafiName, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.AfiSafiName, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiStatePrefixes) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiStatePrefixes " +
		"( " +
		" Received INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" Active bool " +
		" PeerGroup TEXT " +
		" Installed INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		" Sent INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiStatePrefixes) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiStatePrefixes (Received, SendCommunity, Description, RemovePrivateAs, AfiSafiName, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, Active, PeerGroup, Installed, PeerType, LocalAs, Sent) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Received, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.AfiSafiName, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.Active, obj.PeerGroup, obj.Installed, obj.PeerType, obj.LocalAs, obj.Sent)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiApplyPolicyConfig " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiApplyPolicyConfig (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, DefaultExportPolicy, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ExportPolicy, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ExportPolicy, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiApplyPolicyState " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiApplyPolicyState (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, DefaultExportPolicy, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ExportPolicy, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ExportPolicy, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4UnicastPrefixLimitConfig (SendDefaultRoute, RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4UnicastPrefixLimitState (SendDefaultRoute, RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4UnicastConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4UnicastConfig (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4UnicastState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4UnicastState (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6UnicastPrefixLimitConfig (SendDefaultRoute, RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6UnicastPrefixLimitState (SendDefaultRoute, RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6UnicastConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6UnicastConfig (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6UnicastState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6UnicastState (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv4LabelledUnicastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiIpv6LabelledUnicastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv4UnicastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv6UnicastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv4MulticastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL3vpnIpv6MulticastPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL2vpnVplsPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL2vpnVplsPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL2vpnVplsPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL2vpnVplsPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL2vpnVplsPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL2vpnVplsPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL2vpnVplsPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL2vpnVplsPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL2vpnEvpnPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL2vpnEvpnPrefixLimitConfig " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL2vpnEvpnPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL2vpnEvpnPrefixLimitConfig (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiL2vpnEvpnPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiL2vpnEvpnPrefixLimitState " +
		"( " +
		" RestartTimer REAL " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" ShutdownThresholdPct INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiL2vpnEvpnPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiL2vpnEvpnPrefixLimitState (RestartTimer, SendCommunity, Description, MaxPrefixes, ShutdownThresholdPct, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTimer, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.ShutdownThresholdPct, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiUseMultiplePathsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiUseMultiplePathsConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiUseMultiplePathsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiUseMultiplePathsState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiUseMultiplePathsEbgpConfig " +
		"( " +
		" AllowMultipleAs bool " +
		" PeerGroup TEXT " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendCommunity INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiUseMultiplePathsEbgpConfig (AllowMultipleAs, PeerGroup, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendCommunity, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.PeerGroup, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendCommunity, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborAfiSafiUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborAfiSafiUseMultiplePathsEbgpState " +
		"( " +
		" AllowMultipleAs bool " +
		" PeerGroup TEXT " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendCommunity INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborAfiSafiUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborAfiSafiUseMultiplePathsEbgpState (AllowMultipleAs, PeerGroup, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendCommunity, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.PeerGroup, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendCommunity, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborGracefulRestartConfig " +
		"( " +
		" RestartTime INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" HelperOnly bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" StaleRoutesTime REAL " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborGracefulRestartConfig (RestartTime, SendCommunity, Description, HelperOnly, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, StaleRoutesTime, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTime, obj.SendCommunity, obj.Description, obj.HelperOnly, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.StaleRoutesTime, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborGracefulRestartState " +
		"( " +
		" RestartTime INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" HelperOnly bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" LocalAs INTEGER " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" StaleRoutesTime REAL " +
		" Mode INTEGER " +
		" PeerRestarting bool " +
		" PeerGroup TEXT " +
		" LocalRestarting bool " +
		" PeerType INTEGER " +
		" PeerAs INTEGER " +
		" PeerRestartTime INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborGracefulRestartState (RestartTime, SendCommunity, Description, AuthPassword, HelperOnly, RemovePrivateAs, Enabled, LocalAs, RouteFlapDamping, NeighborAddressKey, StaleRoutesTime, Mode, PeerRestarting, PeerGroup, LocalRestarting, PeerType, PeerAs, PeerRestartTime) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTime, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.HelperOnly, obj.RemovePrivateAs, obj.Enabled, obj.LocalAs, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.StaleRoutesTime, obj.Mode, obj.PeerRestarting, obj.PeerGroup, obj.LocalRestarting, obj.PeerType, obj.PeerAs, obj.PeerRestartTime)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborApplyPolicyConfig " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborApplyPolicyConfig (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, RemovePrivateAs, DefaultExportPolicy, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ExportPolicy, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.RemovePrivateAs, obj.DefaultExportPolicy, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ExportPolicy, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborApplyPolicyState " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborApplyPolicyState (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, RemovePrivateAs, DefaultExportPolicy, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, ExportPolicy, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.RemovePrivateAs, obj.DefaultExportPolicy, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.ExportPolicy, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborUseMultiplePathsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborUseMultiplePathsConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborUseMultiplePathsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" PeerGroup TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborUseMultiplePathsState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, PeerGroup, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.PeerGroup, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborUseMultiplePathsEbgpConfig " +
		"( " +
		" AllowMultipleAs bool " +
		" PeerGroup TEXT " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendCommunity INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborUseMultiplePathsEbgpConfig (AllowMultipleAs, PeerGroup, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendCommunity, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.PeerGroup, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendCommunity, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpNeighborUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpNeighborUseMultiplePathsEbgpState " +
		"( " +
		" AllowMultipleAs bool " +
		" PeerGroup TEXT " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" NeighborAddressKey TEXT " +
		" SendCommunity INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(NeighborAddressKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpNeighborUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpNeighborUseMultiplePathsEbgpState (AllowMultipleAs, PeerGroup, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, NeighborAddressKey, SendCommunity, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.PeerGroup, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.NeighborAddressKey, obj.SendCommunity, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupConfig " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupName TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerType INTEGER " +
		" PeerAs INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupConfig (RouteFlapDamping, PeerGroupName, SendCommunity, Description, AuthPassword, RemovePrivateAs, PeerType, PeerAs, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupName, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.RemovePrivateAs, obj.PeerType, obj.PeerAs, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupState " +
		"( " +
		" TotalPaths INTEGER " +
		" PeerGroupName TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" LocalAs INTEGER " +
		" RouteFlapDamping bool " +
		" TotalPrefixes INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY( ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupState (TotalPaths, PeerGroupName, SendCommunity, Description, RemovePrivateAs, PeerAs, LocalAs, RouteFlapDamping, TotalPrefixes, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.TotalPaths, obj.PeerGroupName, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.LocalAs, obj.RouteFlapDamping, obj.TotalPrefixes, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupTimersConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupTimersConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" ConnectRetry REAL " +
		" HoldTime REAL " +
		" RemovePrivateAs INTEGER " +
		" KeepaliveInterval REAL " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MinimumAdvertisementInterval REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupTimersConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupTimersConfig (SendCommunity, Description, ConnectRetry, HoldTime, RemovePrivateAs, KeepaliveInterval, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MinimumAdvertisementInterval, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.ConnectRetry, obj.HoldTime, obj.RemovePrivateAs, obj.KeepaliveInterval, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MinimumAdvertisementInterval, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupTimersState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupTimersState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" ConnectRetry REAL " +
		" HoldTime REAL " +
		" RemovePrivateAs INTEGER " +
		" KeepaliveInterval REAL " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MinimumAdvertisementInterval REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupTimersState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupTimersState (SendCommunity, Description, ConnectRetry, HoldTime, RemovePrivateAs, KeepaliveInterval, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MinimumAdvertisementInterval, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.ConnectRetry, obj.HoldTime, obj.RemovePrivateAs, obj.KeepaliveInterval, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MinimumAdvertisementInterval, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupTransportConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupTransportConfig " +
		"( " +
		" RemovePrivateAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MtuDiscovery bool " +
		" PeerAs INTEGER " +
		" LocalAddress_InetIpAddress TEXT " +
		" PassiveMode bool " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" TcpMss INTEGER " +
		" LocalAddress_InetIpAddress_String TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupTransportConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupTransportConfig (RemovePrivateAs, SendCommunity, Description, MtuDiscovery, PeerAs, LocalAddress_InetIpAddress, PassiveMode, AuthPassword, RouteFlapDamping, PeerGroupNameKey, TcpMss, LocalAddress_InetIpAddress_String, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RemovePrivateAs, obj.SendCommunity, obj.Description, obj.MtuDiscovery, obj.PeerAs, obj.LocalAddress_InetIpAddress, obj.PassiveMode, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.TcpMss, obj.LocalAddress_InetIpAddress_String, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupTransportState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupTransportState " +
		"( " +
		" RemovePrivateAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MtuDiscovery bool " +
		" PeerAs INTEGER " +
		" LocalAddress_InetIpAddress TEXT " +
		" PassiveMode bool " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" TcpMss INTEGER " +
		" LocalAddress_InetIpAddress_String TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupTransportState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupTransportState (RemovePrivateAs, SendCommunity, Description, MtuDiscovery, PeerAs, LocalAddress_InetIpAddress, PassiveMode, AuthPassword, RouteFlapDamping, PeerGroupNameKey, TcpMss, LocalAddress_InetIpAddress_String, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RemovePrivateAs, obj.SendCommunity, obj.Description, obj.MtuDiscovery, obj.PeerAs, obj.LocalAddress_InetIpAddress, obj.PassiveMode, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.TcpMss, obj.LocalAddress_InetIpAddress_String, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupErrorHandlingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupErrorHandlingConfig " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" TreatAsWithdraw bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupErrorHandlingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupErrorHandlingConfig (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, TreatAsWithdraw, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.TreatAsWithdraw, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupErrorHandlingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupErrorHandlingState " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" TreatAsWithdraw bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupErrorHandlingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupErrorHandlingState (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, TreatAsWithdraw, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.TreatAsWithdraw, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupLoggingOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupLoggingOptionsConfig " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" LogNeighborStateChanges bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupLoggingOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupLoggingOptionsConfig (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, LogNeighborStateChanges, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.LogNeighborStateChanges, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupLoggingOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupLoggingOptionsState " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" LogNeighborStateChanges bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupLoggingOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupLoggingOptionsState (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, LogNeighborStateChanges, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.LogNeighborStateChanges, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupEbgpMultihopConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupEbgpMultihopConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MultihopTtl INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupEbgpMultihopConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupEbgpMultihopConfig (SendCommunity, Description, MultihopTtl, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MultihopTtl, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupEbgpMultihopState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupEbgpMultihopState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MultihopTtl INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupEbgpMultihopState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupEbgpMultihopState (SendCommunity, Description, MultihopTtl, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MultihopTtl, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupRouteReflectorConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupRouteReflectorConfig " +
		"( " +
		" RouteReflectorClusterId TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RouteReflectorClient bool " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupRouteReflectorConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupRouteReflectorConfig (RouteReflectorClusterId, SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, RouteReflectorClient, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteReflectorClusterId, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RouteReflectorClient, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupRouteReflectorState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupRouteReflectorState " +
		"( " +
		" RouteReflectorClusterId TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RouteReflectorClient bool " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupRouteReflectorState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupRouteReflectorState (RouteReflectorClusterId, SendCommunity, Description, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, RouteReflectorClient, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteReflectorClusterId, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RouteReflectorClient, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAsPathOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAsPathOptionsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AllowOwnAs INTEGER " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ReplacePeerAs bool " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAsPathOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAsPathOptionsConfig (SendCommunity, Description, AuthPassword, RemovePrivateAs, PeerAs, AllowOwnAs, RouteFlapDamping, PeerGroupNameKey, ReplacePeerAs, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.AuthPassword, obj.RemovePrivateAs, obj.PeerAs, obj.AllowOwnAs, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ReplacePeerAs, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAsPathOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAsPathOptionsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AllowOwnAs INTEGER " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ReplacePeerAs bool " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAsPathOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAsPathOptionsState (SendCommunity, Description, AuthPassword, RemovePrivateAs, PeerAs, AllowOwnAs, RouteFlapDamping, PeerGroupNameKey, ReplacePeerAs, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.AuthPassword, obj.RemovePrivateAs, obj.PeerAs, obj.AllowOwnAs, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ReplacePeerAs, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAddPathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAddPathsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Receive bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" SendMax INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAddPathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAddPathsConfig (SendCommunity, Description, Receive, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, SendMax, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.Receive, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.SendMax, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAddPathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAddPathsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Receive bool " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" SendMax INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAddPathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAddPathsState (SendCommunity, Description, Receive, RemovePrivateAs, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, SendMax, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.Receive, obj.RemovePrivateAs, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.SendMax, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiGracefulRestartConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiGracefulRestartConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiGracefulRestartState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiGracefulRestartState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiConfig (SendCommunity, Description, RemovePrivateAs, AfiSafiName, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.AfiSafiName, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" AfiSafiName TEXT " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiState (SendCommunity, Description, RemovePrivateAs, AfiSafiName, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.AfiSafiName, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiApplyPolicyConfig " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ExportPolicy TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiApplyPolicyConfig (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, DefaultExportPolicy, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ExportPolicy, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ExportPolicy, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiApplyPolicyState " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ExportPolicy TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiApplyPolicyState (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, DefaultExportPolicy, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ExportPolicy, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.DefaultExportPolicy, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ExportPolicy, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitConfig (SendDefaultRoute, SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4UnicastPrefixLimitState (SendDefaultRoute, SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4UnicastConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4UnicastConfig (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4UnicastState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4UnicastState (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitConfig (SendDefaultRoute, SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6UnicastPrefixLimitState (SendDefaultRoute, SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6UnicastConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6UnicastConfig " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6UnicastConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6UnicastConfig (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6UnicastState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6UnicastState " +
		"( " +
		" SendDefaultRoute bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6UnicastState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6UnicastState (SendDefaultRoute, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendDefaultRoute, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv4LabelledUnicastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiIpv6LabelledUnicastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv4UnicastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv6UnicastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv4MulticastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL3vpnIpv6MulticastPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL2vpnVplsPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitConfig (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" MaxPrefixes INTEGER " +
		" RestartTimer REAL " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ShutdownThresholdPct INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiL2vpnEvpnPrefixLimitState (SendCommunity, Description, MaxPrefixes, RestartTimer, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ShutdownThresholdPct, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.MaxPrefixes, obj.RestartTimer, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ShutdownThresholdPct, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsEbgpConfig " +
		"( " +
		" AllowMultipleAs bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsEbgpConfig (AllowMultipleAs, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsEbgpState " +
		"( " +
		" AllowMultipleAs bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsEbgpState (AllowMultipleAs, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsIbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsIbgpConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsIbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsIbgpConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiUseMultiplePathsIbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiUseMultiplePathsIbgpState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiUseMultiplePathsIbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiUseMultiplePathsIbgpState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, AfiSafiNameKey, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.AfiSafiNameKey, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiRouteSelectionOptionsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiRouteSelectionOptionsConfig " +
		"( " +
		" EnableAigp bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" LocalAs INTEGER " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AlwaysCompareMed bool " +
		" AdvertiseInactiveRoutes bool " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" PeerAs INTEGER " +
		" IgnoreAsPathLength bool " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiRouteSelectionOptionsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiRouteSelectionOptionsConfig (EnableAigp, SendCommunity, Description, AuthPassword, IgnoreNextHopIgpMetric, ExternalCompareRouterId, RemovePrivateAs, Enabled, LocalAs, RouteFlapDamping, PeerGroupNameKey, AlwaysCompareMed, AdvertiseInactiveRoutes, AfiSafiNameKey, PeerType, PeerAs, IgnoreAsPathLength) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.EnableAigp, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.RemovePrivateAs, obj.Enabled, obj.LocalAs, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AlwaysCompareMed, obj.AdvertiseInactiveRoutes, obj.AfiSafiNameKey, obj.PeerType, obj.PeerAs, obj.IgnoreAsPathLength)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupAfiSafiRouteSelectionOptionsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupAfiSafiRouteSelectionOptionsState " +
		"( " +
		" EnableAigp bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" AuthPassword TEXT " +
		" IgnoreNextHopIgpMetric bool " +
		" ExternalCompareRouterId bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" LocalAs INTEGER " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" AlwaysCompareMed bool " +
		" AdvertiseInactiveRoutes bool " +
		" AfiSafiNameKey TEXT " +
		" PeerType INTEGER " +
		" PeerAs INTEGER " +
		" IgnoreAsPathLength bool " +
		"PRIMARY KEY(PeerGroupNameKey, AfiSafiNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupAfiSafiRouteSelectionOptionsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupAfiSafiRouteSelectionOptionsState (EnableAigp, SendCommunity, Description, AuthPassword, IgnoreNextHopIgpMetric, ExternalCompareRouterId, RemovePrivateAs, Enabled, LocalAs, RouteFlapDamping, PeerGroupNameKey, AlwaysCompareMed, AdvertiseInactiveRoutes, AfiSafiNameKey, PeerType, PeerAs, IgnoreAsPathLength) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.EnableAigp, obj.SendCommunity, obj.Description, obj.AuthPassword, obj.IgnoreNextHopIgpMetric, obj.ExternalCompareRouterId, obj.RemovePrivateAs, obj.Enabled, obj.LocalAs, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.AlwaysCompareMed, obj.AdvertiseInactiveRoutes, obj.AfiSafiNameKey, obj.PeerType, obj.PeerAs, obj.IgnoreAsPathLength)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupGracefulRestartConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupGracefulRestartConfig " +
		"( " +
		" RestartTime INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" HelperOnly bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" StaleRoutesTime REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupGracefulRestartConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupGracefulRestartConfig (RestartTime, SendCommunity, Description, HelperOnly, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, StaleRoutesTime, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTime, obj.SendCommunity, obj.Description, obj.HelperOnly, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.StaleRoutesTime, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupGracefulRestartState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupGracefulRestartState " +
		"( " +
		" RestartTime INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" HelperOnly bool " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" StaleRoutesTime REAL " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupGracefulRestartState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupGracefulRestartState (RestartTime, SendCommunity, Description, HelperOnly, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, StaleRoutesTime, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RestartTime, obj.SendCommunity, obj.Description, obj.HelperOnly, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.StaleRoutesTime, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupApplyPolicyConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupApplyPolicyConfig " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupApplyPolicyConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupApplyPolicyConfig (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, RemovePrivateAs, DefaultExportPolicy, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ExportPolicy, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.RemovePrivateAs, obj.DefaultExportPolicy, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ExportPolicy, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupApplyPolicyState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupApplyPolicyState " +
		"( " +
		" ImportPolicy TEXT " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" DefaultImportPolicy INTEGER " +
		" RemovePrivateAs INTEGER " +
		" DefaultExportPolicy INTEGER " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" ExportPolicy TEXT " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupApplyPolicyState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupApplyPolicyState (ImportPolicy, SendCommunity, Description, DefaultImportPolicy, RemovePrivateAs, DefaultExportPolicy, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, ExportPolicy, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.ImportPolicy, obj.SendCommunity, obj.Description, obj.DefaultImportPolicy, obj.RemovePrivateAs, obj.DefaultExportPolicy, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.ExportPolicy, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsConfig " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Enabled bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsConfig (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, Enabled, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.Enabled, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsState " +
		"( " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" RemovePrivateAs INTEGER " +
		" PeerAs INTEGER " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" Enabled bool " +
		" LocalAs INTEGER " +
		" PeerType INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsState (RouteFlapDamping, PeerGroupNameKey, RemovePrivateAs, PeerAs, SendCommunity, Description, Enabled, LocalAs, PeerType, AuthPassword) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.RemovePrivateAs, obj.PeerAs, obj.SendCommunity, obj.Description, obj.Enabled, obj.LocalAs, obj.PeerType, obj.AuthPassword)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsEbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsEbgpConfig " +
		"( " +
		" AllowMultipleAs bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsEbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsEbgpConfig (AllowMultipleAs, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsEbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsEbgpState " +
		"( " +
		" AllowMultipleAs bool " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsEbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsEbgpState (AllowMultipleAs, SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AllowMultipleAs, obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsIbgpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsIbgpConfig " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsIbgpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsIbgpConfig (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}

func (obj BgpPeerGroupUseMultiplePathsIbgpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BgpPeerGroupUseMultiplePathsIbgpState " +
		"( " +
		" SendCommunity INTEGER " +
		" Description TEXT " +
		" RemovePrivateAs INTEGER " +
		" Enabled bool " +
		" PeerAs INTEGER " +
		" AuthPassword TEXT " +
		" RouteFlapDamping bool " +
		" PeerGroupNameKey TEXT " +
		" MaximumPaths INTEGER " +
		" PeerType INTEGER " +
		" LocalAs INTEGER " +
		"PRIMARY KEY(PeerGroupNameKey) ) "
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", dbCmd)
	_, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	return nil
}
func (obj BgpPeerGroupUseMultiplePathsIbgpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO BgpPeerGroupUseMultiplePathsIbgpState (SendCommunity, Description, RemovePrivateAs, Enabled, PeerAs, AuthPassword, RouteFlapDamping, PeerGroupNameKey, MaximumPaths, PeerType, LocalAs) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SendCommunity, obj.Description, obj.RemovePrivateAs, obj.Enabled, obj.PeerAs, obj.AuthPassword, obj.RouteFlapDamping, obj.PeerGroupNameKey, obj.MaximumPaths, obj.PeerType, obj.LocalAs)
	fmt.Println("**** Create Object called with ", obj)

	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart a transaction")
	}
	fmt.Println("**** Executing DB command ", insertsql)
	result, err1 := dbHdl.Exec(insertsql)
	if err1 != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction")
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	} else {
		fmt.Println("### Object ID return ", objectId)
	}

	return objectId, nil
}
