package models

import "time"

type Address struct {
	Area      string `json:"area"`  //区域 类似 S I O Q WR ……
	Slave     uint8  `json:"slave"` //从站号 modbus
	Offset    uint16 `json:"offset"`
	ReadCode  uint8  `json:"read_code"`
	WriteCode uint8  `json:"write_code"`
}

type Element struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"` //来源
	UUID        string `json:"uuid"`   //唯一码，自动生成

	Manufacturer string `json:"manufacturer"` //厂商
	Model        string `json:"model"`        //型号
	Version      string `json:"version"`      //版本

	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}

type ElementVariable struct {
	Id        int64 `json:"id"`
	ElementId int64 `json:"element_id"`

	Offset    uint16 `json:"offset"`
	ReadCode  uint8  `json:"read_code"`
	WriteCode uint8  `json:"write_code"`

	Name  string `json:"name"`
	Alias string `json:"alias"` //别名，用于编程
	Type  string `json:"type"`
	Unit  string `json:"unit"` //单位

	Scale   float32 `json:"scale"`   //倍率，比如一般是 整数÷10，得到
	Correct float32 `json:"correct"` //校准

	Default  string `json:"default"`
	Writable bool   `json:"writable"` //可写，用于输出（如开关）

	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}

type ElementBatch struct {
	Id        int64 `json:"id"`
	ElementId int64 `json:"element_id"`

	Type   string `json:"type"` //read write
	Code   uint8  `json:"code"` //功能码 3,4
	Offset uint16 `json:"offset"`
	Size   uint16 `json:"size"`

	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}
