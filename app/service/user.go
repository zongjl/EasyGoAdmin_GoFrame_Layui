/**
 *
 * @author 摆渡人
 * @since 2021/7/27
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// 中间件管理服务
var User = new(userService)

type userService struct{}

func (s *userService) GetList(req *model.UserPageReq) ([]model.UserInfoVo, int, error) {
	// 创建查询实例
	query := dao.User.Where("mark=1")
	// 查询条件
	if req != nil {
		// 用户姓名
		query = query.Where("realname like ?", "%"+req.Realname+"%")
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.User
	query.Structs(&list)

	// 获取职级列表
	levelList, _ := dao.Level.Where("mark=1").Fields("id,name").All()
	var levelMap = make(map[int]string)
	for _, v := range levelList {
		levelMap[v.Id] = v.Name
	}
	// 获取岗位列表
	positionList, _ := dao.Position.Where("mark=1").Fields("id,name").All()
	var positionMap = make(map[int]string)
	for _, v := range positionList {
		positionMap[v.Id] = v.Name
	}
	// 部门
	deptList, _ := dao.Dept.Where("mark=1").Fields("id,name").All()
	var deptMap = make(map[int]string)
	for _, v := range deptList {
		deptMap[v.Id] = v.Name
	}

	// 数据处理
	var result []model.UserInfoVo
	for _, v := range list {
		item := model.UserInfoVo{}
		item.User = v
		// 性别
		if v.Gender > 0 {
			item.GenderName = utils.GENDER_LIST[v.Gender]
		}
		// 职级
		if v.LevelId > 0 {
			item.LevelName = levelMap[v.LevelId]
		}
		// 岗位
		if v.PositionId > 0 {
			item.PositionName = positionMap[v.PositionId]
		}
		// 部门
		if v.DeptId > 0 {
			item.DeptName = deptMap[v.DeptId]
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *userService) Add(req *model.UserAddReq) (int64, error) {
	// 实例化对象
	var entity model.User
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Avatar = req.Avatar
	entity.Mobile = req.Mobile
	entity.Email = req.Email
	entity.Birthday = req.Birthday
	entity.DeptId = req.DeptId
	entity.LevelId = req.LevelId
	entity.PositionId = req.PositionId
	entity.ProvinceCode = req.ProvinceCode
	entity.CityCode = req.CityCode
	entity.DistrictCode = req.DistrictCode
	entity.Address = req.Address
	entity.Username = req.Username
	entity.Password = req.Password
	entity.Intro = req.Intro
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}

	// 插入对象
	result, err := dao.User.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *userService) Update(req *model.UserUpdateReq) (int64, error) {
	// 查询记录
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置对象
	info.Realname = req.Realname
	info.Nickname = req.Nickname
	info.Gender = req.Gender
	info.Avatar = req.Avatar
	info.Mobile = req.Mobile
	info.Email = req.Email
	info.Birthday = req.Birthday
	info.DeptId = req.DeptId
	info.LevelId = req.LevelId
	info.PositionId = req.PositionId
	info.ProvinceCode = req.ProvinceCode
	info.CityCode = req.CityCode
	info.DistrictCode = req.DistrictCode
	info.Address = req.Address
	info.Username = req.Username
	info.Intro = req.Intro
	info.Status = req.Status
	info.Note = req.Note
	info.Sort = req.Sort

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		info.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		info.Avatar = avatar
	}

	// 更新记录
	result, err := dao.User.Save(info)
	if err != nil {
		return 0, err
	}

	// 删除用户角色关系
	// 待完善
	userId := 1
	dao.UserRole.Delete("user_id=?", userId)
	// 创建人员角色关系
	roleIds := strings.Split(req.RoleIds, ",")
	for _, v := range roleIds {
		var userRole model.UserRole
		userRole.UserId = userId
		userRole.RoleId = gconv.Int(v)
		dao.UserRole.Insert(userRole)
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.User.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) Status(req *model.UserStatusReq) (int64, error) {
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.User.Data(g.Map{
		"status":      req.Status,
		"update_user": 1,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *userService) ResetPwd(id int) (int64, error) {
	// 查询记录
	info, err := dao.User.FindOne("id=?", id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置初始密码
	password, err := utils.Md5("123456" + info.Username)
	if err != nil {
		return 0, err
	}

	// 初始化密码
	result, err := dao.User.Data(g.Map{
		"password":    password,
		"update_user": 1,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
