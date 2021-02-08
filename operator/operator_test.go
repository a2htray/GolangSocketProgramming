package operator

import (
	"fmt"
	"testing"
)

func TestParseExpression(t *testing.T) {
	fmt.Println(ParseExpression("11+22"))
	fmt.Println(ParseExpression("11*22"))
	fmt.Println(ParseExpression("11/22"))
	fmt.Println(ParseExpression("11-22"))
	fmt.Println(ParseExpression("11022"))
}
