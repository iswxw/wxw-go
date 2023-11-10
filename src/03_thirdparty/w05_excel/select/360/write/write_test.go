// @Time : 2022/8/18 11:47
// @Author : xiaoweiwei
// @File : write_test

package write

import (
	"log"
	"testing"
)

func TestGetPath(t *testing.T) {
	log.Println(GetPath("11"))
}

func TestBatchWrite(t *testing.T) {
	batchWrite()

}

func TestArray2Struct(t *testing.T) {
	Array2Struct()
}
