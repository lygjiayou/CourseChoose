package service

import (
	"CourseChoose/model"
	_ "github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

// CheckCreateMemberParamService 参数校验服务
func CheckCreateMemberParamService(req *model.CreateMemberRequest) bool {
	//var resp model.CreateMemberRespone
	//member := &model.Member{
	//	Nickname: req.Nickname,
	//	Username: req.Username,
	//	Password: req.Password,
	//}
	//if member
	nickname := req.Nickname
	username := req.Username
	password := req.Password
	if len(nickname)<4 || len(nickname)>20 {
		//resp.Code = model.ParamInvalid
		return false
	}
	if len(username)<8 || len(username)>20 {
		return false
	}
	if len(password)<8 || len(password)>20 {
		return false
	}
	// 用户名是否支持大小写(有大写也好，有小写也好，同时有大小写也好)
	for i:= range username {
		if (username[i] >= 'a' && username[i] <= 'z') || (username[i] >= 'A' && username[i] <= 'Z') {
			continue
		} else {
			return false
		}
	}
	// 密码是否同时包括大、小写和数字，利用正则表达式
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`}
	for _, pattern := range patternList {
		matchOK, _ := regexp.MatchString(pattern, password)
		if !matchOK {
			return false
		}
	}
	return true
}

// CreateMemberService 创建成员服务
func CreateMemberService(req *model.CreateMemberRequest) *model.CreateMemberRespone {
	var resp model.CreateMemberRespone
	member := &model.Member{
		Nickname: req.Nickname,
		Username: req.Username,
		Password: req.Password,
		UserType: req.UserType,
	}
	err := member.CreateMember()
	if err != nil {
		// 如果出现错误返回用户已经存在，因为可能被软删除了
		resp.Code = model.ParamInvalid
	} else {
		resp.Code = model.OK
		resp.Data = struct{ UserID string }{UserID: strconv.Itoa(member.UserID)}
	}
	// 如果是学生，添加到本地映射中
	if req.UserType == model.Student {
		model.StudentList[member.UserID] = struct{}{} // 这个UserID是数据库生成的
	}
	return &resp
}
