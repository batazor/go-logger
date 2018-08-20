package influxdb

type StateRequest struct {
	measurement string
	function    string
	fields      string
	where       string
}
