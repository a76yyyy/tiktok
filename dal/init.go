/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-07 23:23:09
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:46:55
 * @FilePath: /tiktok/dal/init.go
 * @Description: 初始化数据层
 */

package dal

import (
	db "github.com/a76yyyy/tiktok/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
