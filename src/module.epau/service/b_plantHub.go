/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/31 2:56
 * @version     v1.0
 * @filename    b_plantHub.go
 * @description
 ***************************************************************************/
package service

import (
	"cms/src/core"
	"encoding/csv"
	"github.com/google/wire"
	"os"
)

var PlantHubSet = wire.NewSet(wire.Struct(new(BPlantHub), "*"))

type BPlantHub struct {
}

func (bPlantHub *BPlantHub) ReadData(fileName string) ([][]string, *core.Error) {
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
