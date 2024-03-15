package ports

type UserManagement interface {
	// Change password
	UpdatePassword(newdata string) (string, error)
	// Update profile information (e.g., name, email)
	UpdateDATA(newdata string) (string, error)
	// User registration
	SignUp(email, password, role string) (string, error)
	// User login/logout
	SignIn(email, password string) (string, error)
}

type TaskManagement interface {
	// Create a new task
	CreateTask(taskData TaskData) (string, error)
	// Read/view tasks
	ReadTasks() ([]TaskData, error)
	// Update task details (e.g., title, description, due date, priority)
	UpdateTask(taskID string, newData TaskData) (string, error)
	// Mark tasks as completed/incomplete
	MarkTaskCompletion(taskID string, completed bool) error
	// Delete tasks
	DeleteTask(taskID string) error
	// Filter tasks by various criteria (e.g., due date, priority, status)
	FilterTasks(criteria TaskFilterCriteria) ([]TaskData, error)
}

type ListManagement interface {
	// Create a new list
	CreateList(listData ListData) (string, error)
	// Read/view lists
	ReadLists() ([]ListData, error)
	// Update list details (e.g., title, description)
	UpdateList(listID string, newData ListData) (string, error)
	// Delete lists
	DeleteList(listID string) error
}

type TaskData struct {
	Title       string
	Description string
	DueDate     string
	Priority    string
	Completed   bool
}

type ListData struct {
	Title       string
	Description string
}

type TaskFilterCriteria struct {
	DueDate  string
	Priority string
	Status   string
}
