# Banner

Print cli base info.

## Usage

```go
b := &banner.Banner{
	  Name: "CLI name", // required
		Description: "", // optional
		VerFunc: banner.VFConstant("1.0"), // version generater, optional, defaults to filehash
		Author: "", // optional
		Email: "", // optional
		Link: "", // optional
}

// Default prints to STDERR
b.ShowBanner()

// Or specific to other
b.ShowBannerTo(os.Stdout)
```
