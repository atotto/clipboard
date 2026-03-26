package clipboard

import "os"

func writeOsc52(text string) error {
	_, err := os.Stderr.WriteString("\033]52;c;" + text + "\a")
	return err
}
