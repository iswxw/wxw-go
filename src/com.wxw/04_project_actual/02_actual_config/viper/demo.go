/*
@Time: 2022/1/13 0:23
@Author: wxw
@File: demo
@Link: https://www.liwenzhou.com/posts/Go/viper_tutorial/
*/
package viper

import "github.com/spf13/viper"

// 设置默认值
func setDefault() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}
