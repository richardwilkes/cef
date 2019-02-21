package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

var (
	fileLines    = make(map[string][]string)
	linePosRegex = regexp.MustCompile(`<([^:]+):(\d+)(?::\d+)?, ([^:]+):(\d+)(?::\d+)?> ([^:]+):(\d+)`)
)

type position struct {
	Src       string
	LineStart int
	LineEnd   int
}

func (p *position) update(line string) {
	if result := linePosRegex.FindAllStringSubmatch(line, -1); len(result) > 0 {
		pieces := result[0]
		src := -1
		for i := 5; i > 0; i -= 2 {
			if pieces[i] != "line" && pieces[i] != "col" {
				p.setSrc(pieces[i])
				src = i
				break
			}
		}
		min := 100000000
		max := -1
		for i := 1; i < 6; i += 2 {
			if pieces[i] != "col" && (pieces[i] == "line" || src == i) {
				v := mustAtoi(pieces[i+1])
				if min > v {
					min = v
				}
				if max < v {
					max = v
				}
			}
		}
		if max != -1 {
			p.LineStart = min
			p.LineEnd = max
		}
	}
}

func (p *position) setSrc(src string) {
	p.Src = strings.TrimPrefix(src, "./")
}

func mustAtoi(in string) int {
	value, err := strconv.Atoi(in)
	jot.FatalIfErr(err)
	return value
}

func (p position) FileLines() []string {
	lines, ok := fileLines[p.Src]
	if !ok {
		path := p.Src
		if !filepath.IsAbs(path) {
			path = filepath.Join(cefBaseDir, path)
		}
		f, err := os.Open(path)
		if err != nil {
			if !filepath.IsAbs(p.Src) {
				jot.Error(errs.NewWithCause(path, err))
			}
		} else {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			jot.FatalIfErr(f.Close())
		}
		fileLines[p.Src] = lines
	}
	return lines
}

func (p position) Comment() string {
	lines := p.FileLines()
	if p.LineStart > len(lines) {
		return ""
	}
	var commentLines []string
	i := p.LineStart - 2 // Account for 1-based line numbers
	for i >= 0 {
		line := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(line, "//") {
			break
		}
		if !strings.HasPrefix(line, "///") {
			commentLines = append(commentLines, line)
		}
		i--
	}
	if len(commentLines) == 0 {
		return ""
	}
	var buffer strings.Builder
	for i = len(commentLines) - 1; i >= 0; i-- {
		buffer.WriteString(commentLines[i])
		if i != 0 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

// Text returns the text section from the requested line. All params to this
// function are 1-based.
func (p position) Text(line, startCol, endCol int) string {
	if line < 1 {
		return ""
	}
	lines := p.FileLines()
	if line > len(lines) {
		return ""
	}
	text := lines[line-1]
	if text == "" {
		return ""
	}
	if startCol < 1 {
		startCol = 1
	}
	if endCol > len(text) {
		endCol = len(text)
	}
	if endCol < startCol {
		endCol = startCol
	}
	return text[startCol-1 : endCol]
}
