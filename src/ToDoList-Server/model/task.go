package model

import (
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
func GetTasks() ([]Task, error) {
	// tasksを空のTask構造体のスライスとして定義する
	var tasks []Task

	// db.Find: tasksにDBのタスク全てを代入する
	// .Error: 成功時にはnil, 失敗時にエラーの内容が返される。この操作の可否をerrと定義する
	err := db.Find(&tasks).Error

	// tasksとerrを返す
	return tasks, err
}
