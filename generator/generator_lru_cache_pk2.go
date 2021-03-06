package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func (sd StructDescription) generateLRUCache2PK(pk1Field FieldDescriptoin, pk2Field FieldDescriptoin, property Property, srcDir string) bool {
	pk1LRUCacheType := sd.getLRUCacheType(pk1Field)
	if pk1LRUCacheType == "" {
		return false
	}
	pk2LRUCacheType := sd.getLRUCacheType(pk2Field)
	if pk2LRUCacheType == "" {
		return false
	}

	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("storage_%s_lru_cache_template.go", strings.ToLower(sd.StructName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer outputF.Close()

	s := strings.Replace(LRUCacheTemplatePK2, "__package__", property.PackageName, -1)
	s = strings.Replace(s, "__Entity__", sd.StructName, -1)
	s = strings.Replace(s, "__PK1Type__", pk1Field.FieldGoType, -1)
	s = strings.Replace(s, "__PK1FieldName__", pk1Field.FieldName, -1)
	s = strings.Replace(s, "__PK2Type__", pk2Field.FieldGoType, -1)
	s = strings.Replace(s, "__PK2FieldName__", pk2Field.FieldName, -1)
	if strings.Contains(pk1LRUCacheType, "Shard") {
		s = strings.Replace(s, "__NewLRUCacheType__", "New"+pk1LRUCacheType+"(shardingCnt,size)", -1)
	} else {
		s = strings.Replace(s, "__NewLRUCacheType__", "New"+pk1LRUCacheType+"(size)", -1)
	}
	s = strings.Replace(s, "__LRUCacheType__", pk1LRUCacheType, -1)

	var pk2TypeList string
	if isTypeKeySum(pk2Field.FieldGoType) {
		pk2TypeList = fmt.Sprintf("%sList", pk2Field.FieldGoType)
	} else {
		pk2TypeList = fmt.Sprintf("[]%s", pk2Field.FieldGoType)
	}

	s = strings.Replace(s, "__PK2TypeList__", pk2TypeList, -1)

	formatSrc, err := format.Source([]byte(s))
	if err == nil {
		outputF.WriteString(string(formatSrc))
	} else {
		outputF.WriteString(s)
	}

	return true
}

const (
	LRUCacheTemplatePK2 = `// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package __package__

import (
	"github.com/0studio/goutils"
	"github.com/0studio/lru"
	key "github.com/0studio/storage_key"
	"time"
)

var __importKeyL__Entity__ key.KeyUint64
var __importGoutilsL__Entity__ goutils.Int32List

type LRUCache__Entity__Storage struct {
	cache     *lru.__LRUCacheType__
	cacheList *lru.__LRUCacheType__
}

func NewLRUCache__Entity__Storage(shardingCnt int, size int64) (local__Entity__Storage *LRUCache__Entity__Storage) {
	local__Entity__Storage = &LRUCache__Entity__Storage{
		cache: lru.__NewLRUCacheType__,
	    cacheList: lru.__NewLRUCacheType__,
	}

	return
}

func (m *LRUCache__Entity__Storage) getMap(__PK1FieldName__ __PK1Type__) (sm __Entity__Map) { // nil or map
	cacheObj, ok := m.cache.Get(__PK1FieldName__)
	if !ok {
		return
	}
	sm = cacheObj.(__Entity__Map)
	return

}

func (m *LRUCache__Entity__Storage) Get(__PK1FieldName__ __PK1Type__, __PK2FieldName__ __PK2Type__, now time.Time) (e __Entity__, ok bool) {
	sm := m.getMap(__PK1FieldName__)
	if sm == nil {
		return
	}
	e, ok = sm[__PK2FieldName__]
	return
}
func (m *LRUCache__Entity__Storage) set(__PK1FieldName__ __PK1Type__, eMap __Entity__Map) {
	m.cache.Set(__PK1FieldName__, eMap)
}
func (m *LRUCache__Entity__Storage) Set(e *__Entity__, now time.Time) (ok bool) {
	sm := m.getMap(e.__PK1FieldName__)
	if sm == nil {
		sm = make(__Entity__Map)
		sm[e.__PK2FieldName__] = *e
		m.cache.Set(e.__PK1FieldName__, sm)
		return true
	}
	sm[e.__PK2FieldName__] = *e
	return true
}
func (m *LRUCache__Entity__Storage) Add(e *__Entity__, now time.Time) (ok bool) {
    ok = m.add(e,now)
    if ok &&LRU_CACHE_USE_LIST___Entity__{
        list,listOk := m.GetIdListByPK1(e.__PK1FieldName__,now)
        if listOk {
             var isInList bool
             for _, id:= range list{
                 if id==e.__PK2FieldName__{
                      isInList = true
                      break
                 }
             }
             if !isInList{
                  list = append(list,e.__PK2FieldName__)
                  m.SetIdListByPK1(e.__PK1FieldName__,&list,now)
             }
        }
    }
    return
}
func (m *LRUCache__Entity__Storage) add(e *__Entity__, now time.Time) (ok bool) {
	sm := m.getMap(e.__PK1FieldName__)
	if sm == nil {
		sm = make(__Entity__Map)
		sm[e.__PK2FieldName__] = *e
		m.cache.Set(e.__PK1FieldName__, sm)
		return true
	}
	sm[e.__PK2FieldName__] = *e

	return true
}
func (m *LRUCache__Entity__Storage) MultiGet(__PK1FieldName__ __PK1Type__, keys __PK2TypeList__, now time.Time) (eMap __Entity__Map, ok bool) {
	eMap = make(__Entity__Map)
	ok = true
	sm := m.getMap(__PK1FieldName__)
	for _, k := range keys {
		st, ok := sm[k]
		if ok {
			eMap[st.__PK2FieldName__] = st
		}

	}

	return
}
func (m *LRUCache__Entity__Storage) MultiUpdate(__PK1FieldName__ __PK1Type__, eMap __Entity__Map, now time.Time) (ok bool) {
	for _, e := range eMap {
		m.Set(&e, now)
	}
	return true
}
func (m *LRUCache__Entity__Storage) MultiAdd(__PK1FieldName__ __PK1Type__, eMap __Entity__Map, now time.Time) (ok bool) {
	for _, e := range eMap {
		m.Add(&e, now)
	}
	return true

}
func (m *LRUCache__Entity__Storage) del(__PK1FieldName__ __PK1Type__, now time.Time) (ok bool) {
	m.cache.Delete(__PK1FieldName__)
	return true
}
func (m *LRUCache__Entity__Storage) Delete(__PK1FieldName__ __PK1Type__, __PK2FieldName__ __PK2Type__) (ok bool) {
	sm := m.getMap(__PK1FieldName__)
	if sm == nil {
		return true
	}
	delete(sm, __PK2FieldName__)
    now:=time.Now()
    if LRU_CACHE_USE_LIST___Entity__{
		list,listOk := m.GetIdListByPK1(__PK1FieldName__,now)
		if listOk {
				newList:=make(__PK2TypeList__,0,2*len(list))
				for _, tmpId:= range list{
					if tmpId!=__PK2FieldName__{
						newList = append(newList,tmpId)
					}
				}
				m.SetIdListByPK1(__PK1FieldName__,&newList,now)
		}
    }
	return true
}
func (m *LRUCache__Entity__Storage) MultiDelete(__PK1FieldName__ __PK1Type__, keys __PK2TypeList__) (ok bool) {
	for _, __PK2FieldName__ := range keys {
		m.Delete(__PK1FieldName__, __PK2FieldName__)
	}
	return true
}

func (m *LRUCache__Entity__Storage) SetIdListByPK1(__PK1FieldName__ __PK1Type__, idList *__PK2TypeList__, now time.Time) bool {
    if !LRU_CACHE_USE_LIST___Entity__{
        return true
    }
	m.cacheList.Set(__PK1FieldName__, *idList)
	return true
}

func (m *LRUCache__Entity__Storage) GetIdListByPK1(__PK1FieldName__ __PK1Type__, now time.Time) (list __PK2TypeList__, ok bool) {
    if !LRU_CACHE_USE_LIST___Entity__{
        return
    }
	cacheObj, ok := m.cacheList.Get(__PK1FieldName__)
	if !ok {
		return
	}
	list = cacheObj.(__PK2TypeList__)
	return
}
func (m *LRUCache__Entity__Storage) DeleteIdListByPK1(__PK1FieldName__ __PK1Type__) (ok bool) {
    if !LRU_CACHE_USE_LIST___Entity__{
        return
    }
	m.cacheList.Delete(__PK1FieldName__)
	return true
}
func (m *LRUCache__Entity__Storage) GetAllDirty(__PK1FieldName__ __PK1Type__, now time.Time) (eMap __Entity__Map) {
	sm := m.getMap(__PK1FieldName__)
	if sm == nil {
		return
	}

	for key, s := range sm {
		if s.IsFlagDirty() {
			if eMap == nil {
				eMap = make(__Entity__Map)
			}
			eMap[key] = s
		}
	}
	return
}
`
)
