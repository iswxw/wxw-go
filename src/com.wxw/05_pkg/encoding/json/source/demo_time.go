/*
@Time: 2022/2/22 8:35
@Author: wxw
@File: demo_time
*/
package main

import (
	"fmt"
	"time"
)

// time 重写 MarshalJSON、UnmarshalJSON
func main() {
	t := time.Time{}
	marshalJSON, err := t.MarshalJSON()
	if err != nil {
		return
	}
	if err = t.UnmarshalJSON(nil); err != nil {
		return
	}
	fmt.Println(marshalJSON)
}
