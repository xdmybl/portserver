package main

import (
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"os"
	"strings"
	"sync"
)

// 根据 pool , section , big_group, small_group, port 的概念分层,不可以跳过创建 .
// section , big_group, small_group, port  都保存他的父层级的 id.  pool 不保存, 因为 pool 上层是 total_pool

// BigJson
// 保存到文件,启动时加载,程序退出时或隔一段时间保存一次到本地
type BigJson struct {
	sync.Mutex
	MaxId        int `json:"max_id"`
	TotalPool    `json:"totalPool"`
	PoolLs       `json:"pool_ls"`
	SectionLs    `json:"section_ls"`
	BigGroupLs   `json:"big_group_ls"`
	SmallGroupLs `json:"small_group_ls"`
	PortLs       `json:"port_ls"`
}

// SaveToFile
// 保存 BigJson 到文件
func (b *BigJson) SaveToFile() error {
	jsonB, err := json.Marshal(b)
	if err != nil {
		logger.Debug("json ")
		return err
	}
	_, err = os.Stat("./data")
	fd := &os.File{}
	if err != nil {
		fd, err = os.Create("./data")
		if err != nil {
			return err
		}
	} else {
		fd, err = os.OpenFile("./data", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
	}
	defer fd.Close()
	fd.Write(jsonB)
	return nil
}

func (b *BigJson) UpdatePoolLsFromId(id int, pool Pool) error {
	newPLs := PoolLs{}
	for _, _pool := range b.PoolLs {
		if _pool.Id != id {
			newPLs = append(newPLs, _pool)
		}
	}
	newPLs = append(newPLs, pool)
	b.PoolLs = newPLs
	return nil
}

func (b *BigJson) UpdateSectionLsFromId(id int, section Section) error {
	newSLs := SectionLs{}
	for _, _section := range b.SectionLs {
		if _section.Id != id {
			newSLs = append(newSLs, _section)
		}
	}
	newSLs = append(newSLs, section)
	b.SectionLs = newSLs
	return nil
}

func (b *BigJson) UpdateBigGroupLsFromId(id int, bigGroup BigGroup) error {
	bigGroupLs := BigGroupLs{}
	for _, _bigGroup := range b.BigGroupLs {
		if _bigGroup.Id != id {
			bigGroupLs = append(bigGroupLs, _bigGroup)
		}
	}
	bigGroupLs = append(bigGroupLs, bigGroup)
	b.BigGroupLs = bigGroupLs
	return nil
}

func (b *BigJson) UpdateSmallGroupLsFromId(id int, bigGroup SmallGroup) error {
	smallGroupLs := SmallGroupLs{}
	for _, _smallGroup := range b.SmallGroupLs {
		if _smallGroup.Id != id {
			smallGroupLs = append(smallGroupLs, _smallGroup)
		}
	}
	smallGroupLs = append(smallGroupLs, bigGroup)
	b.SmallGroupLs = smallGroupLs
	return nil
}

func (b *BigJson) UpdatePortLsFromId(id int, port Port) error {
	portLs := PortLs{}
	for _, _smallGroup := range b.PortLs {
		if _smallGroup.Id != id {
			portLs = append(portLs, _smallGroup)
		}
	}
	portLs = append(portLs, port)
	b.PortLs = portLs
	return nil
}

// LoadJson
// 启动时加载 Json
func LoadJson() (*BigJson, error) {
	bigJson := &BigJson{}
	_, err := os.Stat("./data")
	if err != nil {
		bigJson.TotalPool = *GlobalTotalPool
		err = bigJson.SaveToFile()
		if err != nil {
			os.Exit(-1)
		}
		return bigJson, nil
	}
	jsonB, err := os.ReadFile("./data")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonB, bigJson)
	if err != nil {
		return nil, err
	}
	return bigJson, nil
}

var GlobalTotalPool = new(TotalPool)

type TotalPool struct {
	*BasePorter
}

func (t *TotalPool) Init() *TotalPool {
	t.BasePorter = &BasePorter{
		Id:          0,
		Name:        "total_pool",
		DisplayName: "总池",
		PortRangeLs: []PortRange{*TotalPoolRange},
		Description: "总的端口池",
		Creator:     "xiaoy",
		CreateTime:  "2023/3/6",
	}
	return t
}

type PortRange struct {
	PortMax int `json:"port_max"`
	PortMin int `json:"port_min"`
}

// BasePorter
// 抽象类, 将使用 PortRanger 的概念抽象
type BasePorter struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	PortRangeLs []PortRange `json:"port_range_ls"`
	Description string      `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"create_time"`
}

// Pool 池
type Pool struct {
	BasePorter
}

type PoolLs []Pool

func (pLs *PoolLs) Get(id int) (Pool, error) {
	ret := Pool{}
	for _, pool := range *pLs {
		if id == pool.Id {
			ret = pool
		}
		return ret, nil
	}
	return ret, fmt.Errorf("no found pool id %d", id)
}

type SectionLs []Section

// Section 段
type Section struct {
	BasePorter
	SuperId int `json:"super_id"` // 父 id
}

