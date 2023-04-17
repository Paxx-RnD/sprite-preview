package types

import (
	"flag"
	"fmt"
)

type Flags struct {
	Input     string
	Prefix    string
	Vtt       string
	Rows      int
	Columns   int
	Frequency int
	Width     int
	Height    int
	Help      bool
}

func (f *Flags) Set() error {
	input := flag.String("i", "", "path of the input video")
	prefix := flag.String("prefix", "", "prefix sprites")
	vtt := flag.String("vtt", "", "specify the prefix for the vtt file")
	frequency := flag.Int("f", 3, "extract frames every n seconds")
	rows := flag.Int("row", 10, "how many rows")
	columns := flag.Int("col", 10, "how many columns")
	width := flag.Int("w", 160, "frame width")
	height := flag.Int("h", 90, "frame height")
	help := flag.Bool("help", false, "shows help")

	flag.Parse()

	f.Input = *input
	f.Prefix = *prefix
	f.Vtt = *vtt
	f.Frequency = *frequency
	f.Rows = *rows
	f.Columns = *columns
	f.Width = *width
	f.Height = *height
	f.Help = *help

	if f.Input == "" {
		return fmt.Errorf("need to specify input")
	}
	if f.Prefix == "" {
		return fmt.Errorf("need to specify prefix")
	}
	if f.Frequency <= 0 {
		return fmt.Errorf("need to specify valid target")
	}

	return nil
}
