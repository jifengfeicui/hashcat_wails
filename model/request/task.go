package request

type CreateTask struct {
	FileID       uint   `json:"file_id"`     // 文件ID
	AttackMode   int    `json:"attack_mode"` // 攻击模式     0 字典 3 掩码 6 字典+掩码
	IncrementMin int    `json:"increment_min"`
	IncrementMax int    `json:"increment_max"`
	OutFilePath  string `json:"out_file_path"`
	DictFilePath string `json:"dict_file_path"`
	MaskType     int    `json:"mask_type"` //掩码类型   0 纯数字
	//         1 小写字母
	//         2 大写字母
	//         3 小写字母加数字，
	//	     4 大小写字母加数字
	//	     5 大小写字母+数字+可见特殊符号
}
