package ros_msg

type OccupancyGrid struct{
	Header Header `json:"header"`
	Info MapMetaData `json:"info"`
	Data []int8 `json:"data"`
}

type MapMetaData struct{
	Time TimeStamp `json:"time"`
	Resolution float32 `json:"resolution"`
	Width uint32 `json:"width"`
	Height uint32 `json:"height"`
	Origin Pose `json:"origin"`
}

type Path struct {
	Header Header    `json:"header"`
	Poses  []PoseStamped `json:"poses"`
}