func (sLs *SectionLs) Get(id int) (Section, error) {
	ret := Section{}
	for _, sec := range *sLs {
		if id == sec.Id {
			ret = sec
		}
		return ret, nil
	}
	return ret, fmt.Errorf("no found sec id %d", id)
}

type BaseFilter struct {
	SuperId                int    `json:"super_id"`
	Name                   string `json:"name"`
	DescriptionLikePattern string `json:"description_like"`
	Page                   int    `json:"page"`
	Size                   int    `json:"size"`
}

// LikeFilter golang 实现 mysql
// Like pattern
func LikeFilter(s, pattern string) bool {
	if pattern == "" {
		return true
	}
	if pattern == "%" {
		return true
	}
	if !strings.Contains(pattern, "%") {
		return s == pattern
	}
	prefix := strings.Split(pattern, "%")[0]
	if prefix != "" && !strings.HasPrefix(s, prefix) {
		return false
	}
	suffix := strings.Split(pattern, "%")[1]
	if suffix != "" && !strings.HasSuffix(s, suffix) {
		return false
	}
	return true
}

// SectionFilter 除了 size 不需要考虑
// 每个对象必须满足 filter 的所有条件, 通过返回 true, 不通过返回 false
func SectionFilter(section Section, filter BaseFilter) bool {
	// super id = 0 则不过滤
	if filter.SuperId != 0 && section.SuperId != filter.SuperId {
		return false
	}
	if filter.Name != "" && section.Name != filter.Name {
		return false
	}
	if filter.DescriptionLikePattern != "" && !LikeFilter(section.Description, filter.DescriptionLikePattern) {
		return false
	}
	// 都满足则返回 true
	return true
}

// SectionLsFilter
// PoolLsFilter
// 纯函数
// sectionLs 为消费队列, retLs 为响应的队列, 采用递归方式执行
func SectionLsFilter(sectionLs SectionLs, retLs SectionLs, filter BaseFilter) SectionLs {
	filteredLs := SectionLs{}
	for _, section := range sectionLs {
		if SectionFilter(section, filter) == true {
			filteredLs = append(filteredLs, section)
		}
	}
	retLs = filteredLs[filter.Page*filter.Size : (filter.Page+1)*filter.Size]
	return retLs
}

func (sLs *SectionLs) Filter(filter BaseFilter) SectionLs {
	return SectionLsFilter(*sLs, SectionLs{}, filter)
}

// BigGroupLs 大组
type BigGroupLs []BigGroup

type BigGroup struct {
	BasePorter
	SuperId int `json:"super_id"` // 父 id
}

func (bgLs *BigGroupLs) Get(id int) (BigGroup, error) {
	ret := BigGroup{}
	for _, bigGroup := range *bgLs {
		if id == bigGroup.Id {
			ret = bigGroup
		}
		return ret, nil
	}
	return ret, fmt.Errorf("no found bigGroup id %d", id)
}

func BigGroupFilter(bigGroup BigGroup, filter BaseFilter) bool {
	// super id = 0 则不过滤
	if filter.SuperId != 0 && bigGroup.SuperId != filter.SuperId {
		return false
	}
	if filter.Name != "" && bigGroup.Name != filter.Name {
		return false
	}
	if filter.DescriptionLikePattern != "" && !LikeFilter(bigGroup.Description, filter.DescriptionLikePattern) {
		return false
	}
	// 都满足则返回 true
	return true
}

func BigGroupLsFilter(bigGroupLs BigGroupLs, retLs BigGroupLs, filter BaseFilter) BigGroupLs {
	filteredLs := BigGroupLs{}
	for _, bigGroup := range bigGroupLs {
		if BigGroupFilter(bigGroup, filter) == true {
			filteredLs = append(filteredLs, bigGroup)
		}
	}
	retLs = filteredLs[filter.Page*filter.Size : (filter.Page+1)*filter.Size]
	return retLs
}

func (bgLs *BigGroupLs) Filter(filter BaseFilter) BigGroupLs {
	return BigGroupLsFilter(*bgLs, BigGroupLs{}, filter)
}

// SmallGroupLs
// SmallGroup
// 小组
type SmallGroupLs []SmallGroup

type SmallGroup struct {
	BasePorter
	SuperId int `json:"super_id"` // 父 id
}

func (sgLs *SmallGroupLs) Get(id int) (SmallGroup, error) {
	ret := SmallGroup{}
	for _, smallGroup := range *sgLs {
		if id == smallGroup.Id {
			ret = smallGroup
		}
		return ret, nil
	}
	return ret, fmt.Errorf("no found smallGroup id %d", id)
}

type PortLs []Port

type Port struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SuperId     int    `json:"super_id"` // 父 id
	DisplayName string `json:"display_name"`
	//PortRangeLs []PortRange `json:"port_range_ls"`
	PortValue   int    `json:"port_value"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	CreateTime  string `json:"create_time"`
}

func (portLs *PortLs) Get(id int) (Port, error) {
	ret := Port{}
	for _, port := range *portLs {
		if id == port.Id {
			ret = port
		}
		return ret, nil
	}
	return ret, fmt.Errorf("no found port id %d", id)
}
