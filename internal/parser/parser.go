package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/matsuev/klsh-email-sender/internal/config"
)

// TableParser struct
type TableParser struct {
	Keys  []string
	fh    *os.File
	scan  *bufio.Scanner
	count int
}

// Create function
func Create(cfg *config.AppConfig) (*TableParser, error) {
	fh, err := os.Open(cfg.DataFile)
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(fh)
	scan.Split(bufio.ScanLines)

	tp := &TableParser{
		Keys:  make([]string, 0),
		fh:    fh,
		scan:  scan,
		count: 0,
	}

	if err = tp.readKeys(); err != nil {
		return nil, err
	}

	return tp, nil
}

// Close function
func (p *TableParser) Close() error {
	return p.fh.Close()
}

// readKeys function
func (p *TableParser) readKeys() error {
	if !p.scan.Scan() {
		return io.EOF
	}

	p.Keys = strings.Split(strings.TrimSpace(p.scan.Text()), ";")
	p.count++

	return nil
}

// Scan function
func (p *TableParser) Scan() bool {
	return p.scan.Scan()
}

// GetLine function
func (p *TableParser) GetLine() (map[string]string, error) {
	csvLine := strings.TrimSpace(p.scan.Text())
	p.count++

	csvParts := strings.Split(csvLine, ";")

	if len(p.Keys) != len(csvParts) {
		return nil, fmt.Errorf("parse CSV file error with line number: %d, and bytes: %v", p.count, []byte(csvLine))
	}

	result := make(map[string]string)

	for idx, key := range p.Keys {
		result[key] = csvParts[idx]
	}

	return result, nil
}
