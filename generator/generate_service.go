package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (sd StructDescription) GenerateService(property Property, srcDir string) {
	pkList := sd.GetPKFieldList()
	if len(pkList) > 2 {
		return
	}

	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("service_%s_stub.go", strings.ToLower(sd.StructName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputF.Close()
	s := strings.Replace(SERVICE_TEMPLATE, "__package__", property.PackageName, -1)
	s = strings.Replace(s, "__Entity__", sd.StructName, -1)
	s = strings.Replace(s, "__LowercaseEntity__", LowerCaseFirstChar(sd.StructName), -1)
	outputF.WriteString(s)

}

const (
	SERVICE_TEMPLATE = `// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package __package__

import (
	"github.com/0studio/databasetemplate"
	"github.com/0studio/logger"
	"github.com/dropbox/godropbox/memcache"
	"time"
)

type __Entity__Service interface {
	__Entity__Storage

	// 这个接口是让你来扩展  __Entity__Service 接口的，
	// 本文件是工具生成， 尽量不要编辑此文件，
	// 当自动生成的文件没法满足你需求的时候
	// 在另一个文件里实现新的接口， 保持这个文件不变
	__Entity__ServiceOther
}

var __LowercaseEntity__Service *__Entity__ServiceImpl

func Get__Entity__Service() __Entity__Service {
	return __LowercaseEntity__Service
}

// db/memecache/lrucache
func New__Entity__Service(dt databasetemplate.DatabaseTemplate, createTable bool, mcClient memcache.Client, log logger.Logger) __Entity__Service {
	lruStorage := NewLRULocal__Entity__Storage(LRU_Cache_Sharding_Cnt, LRU_Cache_Size)
	mcStorage := NewMC__Entity__Storage(mcClient, Memcache_Expired_Seconds, Memcache_Prefix)
	dbStorage := NewDB__Entity__Storage(dt, log, createTable)
	__LowercaseEntity__Service = &__Entity__ServiceImpl{
		lruStorage:   lruStorage,
		lruMCStorage: NewStorageProxy(lruStorage, mcStorage),
		mcDBStorage:  NewStorageProxy(mcStorage, dbStorage),
	}

	__LowercaseEntity__Service.__Entity__Storage = NewStorageProxy(__LowercaseEntity__Service.lruMCStorage, dbStorage)
	return __LowercaseEntity__Service
}

// lrucache and db
func New__Entity__ServiceLocalAndDB(dt databasetemplate.DatabaseTemplate, log logger.Logger, createTable bool) __Entity__Service {
	lruStorage := NewLRULocal__Entity__Storage(LRU_Cache_Sharding_Cnt, LRU_Cache_Size)
	dbStorage := NewDB__Entity__Storage(dt, log, createTable)
	__LowercaseEntity__Service = &__Entity__ServiceImpl{
		lruStorage:        lruStorage,
		lruMCStorage:      lruStorage,
		mcDBStorage:       dbStorage,
		__Entity__Storage: NewStorageProxy(lruStorage, dbStorage),
	}

	return __LowercaseEntity__Service

}

type __Entity__ServiceImpl struct {
	__Entity__Storage
	lruStorage   LRULocal__Entity__Storage // only lrucache
	mcDBStorage  __Entity__Storage         // mc and db proxy
	lruMCStorage __Entity__Storage         // lruCache and mc proxy
}

func (impl *__Entity__ServiceImpl) setOutside(e *__Entity__, now time.Time) {
	// 不改本地缓存， 只改 进程外部的 memcache及db
	// 主要用于当 lrucache 里的内容被各种原因删除时的回调处理
	// func (l __Entity__) OnPurge(why lru.PurgeReason)
	if impl.mcDBStorage != nil {
		impl.mcDBStorage.Set(e, now)
	}
}
`
)
