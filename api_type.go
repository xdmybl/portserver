package main

type BasePortRangeAdd struct {
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	PortRangeLs []PortRange `json:"port_range_ls"`
	Description string      `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"create_time"`
}

type BasePortRangeUpdate struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	PortRangeLs []PortRange `json:"port_range_ls"`
	Description string      `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"create_time"`
}

type SubPortRangeAdd struct {
	Name        string      `json:"name"`
	SuperId     int         `json:"super_id"`
	DisplayName string      `json:"display_name"`
	PortRangeLs []PortRange `json:"port_range_ls"`
	Description string      `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"create_time"`
}

type SubPortRangeUpdate struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	SuperId     int         `json:"super_id"`
	DisplayName string      `json:"display_name"`
	PortRangeLs []PortRange `json:"port_range_ls"`
	Description string      `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"create_time"`
}

type PortAdd struct {
	Name        string `json:"name"`
	SuperId     int    `json:"super_id"` // 父 id
	DisplayName string `json:"display_name"`
	PortValue   int    `json:"port_value"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	CreateTime  string `json:"create_time"`
}

type PortUpdate struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SuperId     int    `json:"super_id"` // 父 id
	DisplayName string `json:"display_name"`
	PortValue   int    `json:"port_value"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	CreateTime  string `json:"create_time"`
}
