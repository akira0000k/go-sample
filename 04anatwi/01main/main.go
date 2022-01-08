package main
 
import (
	"github.com/ChimeraCoder/anaconda"
)
 
func main() {
	anaconda.NewTwitterApiWithCredentials(
		"access-token",
		"access-token-secret",
		"consumer-key",
		"consumer-secret",
	)
}
