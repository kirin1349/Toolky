package toolky

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	mHost string = "localhost:3306"
	mAccount string = "root"
	mPw string = "mysql"
	mDB *sql.DB = nil
)

type DBDataRow map[string]string

type DBDataArray []DBDataRow

func MySQLSetup(host, account, pw string) {
	mHost = host
	mAccount = account
	mPw = pw
}

func MySQLOpen(dbName string) (result bool, err error) {
	dbSrc := mAccount + ":" + mPw + "@tcp(" + mHost + ")/" + dbName + "?charset=utf8";
	db, err := sql.Open("mysql", dbSrc)
	if err != nil {
		return false, err
	}
	mDB = db
	return true, nil
}

func MySQLClose() {
	if mDB != nil {
		mDB.Close()
	}
	mDB = nil
}

func MySQLGetAll(sql string) (result bool, data *DBDataArray, err error){
	if mDB == nil {
		return false, nil, nil
	}
	rows, err := mDB.Query(sql)
	if err != nil {
		return false, nil, err
	}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var dataArr DBDataArray = make([]DBDataRow, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		dataRow := record
		dataArr = append(dataArr, dataRow)
	}
	return true, &dataArr, nil
}

//插入demo
func MySQLInsert(table string, columns []string, values []string) (result bool, err error) {
	columnStr := ""
	valueStr := ""
	for i := range columns {
		if i > 0 {
			columnStr += ","
			valueStr += ","
		}
		columnStr += columns[i]
		valueStr += "?"
	}
	sql := "INSERT " + table + " (" + columnStr + ") values (" + valueStr + ")"
	stmt, err := mDB.Prepare(sql)
	if err != nil {
		return false, err
	}
	valueArgs := make([]interface{}, len(values))
	for i := range values {
		valueArgs[i] = &values[i]
	}
	res, err := stmt.Exec(valueArgs...)
	if err != nil {
		return false, err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return false, err
	}
	return true, nil
}

func MySQLUpdate(table string, whereColumns []string, columns []string, whereValues []string, values []string) (result bool, affect int64, err error) {
	columnStr := ""
	for i := range columns {
		if i > 0 {
			columnStr += ","
		}
		columnStr += columns[i] + "=?"
	}
	whereColumnStr := ""
	for i := range whereColumns {
		if i > 0 {
			whereColumnStr += " AND "
		} else {
			whereColumnStr += " WHERE "
		}
		whereColumnStr += whereColumns[i] + "=?"
	}
	sql := "UPDATE " + table + " SET " + columnStr + whereColumnStr
	stmt, err := mDB.Prepare(sql)
	if err != nil {
		return false, 0, err
	}
	valueCount := len(values)
	count := valueCount + len(whereValues)
	valueArgs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		if i < valueCount {
			valueArgs[i] = &values[i]
		} else {
			valueArgs[i] = &whereValues[i - valueCount]
		}
	}
	res, err := stmt.Exec(valueArgs...)
	if err != nil {
		return false, 0, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return false, 0, err
	}
	return true, num, nil
}