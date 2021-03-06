package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func GenerateUtils(property Property, srcDir string) {
	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("utils_%s_template.go", strings.ToLower(property.PackageName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputF.Close()
	s := fmt.Sprintf(
		`// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package %s

import (
	"github.com/0studio/logger"
	key "github.com/0studio/storage_key"
	"runtime"
	"strconv"
	"strings"
    "time"
)


func bool2int(b bool) int {
    if b {
        return 1
    } else {
        return 0
    }
}
func int2bool(i int)bool {
    if i == 0 {
        return false
    } else {
        return true
    }
}
func formatTime(t time.Time)string {
    if t.IsZero() {
        return "0"
    } else {
        return t.Format("20060102150405")
    }
}
func formatTimeUnix(t time.Time)int64 {
    if t.IsZero() {
        return 0
    } else {
        return t.Unix()
    }
}
func newTime(unix int64) (t time.Time){
    if unix==0 {
        return
    } else {
        return time.Unix(unix,0)
    }
}
func stringListJoin(sList []string, separator string) string {
	return strings.Join(sList, separator)
}
func stringSplit(s string, separator string) []string {
	return strings.Split(s, separator)
}


func uint64ListJoin(intLi []uint64, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatUint(v, 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func int64ListJoin(intLi []int64, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatInt(v, 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func uint32ListJoin(intLi []uint32, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatUint(uint64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func int32ListJoin(intLi []int32, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatInt(int64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func intListJoin(intLi []int, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.Itoa(v)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func uint16ListJoin(intLi []uint16, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatUint(uint64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func int16ListJoin(intLi []int16, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatInt(int64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func uint8ListJoin(intLi []uint8, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatUint(uint64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func int8ListJoin(intLi []int8, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.FormatInt(int64(v), 10)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func str2uint64List(str string, separator string) (uintList []uint64) {
	strList := strings.Split(str, separator)
	uintList = make([]uint64, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseUint(token, 10, 0)
		uintList[i] = v
	}
	return
}
func str2uint32List(str string, separator string) (uintList []uint32) {
	strList := strings.Split(str, separator)
	uintList = make([]uint32, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseUint(token, 10, 0)
		uintList[i] = uint32(v)
	}
	return
}
func str2uint16List(str string, separator string) (uintList []uint16) {
	strList := strings.Split(str, separator)
	uintList = make([]uint16, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseUint(token, 10, 0)
		uintList[i] = uint16(v)
	}
	return
}
func str2uint8List(str string, separator string) (uintList []uint8) {
	strList := strings.Split(str, separator)
	uintList = make([]uint8, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseUint(token, 10, 0)
		uintList[i] = uint8(v)
	}
	return
}

func str2int64List(str string, separator string) (intList []int64) {
	strList := strings.Split(str, separator)
	intList = make([]int64, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseInt(token, 10, 0)
		intList[i] = v
	}
	return
}
func str2intList(str string, separator string) (intList []int) {
	strList := strings.Split(str, separator)
	intList = make([]int, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseInt(token, 10, 0)
		intList[i] = int(v)
	}
	return
}
func str2int32List(str string, separator string) (intList []int32) {
	strList := strings.Split(str, separator)
	intList = make([]int32, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseInt(token, 10, 0)
		intList[i] = int32(v)
	}
	return
}
func str2int16List(str string, separator string) (intList []int16) {
	strList := strings.Split(str, separator)
	intList = make([]int16, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseInt(token, 10, 0)
		intList[i] = int16(v)
	}
	return
}
func str2int8List(str string, separator string) (intList []int8) {
	strList := strings.Split(str, separator)
	intList = make([]int8, len(strList))
	for i, token := range strList {
		v, _ := strconv.ParseInt(token, 10, 0)
		intList[i] = int8(v)
	}
	return
}
func sround(s, c string) string {
	return c + s + c
}
func sroundJoin(sList []string, c, sep string) string {
	var tmpList []string = make([]string, len(sList))
	for i, s := range sList {
		tmpList[i] = sround(s, c)
	}
	return strings.Join(tmpList, sep)
}
func sroundJoin2(sList []key.String, c, sep string) string {
    var sList2 []string=make([]string,len(sList))
	for i, s := range sList {
		sList2[i] = sround(string(s), c)
	}
	return strings.Join(sList2, sep)
}
func sroundJoin3(sList []key.KeyString, c, sep string) string {
    var sList2 []string=make([]string,len(sList))
	for i, s := range sList {
		sList2[i] = sround(string(s), c)
	}
	return strings.Join(sList2, sep)
}



func join(list []string, sep string) string {
	return strings.Join(list,sep)
}

func intList2intList(in []int) (out []int) {
    return in
}
func intList2int64List(in []int) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func intList2int32List(in []int) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func intList2int16List(in []int) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func intList2int8List(in []int) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func intList2uint64List(in []int) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func intList2uint32List(in []int) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func intList2uint16List(in []int) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func intList2uint8List(in []int) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func int64List2intList(in []int64) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func int64List2int64List(in []int64) (out []int64) {
    return in
}
func int64List2int32List(in []int64) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func int64List2int16List(in []int64) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func int64List2int8List(in []int64) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func int64List2uint64List(in []int64) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func int64List2uint32List(in []int64) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func int64List2uint16List(in []int64) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func int64List2uint8List(in []int64) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func int32List2intList(in []int32) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func int32List2int64List(in []int32) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func int32List2int32List(in []int32) (out []int32) {
    return in
}
func int32List2int16List(in []int32) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func int32List2int8List(in []int32) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func int32List2uint64List(in []int32) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func int32List2uint32List(in []int32) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func int32List2uint16List(in []int32) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func int32List2uint8List(in []int32) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func int16List2intList(in []int16) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func int16List2int64List(in []int16) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func int16List2int32List(in []int16) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func int16List2int16List(in []int16) (out []int16) {
    return in
}
func int16List2int8List(in []int16) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func int16List2uint64List(in []int16) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func int16List2uint32List(in []int16) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func int16List2uint16List(in []int16) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func int16List2uint8List(in []int16) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func int8List2intList(in []int8) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func int8List2int64List(in []int8) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func int8List2int32List(in []int8) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func int8List2int16List(in []int8) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func int8List2int8List(in []int8) (out []int8) {
    return in
}
func int8List2uint64List(in []int8) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func int8List2uint32List(in []int8) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func int8List2uint16List(in []int8) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func int8List2uint8List(in []int8) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func uint64List2intList(in []uint64) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func uint64List2int64List(in []uint64) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func uint64List2int32List(in []uint64) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func uint64List2int16List(in []uint64) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func uint64List2int8List(in []uint64) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func uint64List2uint64List(in []uint64) (out []uint64) {
    return in
}
func uint64List2uint32List(in []uint64) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func uint64List2uint16List(in []uint64) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func uint64List2uint8List(in []uint64) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func uint32List2intList(in []uint32) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func uint32List2int64List(in []uint32) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func uint32List2int32List(in []uint32) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func uint32List2int16List(in []uint32) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func uint32List2int8List(in []uint32) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func uint32List2uint64List(in []uint32) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func uint32List2uint32List(in []uint32) (out []uint32) {
    return in
}
func uint32List2uint16List(in []uint32) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func uint32List2uint8List(in []uint32) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func uint16List2intList(in []uint16) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func uint16List2int64List(in []uint16) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func uint16List2int32List(in []uint16) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func uint16List2int16List(in []uint16) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func uint16List2int8List(in []uint16) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func uint16List2uint64List(in []uint16) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func uint16List2uint32List(in []uint16) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func uint16List2uint16List(in []uint16) (out []uint16) {
    return in
}
func uint16List2uint8List(in []uint16) (out []uint8) {
	out = make([]uint8, len(in))
	for idx, v := range in {
		out[idx] = uint8(v)
	}
	return
}
func uint8List2intList(in []uint8) (out []int) {
	out = make([]int, len(in))
	for idx, v := range in {
		out[idx] = int(v)
	}
	return
}
func uint8List2int64List(in []uint8) (out []int64) {
	out = make([]int64, len(in))
	for idx, v := range in {
		out[idx] = int64(v)
	}
	return
}
func uint8List2int32List(in []uint8) (out []int32) {
	out = make([]int32, len(in))
	for idx, v := range in {
		out[idx] = int32(v)
	}
	return
}
func uint8List2int16List(in []uint8) (out []int16) {
	out = make([]int16, len(in))
	for idx, v := range in {
		out[idx] = int16(v)
	}
	return
}
func uint8List2int8List(in []uint8) (out []int8) {
	out = make([]int8, len(in))
	for idx, v := range in {
		out[idx] = int8(v)
	}
	return
}
func uint8List2uint64List(in []uint8) (out []uint64) {
	out = make([]uint64, len(in))
	for idx, v := range in {
		out[idx] = uint64(v)
	}
	return
}
func uint8List2uint32List(in []uint8) (out []uint32) {
	out = make([]uint32, len(in))
	for idx, v := range in {
		out[idx] = uint32(v)
	}
	return
}
func uint8List2uint16List(in []uint8) (out []uint16) {
	out = make([]uint16, len(in))
	for idx, v := range in {
		out[idx] = uint16(v)
	}
	return
}
func uint8List2uint8List(in []uint8) (out []uint8) {
    return in
}


func str2uint64(s string) (out uint64) {
	v, _ := strconv.ParseUint(s, 10, 0)
	return uint64(v)
}
func str2uint32(s string) (out uint32) {
	v, _ := strconv.ParseUint(s, 10, 0)
	return uint32(v)
}
func str2uint16(s string) (out uint16) {
	v, _ := strconv.ParseUint(s, 10, 0)
	return uint16(v)
}
func str2uint8(s string) (out uint8) {
	v, _ := strconv.ParseUint(s, 10, 0)
	return uint8(v)
}
func str2int(s string) (out int) {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int(v)
}
func str2int64(s string) (out int64) {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int64(v)
}
func str2int32(s string) (out int32) {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int32(v)
}
func str2int16(s string) (out int16) {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int16(v)
}
func str2int8(s string) (out int8) {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int8(v)
}
func uint642str(v uint64) (s string) {
	s = strconv.FormatUint(uint64(v), 10)
	return
}
func uint322str(v uint32) (s string) {
	s = strconv.FormatUint(uint64(v), 10)
	return
}
func uint162str(v uint16) (s string) {
	s = strconv.FormatUint(uint64(v), 10)
	return
}
func uint82str(v uint8) (s string) {
	s = strconv.FormatUint(uint64(v), 10)
	return
}
func int2str(v int) (s string) {
	s = strconv.FormatInt(int64(v), 10)
	return
}
func int642str(v int64) (s string) {
	s = strconv.FormatInt(int64(v), 10)
	return
}
func int322str(v int32) (s string) {
	s = strconv.FormatInt(int64(v), 10)
	return
}
func int162str(v int16) (s string) {
	s = strconv.FormatInt(int64(v), 10)
	return
}
func int82str(v int8) (s string) {
	s = strconv.FormatInt(int64(v), 10)
	return
}

func printStack(x interface{}, log logger.Logger) {
	if log != nil {
		log.Error(x)
	}

	for i := 0; i < 20; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			break
		}
		if log != nil {
			log.Errorf("frame %%v:[%%v,file:%%v,line:%%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}
}

`, property.PackageName)
	formatSrc, _ := format.Source([]byte(s))
	if err == nil {
		outputF.WriteString(string(formatSrc))
	} else {
		outputF.WriteString(s)
	}
}
