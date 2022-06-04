package builder

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	TITLE = "title"
	ERROR = "error"
)

type markDownText struct {
	content string
	cType 	string
}

func (t markDownText) parseType() string {
	matchTitle, _ := regexp.MatchString("^h[1-6]$", t.cType)
	if (matchTitle) {
		return TITLE
	}

	return ""
}

func (t markDownText) validationOfType() error {
	tpe := t.parseType()

	if tpe != "" {
		return nil
	}

	return errors.New("Type not supported at the moment, only support: " + strings.Join(getSupportedType(), ",") )
}

func NewMarkDowText(content string, ctype string) (MarkDownBuilder, error) {
	t := markDownText{
		cType: ctype,
		content: content,
	}

	err := t.validationOfType()

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t markDownText) Render() (string, error) {
	var renderString string

	switch t.parseType() {
    case TITLE:
		titleNumber, err := strconv.ParseInt(
			strings.ReplaceAll(
				t.cType,
				"h",
				"",
			),
			10,
			64,
		)

		if err != nil {
			return "", err
		}

        renderString = strings.Repeat(
			"#",
			int(titleNumber),
		) + " " + t.content
    }

	return renderString, nil
}

func getSupportedType() []string {
	return []string{"h1", "h2", "h3", "h4", "h5", "h6"}
}