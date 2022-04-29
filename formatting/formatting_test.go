package formatting

import (
	"strings"
	"testing"
	"time"
)

func testCheck(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
func TestSplitTitleFromPost(t *testing.T) {
	post := `---
title: "yolo"
category: whatever
layout: post
---

# heading

first paragraph

## subheading

whatever`

	title, body, err := SplitTitleFromPost(post)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(title)
	t.Log(body)

	if title != "yolo" {
		t.Fatalf("Title was %q", title)
	}

	// check that closing frontmatter dashes were cut
	if strings.Split(body, "\n")[0] == "---" {
		t.Fatal("First line was frontmatter dashes (---)")
	}

}

func TestExtractDateFromFilename(t *testing.T) {
	goodTitle := "2022-03-09-on-reuse-and-waste.markdown"
	extractedDate, err := ExtractDateFromFilename(goodTitle)
	testCheck(t, err)
	referenceDate, err := time.Parse("2006-01-02", "2022-03-09")
	testCheck(t, err)
	if extractedDate != referenceDate {
		t.Fatalf("Dates didn't match!\nextracted: %q, reference: %q", extractedDate, referenceDate)
	}

	badTitle := "2004-Jan-09-on-reuse-and-waste.markdown"
	_, err = ExtractDateFromFilename(badTitle)
	if err == nil {
		t.Fatal("No error reported even though the date was in the wrong format")
	}

}

// TODO: test matrix for other dates
func TestPrettyDate(t *testing.T) {
	referenceDate, err := time.Parse("2006-01-02", "2022-03-09")
	testCheck(t, err)

	pretty := PrettyDate(referenceDate)
	if pretty != "March 9th, 2022" {
		t.Fatal("Expected th, got: " + pretty)
	}
}
