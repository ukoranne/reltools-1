package genmodels

import (
	"database/sql"
	"fmt"
)

func (obj VlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS VlanConfig " +
		"( " +
		" Status INTEGER " +
		" Name TEXT " +
		" VlanId INTEGER " +
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
func (obj VlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO VlanConfig (Status, Name, VlanId) VALUES (%v, %v, %v);",
		obj.Status, obj.Name, obj.VlanId)
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

func (obj VlanState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS VlanState " +
		"( " +
		" Status INTEGER " +
		" Name TEXT " +
		" VlanId INTEGER " +
		" MemberPorts TEXT " +
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
func (obj VlanState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	insertsql := fmt.Sprintf("INSERT INTO VlanState (Status, Name, VlanId, MemberPorts) VALUES (%v, %v, %v, %v);",
		obj.Status, obj.Name, obj.VlanId, obj.MemberPorts)
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
