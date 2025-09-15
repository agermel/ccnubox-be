package biz

import (
	"fmt"
	"time"
)

type ClassInfo struct {
	ID           string //集合了课程信息的字符串，便于标识（课程ID）
	CreatedAt    time.Time
	UpdatedAt    time.Time
	JxbId        string  //教学班ID
	Day          int64   //星期几
	Teacher      string  //任课教师
	Where        string  //上课地点
	ClassWhen    string  //上课是第几节（如1-2,3-4）
	WeekDuration string  //上课的周数
	Classname    string  //课程名称
	Credit       float64 //学分
	Weeks        int64   //哪些周
	Semester     string  //学期
	Year         string  //学年
	Note         string  //备注
}

func (ci *ClassInfo) UpdateID() {
	ci.ID = fmt.Sprintf("Class:%s:%s:%s:%d:%s:%s:%s:%d", ci.Classname, ci.Year, ci.Semester, ci.Day, ci.ClassWhen, ci.Teacher, ci.Where, ci.Weeks)
}

type StudentCourse struct {
	StuID           string //学号
	ClaID           string //课程ID
	Year            string //学年
	Semester        string //学期
	IsManuallyAdded bool   //是否为手动添加
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
