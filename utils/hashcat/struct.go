package hashcat

type HashcatResponse struct {
	Session         string `json:"session"`
	Guess           Guess  `json:"guess"`
	Status          int    `json:"status"`
	Target          string `json:"target"`
	Progress        []int  `json:"progress"`
	RestorePoint    int    `json:"restore_point"`
	RecoveredHashes []int  `json:"recovered_hashes"`
	RecoveredSalts  []int  `json:"recovered_salts"`
	Rejected        int    `json:"rejected"`
	//Devices         []Devices `json:"devices"`
	TimeStart     int `json:"time_start"`
	EstimatedStop int `json:"estimated_stop"`
}
type Guess struct {
	GuessBase        string      `json:"guess_base"`
	GuessBaseCount   int         `json:"guess_base_count"`
	GuessBaseOffset  int         `json:"guess_base_offset"`
	GuessBasePercent float64     `json:"guess_base_percent"`
	GuessMaskLength  int         `json:"guess_mask_length"`
	GuessMod         interface{} `json:"guess_mod"`
	GuessModCount    int         `json:"guess_mod_count"`
	GuessModOffset   int         `json:"guess_mod_offset"`
	GuessModPercent  float64     `json:"guess_mod_percent"`
	GuessMode        int         `json:"guess_mode"`
}
type Devices struct {
	DeviceID   int    `json:"device_id"`
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
	Speed      int    `json:"speed"`
	Temp       int    `json:"temp"`
	Util       int    `json:"util"`
}
