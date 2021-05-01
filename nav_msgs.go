package ros_msg

type OccupancyGrid struct {
	Header Header      `json:"header"`
	Info   MapMetaData `json:"info"`
	Data   []int8      `json:"data"`
}

type MapMetaData struct {
	Time       TimeStamp `json:"time"`
	Resolution float32   `json:"resolution"`
	Width      uint32    `json:"width"`
	Height     uint32    `json:"height"`
	Origin     Pose      `json:"origin"`
}

type Path struct {
	Header Header        `json:"header"`
	Poses  []PoseStamped `json:"poses"`
}

type Odometry struct {
	Header         Header              `json:"header"`
	Child_Frame_ID string              `json:"child_frame_id"`
	Pose           PoseWithCovariance  `json:"pose"`
	Twist          TwistWithCovariance `json:"twist"`
}

type PoseWithCovariance struct {
	Pose       Pose        `json:"pose"`
	Covariance [36]float64 `json:"covariance"`
}

type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
type Twist struct {
	Linear  Vector3 `json:"linear"`
	Angular Vector3 `json:"angular"`
}

type TwistWithCovariance struct {
	Twist      Twist       `json:"twist"`
	Covariance [36]float64 `json:"covariance"`
}
