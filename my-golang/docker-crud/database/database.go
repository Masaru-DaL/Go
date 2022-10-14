type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Task string `json:"task"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Todo{})

	return nil
}
