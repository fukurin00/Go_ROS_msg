package ros_msg

import "math"

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (p Point) Distance(o Point) float64 {
	return math.Hypot(p.X-o.X, p.Y-o.Y)
}

type Quaternion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	W float64 `json:"w"`
}

type Pose struct {
	Position    Point      `json:"position"`
	Orientation Quaternion `json:"orientation"`
}

type PoseStamped struct {
	Header Header `json:"header"`
	Pose   Pose   `json:"pose"`
}
