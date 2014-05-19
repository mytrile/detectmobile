package detectmobile

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

func Test_DetectMobile(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.New()

	m.Use(DetectMobile())
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko)")

	m.ServeHTTP(recorder, r)

	if recorder.HeaderMap["X-Mobile-Device"][0] != "true" {
		t.Error("iPhone user agent not recognized")
	}
}

func Test_checkAll(t *testing.T) {
	ua_header_iphone := "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko)"
	ua_header_blackberry := "Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en) AppleWebKit/534.11+ (KHTML, like Gecko)"

	accept_header := "application/vnd.wap.wmlscriptc"

	if checkAll(ua_header_iphone, "") != "true" {
		t.Error("iPhone user agent not recognized")
	}

	if checkAll(ua_header_blackberry, "") != "true" {
		t.Error("Blackberry user agent not recognized")
	}

	if checkAll("", accept_header) != "true" {
		t.Error("Mobile device with wap in accept header not recognized")
	}

	if checkAll("", "") != "false" {
		t.Error("Mobibe device detected incorrectly")
	}

}

func Test_checkForUaTargeteDevice(t *testing.T) {
	ua_iphone := "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko)"
	ua_desktop := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.75.14 (KHTML, like Gecko)"

	if result := checkForUaTargetedDevice(ua_iphone); result != true {
		t.Error("iPhone user agent not recognized")
	}

	if result := checkForUaTargetedDevice(ua_desktop); result == true {
		t.Error("Safari desktop recognized as mobile device")
	}
}

func Test_catchAll(t *testing.T) {
	ua_blackberry := "Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en) AppleWebKit/534.11+ (KHTML, like Gecko)"

	if result := catchAll(ua_blackberry); result != true {
		t.Error("BlackBerry user agent not recognized")
	}
}

func Test_chekcInAcceptHeader(t *testing.T) {
	accept_header := "application/vnd.wap.wmlscriptc"
	if result := chekcInAcceptHeader(accept_header); result != true {
		t.Error("Mobile device with wap in accept header not recognized")
	}
}
