/*
@Time : 2021/10/5 00:01
@Author : wxw
@File : test
*/
package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// 1. 简单的日志打印
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello world")
	// Output: {"time":1516134303,"level":"debug","message":"hello world"}

	fmt.Println("-----------------------")

	// 2. 上下文字段嵌入型 日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().Str("Scale", "833 cents").Float64("Interval", 833.09).Msg("Fibonacci is everywhere")
	log.Debug().Str("Name", "Tom").Send()
	// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
	// Output: {"level":"debug","Name":"Tom","time":1562212768}

	fmt.Println("-----------------------")

	// 3. 分级打印日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("hello world")
	// Output: {"time":1516134303,"level":"info","message":"hello world"}

	fmt.Println("-----------------------")

	// 4. 不区分等级的日志
	log.Log().Str("foo", "bar").Msg("")
	// {"foo":"bar","time":1633362536}

	fmt.Println("-----------------------")

	// 5. 自定义字段日志
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	log.Info().Msg("hello world")

	// Output: {"l":"info","t":1494567715,"m":"hello world"}

	fmt.Println("-----------------------")

	// 6. 错误日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := errors.New("seems we have an error here")
	log.Error().Err(err).Msg("")
}
