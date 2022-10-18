package model

import (
	"github.com/dgryski/trifles/uuid"
	"github.com/google/uuid" // taskid用
	_ "gorm.io/gorm"
)

/* Task構造体の定義 */
// 仕様定義のパラメータを持つ
// 先頭が大文字でないと、package modelの外から参照できない
type Task struct {
	ID       uuid.UUID
	Name     string
	Finished bool
}

/* DBからtask一覧を取得する関数 */
// GetTasks: 引数無し, 戻り値がTask型のスライスとerror型
func GetTasks() ([]task, error) {
	// tasksを空のTask構造体のスライスとして定義する
	var tasks []Task

	// tasksにDBのタスク全てを代入する
	// この操作の可否をerrと定義する
	err := db.Find(&tasks).Error

	// tasksとerrを返す
	return tasks, err
}
