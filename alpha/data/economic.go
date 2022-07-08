package data

type CpiResponse struct {
	Name     string         `jsonlog:"name"`
	Interval string         `jsonlog:"interval"`
	Unit     string         `jsonlog:"unit"`
	Data     []CpiDataValue `jsonlog:"data"`
}

type CpiDataValue struct {
	Date  string `jsonlog:"date"`
	Value string `jsonlog:"value"`
}
