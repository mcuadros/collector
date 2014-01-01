package format

import . "launchpad.net/gocheck"
import (
	"fmt"
)

type FormatJSONSuite struct{}

var _ = Suite(&FormatJSONSuite{})

func (s *FormatJSONSuite) TestGetRecord(c *C) {
	config := JSONConfig{}

	format := NewJSON(&config)

	record := format.Parse("{\"foo\":{\"foo\":\"bar\",\"bar\":\"qux\"}}")
	fmt.Println(record)

	foo := record["foo"].(map[string]interface{})
	c.Check(foo["foo"], Equals, "bar")
}