package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func (sd StructDescription) GenerateDBStorage(property Property, srcDir string) {
	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("storage_%s_db_template.go", strings.ToLower(sd.StructName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputF.Close()
	s := fmt.Sprintf(
		`// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package %s

import (
    "database/sql"
    "fmt"
    "github.com/0studio/databasetemplate"
	"github.com/0studio/goutils"
    "github.com/0studio/logger"
    key "github.com/0studio/storage_key"
    "time"
)
var ___importTimeDB%s time.Time
var ___importKeyDB%s key.KeyUint64
var ___importGoutilsDB%s goutils.Int32List

`, property.PackageName, sd.StructName, sd.StructName, sd.StructName)

	s += sd.GenerateDBStorageStruct()
	s += sd.GenerateNewDBStorageStruct()
	s += sd.GenerateDBCreateTable()
	s += sd.GenerateDBMapRow()
	s += sd.GenerateDBGetAll()
	s += sd.GenerateDBGet()
	s += sd.GenerateDBAdd()
	s += sd.GenerateDBMultiAdd()
	s += sd.GenerateDBSet()
	s += sd.GenerateDBMultGet()
	s += sd.GenerateDBMultUpdate()
	s += sd.GenerateDBDelete()
	s += sd.GenerateDBMultDelete()
	s += sd.GenerateDBGetIdList()
	s += sd.GenerateDBsetIdList()

	formatSrc, _ := format.Source([]byte(s))
	if err == nil {
		outputF.WriteString(string(formatSrc))
	} else {
		outputF.WriteString(s)
	}

}
func (sd StructDescription) GenerateDBStorageStruct() string {
	s :=
		`type %s struct {
	databasetemplate.DatabaseTemplate
    log logger.Logger
}
`
	return fmt.Sprintf(s, sd.GetDBStorageName())
}
func (sd StructDescription) GetDBStorageName() string {
	return fmt.Sprintf("DB%sStorage", sd.StructName)
}
func (sd StructDescription) GenerateNewDBStorageStruct() string {
	s :=
		`
// log can be nil
func New%s(dt databasetemplate.DatabaseTemplate,log logger.Logger, createTable bool) *%s {
	var dao *%s = &%s{dt,log}
	if createTable {
		dao.CreateTable()
	}
	return dao
}
`
	return fmt.Sprintf(s, sd.GetDBStorageName(), sd.GetDBStorageName(), sd.GetDBStorageName(), sd.GetDBStorageName())
}
func (sd StructDescription) GenerateDBCreateTable() string {
	sql, _ := sd.GenerateCreateTableSql()
	sql = strings.Replace(sql, "\n", "", -1)
	s :=
		`
func (this *%s) CreateTable() bool{
	sql := "%s"
	err := this.DatabaseTemplate.ExecDDL(sql)
	if err != nil {
        if this.log != nil {
            this.log.Error("err",err)
        }else {
		    fmt.Println("err", err)
        }
        return false
	}
    return true
}

`
	return fmt.Sprintf(s, sd.GetDBStorageName(), sql)
}

func (sd StructDescription) generateDBMapRowSpicialType() (s string) {
	for _, field := range sd.Fields {
		if field.IsBool() {
			s += fmt.Sprintf("    var %s int\n", field.FieldName)
		} else if field.IsTimeInt() {
			s += fmt.Sprintf("    var %s int64\n", field.FieldName)
		} else if field.IsIntList() {
			s += fmt.Sprintf("    var %s string\n", field.FieldName)
		} else if field.IsStringList() {
			s += fmt.Sprintf("    var %s string\n", field.FieldName)
		}

	}
	return
}

func (sd StructDescription) generateDBMapRowSpicialTypeTrans() (s string) {
	for _, field := range sd.Fields {
		if field.IsBool() {

			s += fmt.Sprintf("    e.%s = int2bool(%s)\n", field.FieldName, field.FieldName)
		} else if field.IsTimeInt() {
			s += fmt.Sprintf("    e.%s = newTime(%s)\n", field.FieldName, field.FieldName)
		} else if field.IsIntList() {
			switch field.FieldGoType {
			case "[]int":
				s += fmt.Sprintf("    e.%s = str2intList(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]int8":
				s += fmt.Sprintf("    e.%s = str2int8List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]int16":
				s += fmt.Sprintf("    e.%s = str2int16List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]int32":
				s += fmt.Sprintf("    e.%s = str2int32List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]int64":
				s += fmt.Sprintf("    e.%s = str2int64List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]uint8":
				s += fmt.Sprintf("    e.%s = str2uint8List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]uint16":
				s += fmt.Sprintf("    e.%s = str2uint16List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]uint32":
				s += fmt.Sprintf("    e.%s = str2uint32List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "[]uint64":
				s += fmt.Sprintf("    e.%s = str2uint64List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "goutils.Int32List":
				s += fmt.Sprintf("    e.%s = str2int32List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "goutils.Int16List":
				s += fmt.Sprintf("    e.%s = str2int16List(%s, `,`)\n", field.FieldName, field.FieldName)
			case "goutils.IntList":
				s += fmt.Sprintf("    e.%s = str2intList(%s, `,`)\n", field.FieldName, field.FieldName)
			case "goutils.Int8List":
				s += fmt.Sprintf("    e.%s = str2int8List(%s, `,`)\n", field.FieldName, field.FieldName)
			default:
				fmt.Println("should be here generateDBMapRowSpicialTypeTrans", field.FieldGoType)
			}
		} else if field.IsStringList() {
			s += fmt.Sprintf("    e.%s = stringSplit(%s, `,`)\n", field.FieldName, field.FieldName)
		}

	}
	return
}
func (sd StructDescription) generateDBMapRowScan() (s string) {
	s += "\n"
	for _, field := range sd.Fields {
		if field.IsBool() {
			s += fmt.Sprintf("        &%s,\n", field.FieldName)
		} else if field.IsTimeInt() {
			s += fmt.Sprintf("        &%s,\n", field.FieldName)
		} else if field.IsIntList() {
			s += fmt.Sprintf("        &%s,\n", field.FieldName)
		} else if field.IsStringList() {
			s += fmt.Sprintf("        &%s,\n", field.FieldName)
		} else {
			s += fmt.Sprintf("        &e.%s,\n", field.FieldName)
		}
	}
	return
}
func (sd StructDescription) GenerateDBMapRow() string {
	s :=
		`func (this *%s) mapRow(resultSet *sql.Rows) (interface{}, error) {
	e := %s{}
%s
	err := resultSet.Scan(%s    )
	if err != nil {
		return nil, err
	}
%s
	e.ClearFlag()
	return e, nil
}
`
	return fmt.Sprintf(s, sd.GetDBStorageName(), sd.StructName,
		sd.generateDBMapRowSpicialType(),
		sd.generateDBMapRowScan(),
		sd.generateDBMapRowSpicialTypeTrans())

}
func (sd StructDescription) GetPKVarDeclear() (s string) {
	pkList := sd.GetPKFieldList()
	for idx, field := range pkList {
		s += fmt.Sprintf("%s %s", field.FieldName, field.FieldGoType)
		if idx != len(pkList)-1 {
			s += ", "
		}
	}
	return
}
func isTypeKeySum(goType string) bool {
	if strings.HasPrefix(goType, "key.") {
		return true
	}
	return false
}
func (sd StructDescription) GenerateDBGet() string {
	pkList := sd.GetPKFieldList()
	var keySum string = "nil"
	if len(pkList) == 0 {
		return ""
	}

	if isTypeKeySum(pkList[0].FieldGoType) {
		keySum = pkList[0].FieldName
	}

	s :=
		fmt.Sprintf(
			`func (this *%s) Get(%s, now time.Time) (e %s, ok bool) {
	sql := "select  %s from %s where %s "
	var obj interface{}
	var err error
	obj, err = this.DatabaseTemplate.QueryObject(%s, sql, this.mapRow, %s)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.Get %s %%v %%s",%s,err,sql)
        }
		return
	}
	if obj == nil {
		return
	}
	e = obj.(%s)
	ok = true
	return
}
`, sd.GetDBStorageName(), sd.GetPKVarDeclear(), sd.StructName, sd.JoinMysqlFieldNameList(","),
			sd.StructName, sd.GetWherePosStr2(), keySum, sd.GetWherePosValueWithoutThisPrefix(),
			sd.GetDBStorageName(), sd.GetWherePosStr(), sd.GetWherePosValueWithoutThisPrefix(),
			sd.StructName)
	return s

}
func (sd StructDescription) GenerateDBAdd() string {
	pkList := sd.GetPKFieldList()
	var keySum string = "nil"
	if len(pkList) > 0 && isTypeKeySum(pkList[0].FieldGoType) {
		keySum = "e." + pkList[0].FieldName
	}

	s :=
		fmt.Sprintf(
			`func (this *%s) Add(e *%s, now time.Time) bool {
	sql := e.GetInsertSql()
	err := this.DatabaseTemplate.Exec(%s, sql)
	e.ClearFlag()
     if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.Add %%v %%v %%s",*e,err,sql)
        }
     }
	return err == nil
}
`, sd.GetDBStorageName(), sd.StructName, keySum, sd.GetDBStorageName())
	return s

}
func (sd StructDescription) GenerateDBMultiAdd() string {
	pkList := sd.GetPKFieldList()
	var keySum string = "nil"
	if len(pkList) > 2 || len(pkList) == 0 {
		return ""
	}
	if isTypeKeySum(pkList[0].FieldGoType) {
		keySum = pkList[0].FieldName
	}

	if len(pkList) == 2 {

		s :=
			fmt.Sprintf(
				`func (this *%s) MultiAdd(%s %s,eMap %sMap, now time.Time) bool {
	sql := eMap.GetInsertSql()
	err := this.DatabaseTemplate.Exec(%s, sql)
     if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiAdd %%v %%v %%s",%s,err,sql)
        }
        return false
     }
	for k, e := range eMap {
	    e.ClearFlag()
        eMap[k]=e
	}

	return true
}
`, sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType, sd.StructName,
				keySum,
				sd.GetDBStorageName(), pkList[0].FieldName)
		return s

	} else if len(pkList) == 1 {
		s :=
			fmt.Sprintf(
				`func (this *%s) MultiAdd(eMap %sMap, now time.Time) bool {
	sql := eMap.GetInsertSql()
    // maybe bug when sharding ,maybe add at all sharding db
	err := this.DatabaseTemplate.Exec(nil, sql) // pass nil as sharding sum
     if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiAdd %%v  %%s",err,sql)
        }
        return false
     }
	for k, e := range eMap {
	    e.ClearFlag()
        eMap[k]=e
	}

	return true
}
`, sd.GetDBStorageName(), sd.StructName,
				sd.GetDBStorageName())
		return s

	}

	return ""
}
func (sd StructDescription) GenerateDBSet() string {
	pkList := sd.GetPKFieldList()
	if len(pkList) == 0 {
		return ""
	}

	var keySum string = "nil"
	if isTypeKeySum(pkList[0].FieldGoType) {
		keySum = "e." + pkList[0].FieldName
	}

	s :=
		fmt.Sprintf(
			`func (this *%s) Set(e *%s, now time.Time) (succ bool) {
	if !e.IsFlagDirty() {
		return true
	}
	if e.IsFlagNew() {
		return this.Add(e, now)
	}
	sql := e.GetUpdateSql()
	if sql == "" {
		return true
	}
	err := this.DatabaseTemplate.Exec(%s, sql)
	succ = (err == nil)
	if succ {
		e.ClearFlag()
	}else{
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.Set %%v %%v",*e,err)
        }
   }
	return
}
`, sd.GetDBStorageName(), sd.StructName, keySum, sd.GetDBStorageName())
	return s
}
func (sd StructDescription) GenerateDBGetAll() string {

	s :=
		fmt.Sprintf(
			`func (this *%s) GetAll() (eList %sList, ok bool) {
	sql := "select %s from %s"
	var obj interface{}
	var err error
	var arradd []interface{}
	var e %s
	arradd, err = this.DatabaseTemplate.QueryArray(nil, sql, this.mapRow)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.GetAll %%v",err)
        }
		return
	}
	for _, obj = range arradd {
		e = obj.(%s)
		eList = append(eList, e)
	}
	ok = true
	return
}
`,

			sd.GetDBStorageName(), sd.StructName,
			sd.JoinMysqlFieldNameList(","), sd.GetMysqlTableName(),
			sd.StructName,
			sd.GetDBStorageName(),

			sd.StructName,
		)
	return s

}
func (sd StructDescription) GenerateDBMultGet() string {
	var multKey string
	pkList := sd.GetPKFieldList()
	if len(pkList) > 2 || len(pkList) == 0 {
		return ""
	}

	var keySum string = "nil"
	if len(pkList) == 2 {
		if isTypeKeySum(pkList[0].FieldGoType) {
			keySum = pkList[0].FieldName
			multKey = fmt.Sprintf("%sList", pkList[1].FieldGoType)
		} else {
			multKey = fmt.Sprintf("[]%s", pkList[1].FieldGoType)
		}

		s :=
			fmt.Sprintf(
				`func (this *%s) MultiGet(%s %s,%ss %s, now time.Time) (eMap %s, ok bool) {
	eMap = make(%s)
	sql := fmt.Sprintf("select %s from %s where %s=%s and %s in (%%s)", %s,%s)
	var obj interface{}
	var err error
	var arradd []interface{}
	var e %s
	arradd, err = this.DatabaseTemplate.QueryArray(%s, sql, this.mapRow)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiGet %%v %%v ,%%v",%s,%ss,err)
        }
		return
	}
	eMap = make(%s)
	for _, obj = range arradd {
		e = obj.(%s)
		eMap[e.%s] = e
	}
	ok = true
	return
}
`,
				sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType, pkList[1].FieldName, multKey, sd.GetSuggestMapName(),
				sd.GetSuggestMapName(),
				sd.JoinMysqlFieldNameList(","), sd.GetMysqlTableName(), pkList[0].GetMysqlFieldName(), pkList[0].GetFieldPosStr(), pkList[1].GetMysqlFieldName(), pkList[0].GetFieldPosValueWithoutPrefix(), pkList[1].GetFieldListPosValue(),
				sd.StructName,
				keySum,
				sd.GetDBStorageName(), pkList[0].FieldName, pkList[1].FieldName,

				sd.GetSuggestMapName(),
				sd.StructName,
				pkList[1].FieldName,
			)
		return s

	} else if len(pkList) == 1 {
		if isTypeKeySum(pkList[0].FieldGoType) {
			keySum = pkList[0].FieldName + "s"
			multKey = fmt.Sprintf("%sList", pkList[0].FieldGoType)
		} else {
			multKey = fmt.Sprintf("[]%s", pkList[0].FieldGoType)
		}

		s :=
			fmt.Sprintf(
				`func (this *%s) MultiGet(%ss %s, now time.Time) (eMap %s, ok bool) {
	eMap = make(%s)
	sql := fmt.Sprintf("select %s from %s where %s in (%%s)",%s)
	var obj interface{}
	var err error
	var arradd []interface{}
	var e %s
	arradd, err = this.DatabaseTemplate.QueryArray(%s, sql, this.mapRow)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiGet %%v %%v",%ss,err)
        }
		return
	}
	eMap = make(%s)
	for _, obj = range arradd {
		e = obj.(%s)
		eMap[e.%s] = e
	}
	ok = true
	return
}
`,
				sd.GetDBStorageName(), pkList[0].FieldName, multKey, sd.GetSuggestMapName(),
				sd.GetSuggestMapName(),
				sd.JoinMysqlFieldNameList(","), sd.GetMysqlTableName(), pkList[0].GetMysqlFieldName(), pkList[0].GetFieldListPosValue(),
				sd.StructName,
				keySum,
				sd.GetDBStorageName(), pkList[0].FieldName,
				sd.GetSuggestMapName(),
				sd.StructName,
				pkList[0].FieldName,
			)
		return s

	}
	return ""
}

func (sd StructDescription) GenerateDBMultUpdate() string {
	pkList := sd.GetPKFieldList()
	if len(pkList) == 2 {
		s := fmt.Sprintf(
			`func (this *%s) MultiUpdate(%s %s,eMap %s, now time.Time) (ok bool) {
	for k, e := range eMap {
		this.Set(&e, now)
		eMap[k] = e
	}
	return true
}
`, sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType,
			sd.GetSuggestMapName())
		return s

	} else if len(pkList) == 1 {

		s := fmt.Sprintf(
			`func (this *%s) MultiUpdate(eMap %s, now time.Time) (ok bool) {
	for k, e := range eMap {
		this.Set(&e, now)
		eMap[k] = e
	}
	return true
}
`, sd.GetDBStorageName(),
			sd.GetSuggestMapName())
		return s
	}
	return ""
}
func (sd StructDescription) GenerateDBDelete() string {
	pkList := sd.GetPKFieldList()
	if len(pkList) == 0 {
		return ""
	}

	var keySum string = "nil"
	if isTypeKeySum(pkList[0].FieldGoType) {
		keySum = pkList[0].FieldName
	}

	s :=
		fmt.Sprintf(
			`func (this *%s) Delete(%s) (ok bool) {
	sql := "delete from %s where %s "
	var err error
	err = this.DatabaseTemplate.Exec(%s, sql,%s)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.Delete %s %%v",%s,err)
        }
		return
	}
	ok = true
	return
}
`, sd.GetDBStorageName(), sd.GetPKVarDeclear(),
			sd.GetMysqlTableName(), sd.GetWherePosStr2(),
			keySum, sd.GetWherePosValueWithoutThisPrefix(),

			sd.GetDBStorageName(), sd.GetWherePosStr(), sd.GetWherePosValueWithoutThisPrefix(),
		)
	return s

}
func (sd StructDescription) GenerateDBMultDelete() string {
	var multKey string
	pkList := sd.GetPKFieldList()
	if len(pkList) > 2 {
		return ""
	}
	if len(pkList) == 0 {
		return ""
	}

	var keySum string = "nil"
	if len(pkList) == 2 {
		if isTypeKeySum(pkList[0].FieldGoType) {
			keySum = pkList[0].FieldName
			multKey = fmt.Sprintf("%sList", pkList[1].FieldGoType)
		} else {
			multKey = fmt.Sprintf("[]%s", pkList[1].FieldGoType)
		}

		s :=
			fmt.Sprintf(
				`func (this *%s) MultiDelete(%s %s,%ss %s) (ok bool) {
	sql := fmt.Sprintf("delete from %s where %s=%s and %s in (%%s)", %s,%s)
	var err error
	 err = this.DatabaseTemplate.Exec(%s, sql)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiDelete %%v,%%v %%v",%s,%ss,err)
        }
		return
	}
	ok = true
	return
}
`,
				sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType, pkList[1].FieldName, multKey,
				sd.GetMysqlTableName(), pkList[0].GetMysqlFieldName(), pkList[0].GetFieldPosStr(), pkList[1].GetMysqlFieldName(), pkList[0].GetFieldPosValueWithoutPrefix(), pkList[1].GetFieldListPosValue(),
				keySum,
				sd.GetDBStorageName(), pkList[0].FieldName, pkList[1].FieldName,
			)
		return s

	} else if len(pkList) == 1 {
		if isTypeKeySum(pkList[0].FieldGoType) {
			keySum = pkList[0].FieldName + "s"
			multKey = fmt.Sprintf("%sList", pkList[0].FieldGoType)
		} else {
			multKey = fmt.Sprintf("[]%s", pkList[0].FieldGoType)
		}

		s :=
			fmt.Sprintf(
				`func (this *%s) MultiDelete(%ss %s) (ok bool) {
	sql := fmt.Sprintf("delete from %s where %s in (%%s)",%s)
	var err error
	err = this.DatabaseTemplate.Exec(%s, sql)
	if err != nil {
        if this.log != nil {
            this.log.Errorf("[DB.ERR]%s.MultiDelete %%v %%v",%ss,err)
        }
		return
	}
	ok = true
	return
}
`,
				sd.GetDBStorageName(), pkList[0].FieldName, multKey,
				sd.GetMysqlTableName(), pkList[0].GetMysqlFieldName(), pkList[0].GetFieldListPosValue(),
				keySum,
				sd.GetDBStorageName(), pkList[0].FieldName,
			)
		return s

	}
	return ""
}
func (sd StructDescription) GenerateDBGetIdList() string {
	var multKey string
	pkList := sd.GetPKFieldList()
	if len(pkList) != 2 {
		return ""
	}

	var keySum string = "nil"
	if isTypeKeySum(pkList[0].FieldGoType) {
		keySum = pkList[0].FieldName
		multKey = fmt.Sprintf("%sList", pkList[1].FieldGoType)
	} else {
		multKey = fmt.Sprintf("[]%s", pkList[1].FieldGoType)
	}

	s :=
		fmt.Sprintf(
			`
func (this *%s) GetIdListByPK1(%s %s, now time.Time) (idList %s, ok bool) {
	sqlstr := fmt.Sprintf("select %s from %s where %s = ?")
	var rows []interface{}
	mapRow := func(resultSet *sql.Rows) (interface{}, error) {
		var %s %s //need fix maybe
		err := resultSet.Scan(&%s)
		if err != nil {
            if this.log != nil {
                this.log.Errorf("[DB.ERR]%s.GetIdListByPK1.mapRow %%v %%v",%s,err)
            }
			return nil, err
		}
		return %s, err
	}
	rows, err := this.DatabaseTemplate.QueryArray(%s, sqlstr, mapRow, %s)
	if err != nil {
		return
	}
	idList = make(%s, len(rows))
	for i, obj := range rows {
		idList[i] = (obj.(%s))
	}
	return idList, true
}

`,
			sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType, multKey,
			pkList[1].GetMysqlFieldName(), sd.GetMysqlTableName(), pkList[0].GetMysqlFieldName(),
			pkList[1].GetMysqlFieldName(), pkList[1].FieldGoType,
			pkList[1].GetMysqlFieldName(),
			sd.GetDBStorageName(), pkList[0].FieldName,
			pkList[1].GetMysqlFieldName(),
			keySum, pkList[0].FieldName,
			multKey,
			pkList[1].FieldGoType,
		)
	return s

}
func (sd StructDescription) GenerateDBsetIdList() string {
	var multKey string
	pkList := sd.GetPKFieldList()
	if len(pkList) != 2 {
		return ""
	}

	if isTypeKeySum(pkList[0].FieldGoType) {
		multKey = fmt.Sprintf("%sList", pkList[1].FieldGoType)
	} else {
		multKey = fmt.Sprintf("[]%s", pkList[1].FieldGoType)
	}

	s :=
		fmt.Sprintf(
			`
func (this *%s) SetIdListByPK1(%s %s, idList *%s,now time.Time) (ok bool) {
	return true
}
`, sd.GetDBStorageName(), pkList[0].FieldName, pkList[0].FieldGoType, multKey)
	return s

}
