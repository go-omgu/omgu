package banner

import "testing"

func TestShowBanner(t *testing.T) {
	b := new(Banner)
	b.Name = "omgu"
	b.Description = "Common lightweight utils fo Golang."
	b.VerFunc = VFConstant("dev")
	// b.Author = "xbol0"
	// b.Email = "xbolo@duck.com"
	b.Link = "https://github.com/go-omgu/omgu"
	b.ShowBanner()
}
