package model

type SourceHashFile struct {
	FileName string         `json:"file_name"`
	FileType string         `json:"file_type"`
	FilePath string         `json:"file_path"`
	FileMD5  string         `json:"file_md5" gorm:"unique"`
	Tasks    []*HashCatTask `gorm:"foreignKey:FileID" json:"tasks"`
	MODEL
}

type HashCatTask struct {
	FileID           uint            `json:"file_id"`
	File             *SourceHashFile `gorm:"foreignKey:FileID" json:"file"` // 指定外键
	GuessBasePercent float64         `json:"guess_base_percent"`
	GuessModPercent  float64         `json:"guess_mod_percent"`
	CurrentProgress  int             `json:"current_progress"`
	TotalProgress    int             `json:"total_progress"`
	Status           TaskStatus      `json:"status"`
	CMD              string          `json:"cmd"`
	OutFilePath      string          `json:"out_file_path"`
	Result           string          `json:"result"`
	AttackMode       int             `json:"attack_mode"` // 攻击模式     0 字典 3 掩码 6 字典+掩码
	IncrementMin     int             `json:"increment_min"`
	IncrementMax     int             `json:"increment_max"`
	MaskType         int             `json:"mask_type"` //掩码类型   0 纯数字
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
