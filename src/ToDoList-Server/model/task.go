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

/* AddTask: 引数がstring型のname, 戻り値がTaskのポインタとerror型 */
func AddTask(name string) (*Task, error) {
	// 新たなuuidを生成し、これをid、成否をerrとする
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	// ID, Name, Finishedにid, name, falseを代入したTask型のtaskを定義する
	task := Task{
		ID:       id,
		Name:     name,
		Finished: false,
	}

	// taskをDBのTaskテーブルに追加。成否をerrとする
	err = db.Create(&task).Error

	// taskのポインタとerrを返す
	return &task, db.Error
}

/* ChangeFinishedTask: 引数がuuid.UUID型のtaskID, 戻り値がerror型 */
func ChangeFinishedTask(taskID uuid.UUID) error {
	// DBのTaskテーブルから、taskIDと一致するidを探し、そのFinishedをtrueにする
	err := db.Model(&Task{}).Where("id = ?", taskID).Update("finished", true).Error
	return err
}

func DeleteTask(taskID uuid.UUID) error {
	// DBのTaskテーブルからtaskIDと一致するidを探し、そのタスクを削除する
	err := db.Where("id = ?", taskID).Delete(&Task{}).Error
	return err
}
