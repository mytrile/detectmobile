package detectmobile

import (
	"net/http"
	"regexp"

	"github.com/go-martini/martini"
)

const (
	X_HEADER = "X-MOBILE-DEVICE"
)

func DetectMobile() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		ua_header := r.Header.Get("User-Agent")
		accept_header := r.Header.Get("Accept")
		mobileDevice := checkAll(ua_header, accept_header)

		w.Header().Set(X_HEADER, mobileDevice)
		c.Map(mobileDevice)
	}
}

func checkAll(ua_header, accept_header string) string {
	switch {
	case checkForUaTargetedDevice(ua_header) == true:
		return "true"
	case chekcInAcceptHeader(accept_header) == true:
		return "true"
	case catchAll(ua_header) == true:
		return "true"
	default:
		return "false"
	}
}

func checkForUaTargetedDevice(header string) bool {
	match, _ := regexp.MatchString(`(?i)iphone|android|ipod|ipad`, header)
	return match
}

func catchAll(header string) bool {
	catch_all_string := `(?i)palm|blackberry|nokia|phone|midp|mobi|symbian|chtml|ericsson|minimo|audiovox|motorola|samsung|telit|upg1|windows ce|ucweb|astel|plucker|x320|x240|j2me|sgh|portable|sprint|docomo|kddi|softbank|android|mmp|pdxgw|netfront|xiino|vodafone|portalmmm|sagem|mot-|sie-|ipod|up\\.b|webos|amoi|novarra|cdm|alcatel|pocket|ipad|iphone|mobileexplorer|mobile`
	match, _ := regexp.MatchString(catch_all_string, header)
	return match
}

func chekcInAcceptHeader(header string) bool {
	match, _ := regexp.MatchString(`(?)vnd\.wap`, header)
	return match
}
