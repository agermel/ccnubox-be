package biz

import (
	"context"
	"github.com/asynccnu/ccnubox-be/be-classlist/internal/model"
	"github.com/google/wire"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewClassUsecase, NewClassInfoRepo, NewStudentAndCourseRepo, NewClassRepo)

type Transaction interface {
	// 下面2个方法配合使用，在InTx方法中执行ORM操作的时候需要使用DB方法获取db！
	InTx(ctx context.Context, fn func(ctx context.Context) error) error
	DB(ctx context.Context) *gorm.DB
}
type ClassCrawler interface {
	//获取本科生的课表
	GetClassInfosForUndergraduate(ctx context.Context, req model.GetClassInfosForUndergraduateReq) (*model.GetClassInfosForUndergraduateResp, error)
	//获取研究生的课表(未实现)
	GetClassInfoForGraduateStudent(ctx context.Context, req model.GetClassInfoForGraduateStudentReq) (*model.GetClassInfoForGraduateStudentResp, error)
}
type ClassRepoProxy interface {
	//保存课程
	SaveClass(ctx context.Context, stuID, year, semester string, classInfos []*model.ClassInfo, scs []*model.StudentCourse) error
	//获取某个学生某个学期的所有课程
	GetClassesFromLocal(ctx context.Context, req model.GetClassesFromLocalReq) (*model.GetClassesFromLocalResp, error)
	//只获取特定ID的class_info
	GetSpecificClassInfo(ctx context.Context, req model.GetSpecificClassInfoReq) (*model.GetSpecificClassInfoResp, error)
	//添加课程
	AddClass(ctx context.Context, req model.AddClassReq) error
	//删除课程
	DeleteClass(ctx context.Context, req model.DeleteClassReq) error
	//获取某个学生某个学期的处于回收站的课程ID
	GetRecycledIds(ctx context.Context, req model.GetRecycledIdsReq) (*model.GetRecycledIdsResp, error)
	//从回收站中移除课程
	RemoveClassFromRecycledBin(ctx context.Context, req model.RemoveClassFromRecycleBinReq) error
	//更新课程
	UpdateClass(ctx context.Context, req model.UpdateClassReq) error
	//判断课程和学生ID是否有联系
	CheckSCIdsExist(ctx context.Context, req model.CheckSCIdsExistReq) bool
	//获取全校某个学期的所有课程
	GetAllSchoolClassInfos(ctx context.Context, req model.GetAllSchoolClassInfosReq) *model.GetAllSchoolClassInfosResp
	//检查某个class是否存在于回收站中
	CheckClassIdIsInRecycledBin(ctx context.Context, req model.CheckClassIdIsInRecycledBinReq) bool
	//检查回收的课程的是否手动添加
	IsRecycledCourseManual(ctx context.Context, req model.IsRecycledCourseManualReq) bool
	//获取某个学生某个学期的手动添加的课程[直接来自数据库]
	GetAddedClasses(ctx context.Context, req model.GetAddedClassesReq) (*model.GetAddedClassesResp, error)
}
type JxbRepo interface {
	//保存教学班
	SaveJxb(ctx context.Context, stuID string, jxbID []string) error
	//根据教学班ID查询stuID
	FindStuIdsByJxbId(ctx context.Context, jxbId string) ([]string, error)
}
type CCNUServiceProxy interface {
	//从其他服务获取cookie
	GetCookie(ctx context.Context, stuID string) (string, error)
}

type RefreshLogRepo interface {
	InsertRefreshLog(ctx context.Context, stuID, year, semester string) (uint64, error)
	UpdateRefreshLogStatus(ctx context.Context, logID uint64, status string) error
	SearchRefreshLog(ctx context.Context, stuID, year, semester string) (*model.ClassRefreshLog, error)
	GetRefreshLogByID(ctx context.Context, logID uint64) (*model.ClassRefreshLog, error)
	GetLastRefreshTime(ctx context.Context, stuID, year, semester string, beforeTime time.Time) *time.Time
}

type DelayQueue interface {
	Send(key, value []byte) error
	Consume(groupID string, f func(key, value []byte)) error
	Close()
}
