package main

import (
	"crypto/ecdsa"
	"math/big"
)

type Data struct {
	DataID     string `json:"dataID"`     // 车牌号
	Cid        string `json:"cid"`        // ipfs的返回值
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

type Policy struct {
	SubjectA       SA `json:"subjectA"`       // 主题属性(用户属性)
	DataA          DA `json:"dataA"`          // 数据属性
	ActionA        AA `json:"actionA"`        // 行为属性(权限属性)
	EnvironmentalA EA `json:"environmentalA"` // 环境属性
}

type SA struct {
	UserID string `json:"userID"` // 用户ID
	Role   string `json:"role"`   // 角色
	PK     string `json:"PK"`     // 用户的公钥
}

type DA struct {
	DataID string `json:"dataID"` // 数据ID
	Owner  string `json:"owner"`  // 所有者
	Key    string `json:"key"`    // 键值
}

type AA struct {
	Level int `json:"level"` // 操作权限等级
}

type EA struct {
	CreateTime string `json:"createTime"` // 创建时间
	EndTime    string `json:"endTime"`    // 结束时间
	Address    string `json:"address"`    // 地址
}

// AP 访问控制策略
type AP struct {
	PolicyID   string `json:"policyID"`
	RequestID  string `json:"requestID"`
	ResponseID string `json:"responseID"`
	Policy     Policy `json:"policy"` // 访问控制策略
	Sign       string `json:"sign"`   // 数字签名
}

// Sign 数字签名
type Sign struct {
	R *big.Int
	S *big.Int
}

type User struct {
	UserID     string            `json:"userID"`     // 用户ID
	Lnp        string            `json:"lnp"`        // 车牌号
	Org        string            `json:"org"`        // 组织
	Department string            `json:"department"` // 部门
	Position   string            `json:"position"`   // 职称
	PublicKey  *ecdsa.PublicKey  `json:"publicKey"`  // ECDSA公钥
	PrivateKey *ecdsa.PrivateKey `json:"privateKey"` // ECDSA私钥
}

type ActionRecord struct {
	RecordID  string `json:"recordID"`
	PolicyID  string `json:"policyID"`
	UserID    string `json:"userID"`    // 用户ID
	DataID    string `json:"dataID"`    // 数据ID
	Action    string `json:"action"`    // 操作
	TimeStamp string `json:"timeStamp"` // 时间
}

type RequestRecord struct {
	RequestID string `json:"requestID"`
	SubjectA  SA     `json:"subjectA"`
	DataA     DA     `json:"dataA"`
	Level     int    `json:"level"`
	TimeStamp string `json:"timeStamp"`
}

type ResponseRecord struct {
	ResponseID string `json:"responseID"`
	PolicyID   string `json:"policyID"`
	Owner      string `json:"owner"`
	RequestID  string `json:"requestID"`
	EndTime    string `json:"endTime"`
	TimeStamp  string `json:"timeStamp"`
}
