package photo

const (
	Normal  = "常规尺寸" // 开始生成枚举值, 默认为0
	Student = "学生证类" //学生
	Exam    = "考试证类" //考试
	ID      = "证件照"  //证件照
	Signed  = "签证照类" //签证
	Other   = "其他尺寸" // 其他尺寸
)

//证件照分类
type PhotoGroup struct {
	GroupName string  `json:"group_name"`
	Photos    []Photo `json:"photos"`
}

//证件照基本信息
type Photo struct {
	PhotoName string `json:"photo_name"`
	//附加信息
	PhotoDes      string `json:"photo_des"`
	PhotoFileSize string `json:"photo_file_size"`
	WidthMm       int    `json:"width_mm"`
	HeightMm      int    `json:"height_mm"`
	WidthPx       int    `json:"width_px"`
	HeightPx      int    `json:"height_px"`
}
