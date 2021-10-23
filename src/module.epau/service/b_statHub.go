/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/5 23:18
 * @version     v1.0
 * @filename    b_statHub.go
 * @description
 ***************************************************************************/
package service

import (
	"cms/src/core"
	"encoding/csv"
	"github.com/google/wire"
	"os"
)

var StatHubSet = wire.NewSet(wire.Struct(new(BStatHub), "*"))

type BStatHub struct {
}

func (bStatHub *BStatHub) ReadData(fileName string) ([][]string, *core.Error) {
	// deal with small files, read lines all
	fs, err := os.Open(fileName)
	if err != nil {
		return nil, core.NewError(200, err)
	}
	defer fs.Close()
	reader := csv.NewReader(fs)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, core.NewError(201, err)
	}
	return content, nil
}
