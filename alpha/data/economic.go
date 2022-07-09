package data

type EconomicResponse struct {
	Name     string          `jsonlog:"name"`
	Interval string          `jsonlog:"interval"`
	Unit     string          `jsonlog:"unit"`
	Data     []EconomicValue `jsonlog:"data"`
}

type EconomicValue struct {
	Date  string `jsonlog:"date"`
	Value string `jsonlog:"value"`
}
