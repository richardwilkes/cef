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
	linePosRegex = regexp.MustCompile(`(.*):(\d+):(\d+), (col:(\d+)|line:(\d+):(\d+))`)
)

type position struct {
	Src       string
	LineStart int
	ColStart  int
	LineEnd   int
	ColEnd    int
}

func (p *position) update(line string) {
	left := strings.Index(line, " <")
	right := strings.Index(line, "> ")
	if left != -1 && right > left {
		if result := linePosRegex.FindAllStringSubmatch(line[left+2:right], -1); len(result) > 0 {
			if result[0][1] != "line" {
				p.Src = result[0][1]
				if strings.HasPrefix(p.Src, "./") {
					p.Src = p.Src[2:]
				}
			}
			var err error
			p.LineStart, err = strconv.Atoi(result[0][2])
			jot.FatalIfErr(err)
			p.ColStart, err = strconv.Atoi(result[0][3])
			jot.FatalIfErr(err)
			if result[0][5] == "" {
				p.LineEnd, err = strconv.Atoi(result[0][6])
				jot.FatalIfErr(err)
				p.ColEnd, err = strconv.Atoi(result[0][7])
				jot.FatalIfErr(err)
			} else {
				p.LineEnd = p.LineStart
				p.ColEnd, err = strconv.Atoi(result[0][5])
				jot.FatalIfErr(err)
			}
		}
	}
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
	if startCol < 1 {
		startCol = 1
	}
	if endCol > len(lines[line-1]) {
		endCol = len(lines[line-1])
	}
	if endCol < startCol {
		endCol = startCol
	}
	return lines[line-1][startCol-1 : endCol]
}
