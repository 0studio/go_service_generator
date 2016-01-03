package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type pbInfo struct {
	reqType  string
	dataType string
	goType   string
}

var DefaultProtoBufTypeMap map[string]pbInfo = map[string]pbInfo{
	"bool":              pbInfo{goType: "bool", dataType: "bool", reqType: "optional"},
	"int":               pbInfo{goType: "int32", dataType: "int32", reqType: "optional"},
	"int8":              pbInfo{goType: "int32", dataType: "int32", reqType: "optional"},
	"int16":             pbInfo{goType: "int32", dataType: "int32", reqType: "optional"},
	"int32":             pbInfo{goType: "int32", dataType: "int32", reqType: "optional"},
	"int64":             pbInfo{goType: "int64", dataType: "int64", reqType: "optional"},
	"uint8":             pbInfo{goType: "uint32", dataType: "uint32", reqType: "optional"},
	"uint16":            pbInfo{goType: "uint32", dataType: "uint32", reqType: "optional"},
	"uint32":            pbInfo{goType: "uint32", dataType: "uint32", reqType: "optional"},
	"uint64":            pbInfo{goType: "uint64", dataType: "uint64", reqType: "optional"},
	"float32":           pbInfo{goType: "float64", dataType: "double", reqType: "optional"},
	"float64":           pbInfo{goType: "float64", dataType: "double", reqType: "optional"},
	"string":            pbInfo{goType: "string", dataType: "string", reqType: "optional"},
	"time.Time":         pbInfo{goType: "int64", dataType: "int64", reqType: "optional"},
	"key.KeyUint64":     pbInfo{goType: "uint64", dataType: "uint64", reqType: "optional"},
	"key.KeyInt":        pbInfo{goType: "int64", dataType: "int64", reqType: "optional"},
	"key.KeyInt32":      pbInfo{goType: "int32", dataType: "int32", reqType: "optional"},
	"key.String":        pbInfo{goType: "string", dataType: "string", reqType: "optional"},
	"key.KeyString":     pbInfo{goType: "string", dataType: "string", reqType: "optional"},
	"[]int":             pbInfo{goType: "[]int64", dataType: "int64", reqType: "repeated"},
	"[]int32":           pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
	"[]int8":            pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
	"[]int16":           pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
	"[]int64":           pbInfo{goType: "[]int64", dataType: "int64", reqType: "repeated"},
	"[]uint32":          pbInfo{goType: "[]uint32", dataType: "uint32", reqType: "repeated"},
	"[]uint8":           pbInfo{goType: "[]uint32", dataType: "uint32", reqType: "repeated"},
	"[]uint16":          pbInfo{goType: "[]uint32", dataType: "uint32", reqType: "repeated"},
	"[]uint64":          pbInfo{goType: "[]uint64", dataType: "uint64", reqType: "repeated"},
	"goutils.Int32List": pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
	"goutils.Int16List": pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
	"goutils.IntList":   pbInfo{goType: "[]int64", dataType: "int64", reqType: "repeated"},
	"goutils.Int8List":  pbInfo{goType: "[]int32", dataType: "int32", reqType: "repeated"},
}

func (fd FieldDescriptoin) getPBInfoFromTag() (info pbInfo) {
	pbType := fd.ProtoBufTagFieldList.GetValue("type")
	if pbType != "" {
		info.dataType = pbType
		if fd.ProtoBufTagFieldList.Contains("repeated") {
			info.reqType = "repeated"
		} else if fd.ProtoBufTagFieldList.Contains("required") {
			info.reqType = "required"
		} else {
			info.reqType = "optional"
		}
		info.goType = DefaultProtoBufTypeMap[fd.FieldGoType].goType
		return
	}
	return
}
func (fd FieldDescriptoin) getPBInfo() (info pbInfo) {
	info = fd.getPBInfoFromTag()
	if info.dataType != "" {
		return
	}

	var ok bool
	info, ok = DefaultProtoBufTypeMap[fd.FieldGoType]
	if ok {
		return
	}
	if strings.HasPrefix(fd.FieldGoType, "[]") {
		info.reqType = "repeated"
		info.dataType = DefaultProtoBufTypeMap[fd.FieldGoType[2:]].dataType
		return
	}
	return

}

func (sd StructDescription) GenerateGogoPB(property Property, srcDir string) {
	pkFieldList := sd.GetPKFieldList()
	if len(pkFieldList) > 2 {
		return
	}

	sd.generateGogoPB(property, srcDir)
	sd.GeneratePBBuildScript(property, srcDir)
	sd.GenerateEntitySerialUnSerial(property, srcDir)

}
func (sd StructDescription) generateGogoPB(property Property, srcDir string) {
	outputF, err := os.OpenFile(filepath.Join(srcDir, fmt.Sprintf("serial_%s.proto", strings.ToLower(sd.StructName))), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputF.Close()

	outputF.WriteString(
		fmt.Sprintf(`// do not edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
package %s;

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.unmarshaler_all) = true;

`, property.PackageName))

	outputF.WriteString(fmt.Sprintf("message %sPB{\n", sd.StructName))

	var nullable string
	for idx, field := range sd.Fields {
		pbInfo := field.getPBInfo()
		if pbInfo.reqType != "repeated" {
			nullable = " [(gogoproto.nullable) = false]"
		} else {
			nullable = ""
		}
		outputF.WriteString(fmt.Sprintf("    %s %s%s%s%s= %2d%s;\n",
			pbInfo.reqType, pbInfo.dataType, makeString(" ", 10-len(pbInfo.dataType)), field.FieldName,
			makeString(" ", 16-len(field.FieldName)), idx+1, nullable))
	}
	outputF.WriteString("}\n")
}

func (sd StructDescription) GeneratePBBuildScript(property Property, srcDir string) {
	scriptFileName, _ := filepath.Abs(filepath.Join(srcDir, fmt.Sprintf("build_proto.sh")))

	if !IsFileExists(scriptFileName) {
		outputF, err := os.OpenFile(scriptFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer outputF.Close()
		outputF.WriteString(
			`# -*- coding:utf-8 -*-
#you can edit this file ,this is generated by tools(https://github.com/0studio/go_service_generator)
# you need install protobuf and github.com/gogo/protobuf
# and make sure protoc is in your $PATH
#brew install protobuf
#go get github.com/gogo/protobuf/proto
#go get github.com/gogo/protobuf/protoc-gen-gogo
#go get github.com/gogo/protobuf/gogoproto
protoc --proto_path=$GOPATH/src/github.com/gogo/protobuf/protobuf:$GOPATH/src:. --gogo_out=.  serial.proto
`)

	}

	cmd := exec.Command("sh", scriptFileName)
	cmd.Dir = srcDir
	err := cmd.Run()
	if err != nil {
		fmt.Println("-------------------------------------------------")
		fmt.Println("[Warning]:running build_proto.sh failled\nmaybe you need edit this file  and $GOPATH is not empty\nand make sure protobuf and gogo/protobuf/proto is installed on your os")
		fmt.Println(err)
	}

}

func IsFileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)

	// _, err := os.Stat(name)
	// if os.IsNotExist(err) {
	// 	return false, nil
	// }
	// return err != nil, err
}
