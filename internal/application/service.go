package application

import (
	"fmt"
	"getQuestionBot/internal/dao"
	"getQuestionBot/internal/models"
	"getQuestionBot/internal/service"
	"getQuestionBot/pkg/utils"
	"strconv"
)

type Service struct {
	service.CrawlService
	service.DingtalkService
	service.MessageProcessService
	dao.UserRepo
	dao.SystemRepo
}

func NewService(userRepo dao.UserRepo, sysRepo dao.SystemRepo, craw service.CrawlService, ding service.DingtalkService, msg service.MessageProcessService) *Service {
	return &Service{
		CrawlService:          craw,
		DingtalkService:       ding,
		MessageProcessService: msg,
		UserRepo:              userRepo,
		SystemRepo:            sysRepo,
	}
}

// Apply 获取六期同学力扣刷题信息，并推送到钉钉机器人
func (s *Service) Apply(obj string, websiteName string) {
	// 数据库查询同学刷题网站
	websites, err := s.UserRepo.GetUserAndWebsitesByDepartment(obj, websiteName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取系统变量
	system, err := s.SystemRepo.GetSystemConfigById(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	users := make([]models.User, 0)
	for _, website := range websites {
		// 获取静态资源
		source, err := s.CrawlService.GetPageSource(website.UserURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 解析数据
		user, err := s.MessageProcessService.GetProblemListByPageSource(source)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 计算当前解题周期
		week := utils.CalCurWeek(system)
		for i, _ := range user.Problems {
			user.Problems[i].Week = strconv.Itoa(week)
			user.Problems[i].Term = system.CurTerm
		}

		// 保存题目
		err = s.UserRepo.SaveProblem(user.Problems, website.UserID)
		if err != nil {
			fmt.Println(err)
			return
		}

		user.Name = website.User.Name
		users = append(users, *user)
	}

	// 组装消息
	msg, err := s.MessageProcessService.MessageAssembly(users, system)
	if err != nil {
		return
	}

	// 发送消息
	err = s.DingtalkService.SendMessage(msg)
	if err != nil {
		return
	}
}
