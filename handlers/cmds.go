package handlers

import (
	defs "curiosity-questions/question-1/defs"
	"errors"
	"os"
	"strconv"
)

func CheckAndGetCmdParams(params []string) (totalRequestNum, concurrentNum int, err error) {
	if len(os.Args) != defs.CMD_PARAM_NUM {
		return 0, 0, errors.New("Not enough params，2 cmd params requried")
	}
	// 获取命令行参数
	a, b := os.Args[1], os.Args[2]
	// 总共的任务数
	totalRequestNum, err = strconv.Atoi(a)
	if err != nil {
		return 0, 0, errors.New("Need to enter numbers")
	}
	// 并发数
	concurrentNum, err = strconv.Atoi(b)
	if err != nil {
		return 0, 0, errors.New("Need to enter numbers")
	}
	return totalRequestNum, concurrentNum, nil
}
