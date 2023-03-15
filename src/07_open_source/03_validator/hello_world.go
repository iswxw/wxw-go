// @Time : 2023/3/15 17:38
// @Author : xiaoweiwei
// @File : hello_world

package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Inner struct {
	StartDate time.Time
}

type Outer struct {
	InnerStructField *Inner
	CreatedAt        time.Time `validate:"ltecsfield=InnerStructField.StartDate"`
}

// NOTE: when calling validate.Struct(val) topStruct will be the top level struct passed
//
//	into the function
//	when calling validate.VarWithValue(val, field, tag) val will be
//	whatever you pass, struct, field...
//	when calling validate.Field(field, tag) val will be nil
//
// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Singleton
func main() {
	now := time.Now()

	inner := &Inner{
		StartDate: now,
	}

	outer := &Outer{
		InnerStructField: inner,
		CreatedAt:        now,
	}

	err := validator.New().Struct(outer)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("check result = success")
}
