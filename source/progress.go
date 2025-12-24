package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"
)

type progressWriter struct {
	label      string
	total      int64
	written    int64
	lastRender time.Time
	manager    *progressManager
	line       int
}

type progressManager struct {
	out     io.Writer
	mu      sync.Mutex
	lines   int
	active  int
	enabled bool
}

var defaultProgressManager = newProgressManager(os.Stdout)

func newProgressWriter(label string, total int64, out io.Writer) *progressWriter {
	manager := defaultProgressManager
	if out != defaultProgressManager.out {
		manager = newProgressManager(out)
	}

	return &progressWriter{
		label:      label,
		total:      total,
		lastRender: time.Now().Add(-time.Second),
		manager:    manager,
		line:       manager.register(),
	}
}

func (p *progressWriter) Write(data []byte) (int, error) {
	n := len(data)
	p.written += int64(n)

	now := time.Now()
	if now.Sub(p.lastRender) >= 200*time.Millisecond || (p.total > 0 && p.written >= p.total) {
		p.lastRender = now
		p.render()
	}

	return n, nil
}

func (p *progressWriter) Finish() {
	p.render()
	p.manager.finish()
}

func (p *progressWriter) render() {
	p.manager.render(p.line, p.label, p.progressText())
}

func (p *progressWriter) progressText() string {
	if p.total > 0 {
		percent := float64(p.written) / float64(p.total) * 100
		if percent > 100 {
			percent = 100
		}
		return fmt.Sprintf("%3.0f%% (%s / %s)", percent, humanBytes(p.written), humanBytes(p.total))
	}

	return fmt.Sprintf("%s downloaded", humanBytes(p.written))
}

func newProgressManager(out io.Writer) *progressManager {
	return &progressManager{
		out:     out,
		enabled: isTerminal(out),
	}
}

func (m *progressManager) register() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	line := m.lines
	m.lines++
	m.active++
	if m.enabled && line > 0 {
		fmt.Fprintln(m.out)
	}

	return line
}

func (m *progressManager) render(line int, label, text string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.enabled {
		fmt.Fprintf(m.out, "%s: %s\n", label, text)
		return
	}

	linesFromBottom := m.lines - line - 1
	if linesFromBottom > 0 {
		fmt.Fprintf(m.out, "\x1b[%dA", linesFromBottom)
	}
	fmt.Fprintf(m.out, "\r\x1b[2K%s: %s", label, text)
	if linesFromBottom > 0 {
		fmt.Fprintf(m.out, "\x1b[%dB", linesFromBottom)
	}
}

func (m *progressManager) finish() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.active > 0 {
		m.active--
	}

	if m.enabled && m.active == 0 {
		fmt.Fprintln(m.out)
	}
}

func isTerminal(out io.Writer) bool {
	file, ok := out.(*os.File)
	if !ok {
		return false
	}

	info, err := file.Stat()
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}

func humanBytes(n int64) string {
	if n < 1024 {
		return fmt.Sprintf("%d B", n)
	}

	const unit = 1024.0
	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB"}

	exponent := math.Floor(math.Log(float64(n)) / math.Log(unit))
	if exponent >= float64(len(suffixes)) {
		exponent = float64(len(suffixes) - 1)
	}

	scaled := float64(n) / math.Pow(unit, exponent)
	return fmt.Sprintf("%.1f %s", scaled, suffixes[int(exponent)])
}
