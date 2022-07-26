package model

type Member struct {
	UserID	 int		`gorm:"primarykey;column:userid"`	// 用户id
	Nickname string		`gorm:"column:nickname"`			// 昵称
	Username string		`gorm:"column:username"`			// 用户名
	Password string		`gorm:"column:password"`			// 密码
	UserType UserType	`gorm:"column:usertype"`			// 类型（2：学生，1：管理员 3：教师）
	State 	 bool		`gorm:"column:state"`				// 状态，已删除为true,否则为false
}

func (Member) TableName() string {
	return "member" // 类似于构造函数，返回给Member成员，告诉数据库的名字是member(这个member是要和数据库的表名字对应)
}

// StudentList 存放学生的映射，用以判断学生是否存在
var StudentList = make(map[int]struct{})

// CreateIndexList 初始时将学生信息放入映射
func CreateIndexList() {
	var members []Member
	db.Where("UserType=?", Student).Select("userid").Find(&members)
	for i:=0;i<len(members); i++ {
		StudentList[members[i].UserID] = struct{}{}
	}
}

// CreateMember 创建用户
func (member *Member) CreateMember() error { // error内置接口类型是表示错误条件的常规接口，nil值表示没有错误
	err := db.Create(member).Error
	return err
}

