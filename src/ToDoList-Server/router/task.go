package router

import (
	"ToDoList-Server/model"
	"github.com/labstack/echo/v4"
	"net/http"
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
