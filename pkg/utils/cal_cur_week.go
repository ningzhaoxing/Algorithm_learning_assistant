package utils

import (
	"getQuestionBot/internal/models"
	"time"
)

func CalCurWeek(system models.System) int {
	// 计算当前是第几周
	semesterStart, _ := time.Parse("2006-01-02", system.SemesterStart)
	currentTime := time.Now()
	daysSinceStart := currentTime.Sub(semesterStart).Hours() / 24
	weekNumber := int(daysSinceStart/7) + 1
	return weekNumber
}
