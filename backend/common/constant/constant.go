package constant

import "time"

const (
	USER_ID           = "uid"         // 用户ID
	USER_SESSION_MARK = "user_info"   // 用户信息
	MENU_CACHE        = "menu_cache_" // 登陆用户的菜单列表缓存前缀
)

type BunissType int

//业务类型
const (
	Buniss_Other BunissType = 0 // 0其它
	Buniss_Add   BunissType = 1 // 1新增
	Buniss_Edit  BunissType = 2 // 2修改
	Buniss_Del   BunissType = 3 // 3删除
)

//响应结果
const (
	SUCCESS      = 0   // 成功
	ERROR        = 500 // 错误
	UNAUTHORIZED = 403 // 无权限
	FAIL         = -1  // 失败
)

//错误处理页面
const (
	ERROR_PAGE  = "core/errpage/error.html"  // 错误提示页面
	UNAUTH_PAGE = "core/errpage/unauth.html" // 无权限提示页面
)

// 通用api响应
type CommonRes struct {
	Code  int         `json:"code"`           // 响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg   string      `json:"msg"`            // 消息
	Data  interface{} `json:"data,omitempty"` // 数据内容
	Btype BunissType  `json:"otype"`          // 业务类型
}

// 通用分页表格响应
type TableDataInfo struct {
	Total int         `json:"total"` // 总数
	Rows  interface{} `json:"rows"`  // 数据
	Code  int         `json:"code"`  // 响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   // 消息
}

// 通用的树形结构
type Ztree struct {
	Id      int    `json:"id"`      // 节点ID
	Pid     int    `json:"pId"`     // 节点父ID
	Name    string `json:"name"`    // 节点名称
	Title   string `json:"title"`   // 节点标题
	Checked bool   `json:"checked"` // 是否勾选
	Open    bool   `json:"open"`    // 是否展开
	Nocheck bool   `json:"nocheck"` // 是否能勾选
}

// 通用分页请求参数
type PageReq struct {
	BeginTime string `form:"beginTime"`     //开始时间
	EndTime   string `form:"endTime"`       //结束时间
	PageNum   int    `form:"pageNum"`       //当前页码
	PageSize  int    `form:"pageSize"`      //每页数
	SortName  string `form:"orderByColumn"` //排序字段
	SortOrder string `form:"isAsc"`         //排序方式
	Total     *int   // 分页数据总条数
}

// 模型数据
type ModelData struct {
	CreateBy   string    `db:"create_by" json:"create_by"`     // 创建者
	CreateTime time.Time `db:"create_time" json:"create_time"` // 创建时间
	UpdateBy   string    `db:"update_by" json:"update_by"`     // 更新者
	UpdateTime time.Time `db:"update_time" json:"update_time"` // 更新时间
	Remark     string    `db:"remark" json:"remark"`           // 备注
}

// 通用的删除请求
type RemoveReq struct {
	Ids string `form:"ids" binding:"required"`
}
