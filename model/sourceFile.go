package model

type SourceHashFile struct {
	FileName string         `json:"file_name"`
	FileType string         `json:"file_type"`
	FilePath string         `json:"file_path"`
	Tasks    []*HashCatTask `gorm:"foreignKey:FileID"`
	MODEL
}

type HashCatTask struct {
	FileID           uint       `json:"file_id"`
	Session          string     `gorm:"uniqueIndex" json:"session"`
	GuessBasePercent float64    `json:"guess_base_percent"`
	GuessModPercent  float64    `json:"guess_mod_percent"`
	CurrentProgress  int        `json:"current_progress"`
	TotalProgress    int        `json:"total_progress"`
	Status           TaskStatus `json:"status"`
	CMD              string     `json:"cmd"`
	OutFilePath      string     `json:"out_file_path"`
	Result           string     `json:"result"`
	MODEL
}

type TaskStatus int

const (
	TaskStatusInit TaskStatus = iota + 1
	TaskStatusRunning
	TaskStatusError
	TaskStatusStop
	TaskStatusFinish
)
