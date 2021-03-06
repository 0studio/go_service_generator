package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func (sd StructDescription) GenerateService(property Property, srcDir string) {
	pkList := sd.GetPKFieldList()
	if len(pkList) > 2 || len(pkList) == 0 {
		return
	}

	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("service_%s_template.go", strings.ToLower(sd.StructName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputF.Close()
	s := strings.Replace(SERVICE_TEMPLATE, "__package__", property.PackageName, -1)
	s = strings.Replace(s, "__Entity__", sd.StructName, -1)
	s = strings.Replace(s, "__LowercaseEntity__", LowerCaseFirstChar(sd.StructName), -1)

	getDefaultFunDeclare, getDefaultFunImpl := sd.generateServiceGetDefault()

	s = strings.Replace(s, "__GetDefaultDeclard__", getDefaultFunDeclare, -1)
	s = strings.Replace(s, "__GetDefaultImplments__", getDefaultFunImpl, -1)

	getAllFunDeclare, getAllImpl := sd.generateServiceGetAll()
	s = strings.Replace(s, "__GetAllDeclard__", getAllFunDeclare, -1)
	s = strings.Replace(s, "__GetAllImplments__", getAllImpl, -1)
	clearLocalDeclare, clearLocalImpl := sd.generateServiceClearLocal()
	s = strings.Replace(s, "__ClearLocalDeclard__", clearLocalDeclare, -1)
	s = strings.Replace(s, "__ClearLocalImplments__", clearLocalImpl, -1)

	formatSrc, err := format.Source([]byte(s))
	if err == nil {
		outputF.WriteString(string(formatSrc))
	} else {
		outputF.WriteString(s)
	}

}
func (sd StructDescription) generateServiceGetDefault() (declare string, implements string) {
	pkList := sd.GetPKFieldList()
	if len(pkList) > 2 {
		return
	}
	declare = fmt.Sprintf("    GetDefault(%s, now time.Time) (e %s, ok bool)", sd.GetPKVarDeclear(), sd.StructName)
	implements =
		fmt.Sprintf(`func (impl *%sServiceImpl) GetDefault(%s, now time.Time) (e %s, ok bool) {
	e, ok = impl.Get(%s, now)
	if !ok {
		e = New%s(%s)
		ok = impl.Add(&e, now)
		return
	}
	return
}
`, sd.StructName, sd.GetPKVarDeclear(), sd.StructName,
			sd.GetWherePosValueWithoutThisPrefix(),
			sd.StructName, sd.GetWherePosValueWithoutThisPrefix())

	return
}

func (sd StructDescription) generateServiceGetAll() (declare string, implements string) {
	pkList := sd.GetPKFieldList()
	if len(pkList) != 2 {
		return
	}
	declare = fmt.Sprintf("    GetAll(%s %s, now time.Time) (eMap %s, ok bool)",
		LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType, sd.GetSuggestMapName())
	implements =
		fmt.Sprintf(`func (impl *%sServiceImpl) GetAll(%s %s, now time.Time) (eMap %s, ok bool) {
     // maybe bug if you do not update idlist  when you add ,delete
     // maybe the cachedIdList is dirty by some reason
     // so you should use this carefully
	idList, ok := impl.GetIdListByPK1(%s, now)
	if !ok {
		return
	}
	return impl.MultiGet(%s, idList, now)
}
`, sd.StructName, LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType, sd.GetSuggestMapName(),
			LowerCaseFirstChar(pkList[0].FieldName),
			LowerCaseFirstChar(pkList[0].FieldName))

	return
}

func (sd StructDescription) generateServiceClearLocal() (declare string, implements string) {
	pkList := sd.GetPKFieldList()
	if len(pkList) == 1 {
		declare = fmt.Sprintf("    ClearLocal(%s %s, now time.Time)",
			LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType)
		implements =
			fmt.Sprintf(`func (impl *%sServiceImpl) ClearLocal(%s %s, now time.Time) {
    if impl.lruStorage != nil {
		impl.lruStorage.Delete(%s)
    }
}
`, sd.StructName, LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType,
				LowerCaseFirstChar(pkList[0].FieldName))

		return

	} else if len(pkList) == 2 {
		declare = fmt.Sprintf("    ClearLocal(%s %s, now time.Time)",
			LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType)
		implements =
			fmt.Sprintf(`func (impl *%sServiceImpl) ClearLocal(%s %s, now time.Time) {
	idList, ok := impl.GetIdListByPK1(%s, now)
	if !ok {
		return
	}
    if impl.lruStorage != nil {
		ok = impl.lruStorage.MultiDelete(%s, idList)
		impl.lruStorage.DeleteIdListByPK1(%s)
    }

}
`, sd.StructName, LowerCaseFirstChar(pkList[0].FieldName), pkList[0].FieldGoType,
				LowerCaseFirstChar(pkList[0].FieldName),
				LowerCaseFirstChar(pkList[0].FieldName),
				LowerCaseFirstChar(pkList[0].FieldName))

		return

	}
	return
}

const (
	SERVICE_TEMPLATE = `// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package __package__

import (
	"github.com/0studio/databasetemplate"
	"github.com/0studio/logger"
	"github.com/dropbox/godropbox/memcache"
	"github.com/0studio/goutils"
    key "github.com/0studio/storage_key"
	"time"
)
var ___importTimeSrv__Entity__ time.Time
var ___importKeySrv__Entity__ key.KeyUint64
var ___importGoutilsSrv__Entity__ goutils.Int32List


type __Entity__Service interface {
	__Entity__Storage

	// 这个接口是让你来扩展  __Entity__Service 接口的，
	// 本文件是工具生成， 尽量不要编辑此文件，
	// 当自动生成的文件没法满足你需求的时候
	// 在另一个文件里实现新的接口， 保持这个文件不变
    __Entity__ServiceOther
__GetDefaultDeclard__
__GetAllDeclard__
__ClearLocalDeclard__

}

var __LowercaseEntity__Service *__Entity__ServiceImpl

func Get__Entity__Service() __Entity__Service {
	return __LowercaseEntity__Service
}

// db/memecache/lrucache
// log can be nil
func New__Entity__Service(dt databasetemplate.DatabaseTemplate, mcClient memcache.Client, log logger.Logger, createTable bool) __Entity__Service {
	lruStorage := NewLRUCache__Entity__Storage(LRU_Cache_Sharding_Cnt___Entity__, LRU_Cache_Size___Entity__)
	mcStorage := NewMC__Entity__Storage(mcClient, Memcache_Expired_Seconds___Entity__, Memcache_Prefix___Entity__)
	dbStorage := NewDB__Entity__Storage(dt, log, createTable)
	__LowercaseEntity__Service = &__Entity__ServiceImpl{
		lruStorage:   lruStorage,
     	dbStorage:dbStorage,
     	mcStorage:mcStorage,
		lruMCStorage: NewStorageProxy__Entity__(lruStorage, mcStorage),
		mcDBStorage:  NewStorageProxy__Entity__(mcStorage, dbStorage),
        log:               log,
	}

	__LowercaseEntity__Service.__Entity__Storage = NewStorageProxy__Entity__(__LowercaseEntity__Service.lruMCStorage, dbStorage)
	return __LowercaseEntity__Service
}

// lrucache and db
// log can be nil
func New__Entity__ServiceCacheAndDB(dt databasetemplate.DatabaseTemplate, log logger.Logger, createTable bool) __Entity__Service {
	lruStorage := NewLRUCache__Entity__Storage(LRU_Cache_Sharding_Cnt___Entity__, LRU_Cache_Size___Entity__)
	dbStorage := NewDB__Entity__Storage(dt, log, createTable)
	__LowercaseEntity__Service = &__Entity__ServiceImpl{
		lruStorage:        lruStorage,
     	dbStorage:dbStorage,
		lruMCStorage:      lruStorage,
		mcDBStorage:       dbStorage,
		__Entity__Storage: NewStorageProxy__Entity__(lruStorage, dbStorage),
        log:               log,
	}

	return __LowercaseEntity__Service

}
// memcache and db
// log can be nil
func New__Entity__ServiceMCAndDB(dt databasetemplate.DatabaseTemplate, mcClient memcache.Client, log logger.Logger, createTable bool) __Entity__Service {
	mcStorage := NewMC__Entity__Storage(mcClient, Memcache_Expired_Seconds___Entity__, Memcache_Prefix___Entity__)
	dbStorage := NewDB__Entity__Storage(dt, log, createTable)
	__LowercaseEntity__Service = &__Entity__ServiceImpl{
     	dbStorage:dbStorage,
     	mcStorage:mcStorage,
		lruMCStorage: mcStorage,
		mcDBStorage:  NewStorageProxy__Entity__(mcStorage, dbStorage),
		__Entity__Storage:NewStorageProxy__Entity__(mcStorage, dbStorage),
        log:               log,
	}

	return __LowercaseEntity__Service
}



type __Entity__ServiceImpl struct {
	__Entity__Storage
	lruStorage   *LRUCache__Entity__Storage // only lrucache
	dbStorage    *DB__Entity__Storage         // db
	mcStorage    MC__Entity__Storage         // mc
	mcDBStorage  __Entity__Storage         // mc and db proxy
	lruMCStorage __Entity__Storage         // lruCache and mc proxy
    log          logger.Logger
}

func (impl *__Entity__ServiceImpl) setOutside(e *__Entity__, now time.Time) {
	// 不改本地缓存， 只改 进程外部的 memcache及db
	// 主要用于当 lrucache 里的内容被各种原因删除时的回调处理
	// func (l __Entity__) OnPurge(why lru.PurgeReason)
	if impl.mcDBStorage != nil {
		impl.mcDBStorage.Set(e, now)
	}
}
__GetDefaultImplments__
__GetAllImplments__
__ClearLocalImplments__

`
)
