package router

import (
	"ToDoList-Server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/* Task一覧をjsonで返却する関数 */
// GetTaskHandler: 引数がecho.Context型のc, 戻り値はerror型
func GetTaskHandler(c echo.Context) error {
	// model(package)の関数GetTasksを実行し、戻り値をtasks, errと定義する
	tasks, err := model.GetTasks()

	// エラーが起きた場合、StatusBadRequestを返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// エラーが起きなかった場合、StatusOKとtasksを返す
	return c.JSON(http.StatusOK, tasks)
}

/* POSTメソッド用の構造体 */
// RequestTask型は、文字列のNameをパラメータとして持つ
type RequestTask struct {
	// json:"name" -> jsonデータを代入するための識別子
	Name string `json:"name"`
}

/* AddTaskHandler: 引数がecho.Context型, 戻り値はerror型 */
func AddTaskHandler(c echo.Context) error {
	// 空のRequestTaskであるreqを定義
	var req RequestTask

	// bodyのjsonファイルをbind
	err := c.Bind(&req)
	// エラーハンドリング
	// StatusBadRequestを返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// 空のmodel(package)のTaskである、taskを定義
	var task *model.Task

	// model(package)のAddTask関数を実行し、戻り値をtask, errと定義
	task, err = model.AddTask(req.Name)
	// エラーハンドリング
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// エラーでない場合、StatusOKと追加されたtaskを返す
	return c.JSON(http.StatusOK, task)
}

func ChangeFinishedTaskHandler(c echo.Context) error {
	// taskIDのパスパラメータ(string型)を取得し、uuid型に変換
	// その値をtaskID, 成否をerrとする
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// ChangeFinishedTaskを実行
	// 戻り値をerrに代入する(errを更新)
	err = model.ChangeFinishedTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
