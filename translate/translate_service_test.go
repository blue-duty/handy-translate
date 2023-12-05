package translate

import (
	"fmt"
	"handy-translate/config"
	"handy-translate/translate/deepl"
	"testing"

	"github.com/OwO-Network/gdeeplx"
)

func TestGetTransalteWay(t *testing.T) {
	result, err := gdeeplx.Translate("hello", "EN", "ZH", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)
}

func TestGetTransalteWayList(t *testing.T) {
	config.Init("handy-translate")
	v := GetTransalteWay(deepl.Way)
	s, err := v.PostQuery("Software\r\nAnalytics\r\nArchiving and Digital Preservation (DP)\r\nAutomation\r\nBackup\r\nBlogging Platforms\r\nBooking and Scheduling", "", "ZH")
	if err != nil {
		t.Fatal(err)

	}
	fmt.Println(s)

}
