package formatting

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Glamourize renders the given string with the Glamour
// markdown renderer.
func Glamourize(content string, width int) string {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(width),
	)
	out, err := r.Render(string(content))
	check(err)

	return out
}

// SplitTitleFromPost reads the frontmatter of a blogpost and
// parses the title, the returns it alongside the rest of the post
// without frontmatter.
func SplitTitleFromPost(page string) (string, string, error) {
	lines := strings.Split(page, "\n")
	var frontmatterClose int
	var title string

	// we're assuming the frontmatter starts at line 1
	for i, line := range lines[1:] {
		line = strings.TrimRight(line, "\r\n")
		if line == "---" {
			frontmatterClose = i + 1 // + 1 because we skipped first line
			break
		} else if strings.HasPrefix(line, "title:") {
			title = strings.TrimPrefix(line, "title:")
		}
	}

	if frontmatterClose == 0 {
		return "", "", errors.New("No frontmatter detected")
	}

	title = strings.Trim(title, " \"")
	body := strings.Join(lines[frontmatterClose+1:], "\n")
	return title, body, nil
}

// ExtractDateFromFilename parses a filename and extracts the date from
// it. This assumes the filename starts with the date, as Jekyll requires
// for blog posts. See: https://jekyllrb.com/docs/posts/
func ExtractDateFromFilename(filename string) (time.Time, error) {
	r := regexp.MustCompile(`(\d\d\d\d-\d\d-\d\d).*`)
	match := r.FindStringSubmatch(filename)
	if match == nil {
		return time.Now(), errors.New("filename didn't start with a date: " + filename)
	}

	d, err := time.Parse("2006-01-02", match[1])
	check(err)

	return d, nil
}

// PrettyDate formats a time string to look like
// "December 22nd", 2018.
func PrettyDate(date time.Time) string {
	day, err := strconv.Atoi(date.Format("2"))
	check(err)
	day = day % 10
	ending := ""
	switch day {
	case 1:
		ending = "st"
	case 2:
		ending = "nd"
	case 3:
		ending = "rd"
	default:
		ending = "th"
	}
	return date.Format(fmt.Sprintf("January 2%s, 2006", ending))
}
