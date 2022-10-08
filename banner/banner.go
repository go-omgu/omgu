package banner

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

type Banner struct {
	Name        string
	Description string
	VerFunc     VersionFunc
	Author      string
	Email       string
	Link        string
}

type VersionFunc func() string

var (
	errEmptyVar = errors.New("Variable is empty")
)

func VFConstant(v string) VersionFunc {
	return func() string { return v }
}

func VFFileHash() string {
	cmd, _ := os.Executable()
	f, _ := os.Open(cmd)
	buf, _ := io.ReadAll(f)
	s := sha256.Sum256(buf)
	return hex.EncodeToString(s[:4])
}

func (b *Banner) ShowBanner() error {
	return b.ShowBannerTo(os.Stderr)
}

func (b *Banner) ShowBannerTo(o io.Writer) error {
	if len(b.Name) == 0 {
		return fmt.Errorf("%w: %s", errEmptyVar, "Name")
	}
	ver := b.VerFunc()
	if len(ver) == 0 {
		ver = VFFileHash()
	}

	fmt.Fprintf(o, "%s - %s\n", b.Name, ver)

	if len(b.Description) > 0 {
		fmt.Fprintf(o, "    %s\n\n", b.Description)
	}

	if len(b.Author) > 0 {
		fmt.Fprintf(o, "@%s ", b.Author)
	}

	if len(b.Email) > 0 {
		fmt.Fprintf(o, "<%s>", b.Email)
	}

	if len(b.Link) > 0 {
		fmt.Fprintf(o, "\nVisit for more: %s", b.Link)
	}

	fmt.Fprintf(o, "\n\n")

	return nil
}
