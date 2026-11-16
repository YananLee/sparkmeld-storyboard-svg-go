package sparkmeld

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

var hexColor = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

type Options struct {
	Width     int
	CardColor string
}

// RenderStoryboard arranges 1 to 12 scene prompts into a planning SVG.
func RenderStoryboard(prompts []string, o Options) (string, error) {
	if len(prompts) < 1 || len(prompts) > 12 {
		return "", fmt.Errorf("provide 1 to 12 prompts")
	}
	for _, p := range prompts {
		if strings.TrimSpace(p) == "" {
			return "", fmt.Errorf("prompt must not be empty")
		}
	}
	if o.Width == 0 {
		o.Width = 900
	}
	if o.CardColor == "" {
		o.CardColor = "#e9edff"
	}
	if o.Width < 400 || !hexColor.MatchString(o.CardColor) {
		return "", fmt.Errorf("invalid options")
	}
	columns := len(prompts)
	if columns > 3 {
		columns = 3
	}
	rows := (len(prompts) + columns - 1) / columns
	const gap, cardHeight = 16, 180
	cardWidth := float64(o.Width-gap*(columns+1)) / float64(columns)
	height := rows*(cardHeight+gap) + gap
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="Storyboard with %d scenes"><rect width="100%%" height="100%%" fill="#ffffff"/>`, o.Width, height, o.Width, height, len(prompts))
	for i, p := range prompts {
		x := float64(gap) + float64(i%columns)*(cardWidth+gap)
		y := gap + (i/columns)*(cardHeight+gap)
		fmt.Fprintf(&b, `<rect x="%.1f" y="%d" width="%.1f" height="%d" rx="12" fill="%s"/><text x="%.1f" y="%d" font-family="Arial,sans-serif" font-size="18" font-weight="700" fill="#303f91">Scene %d</text>`, x, y, cardWidth, cardHeight, o.CardColor, x+18, y+36, i+1)
		words := strings.Fields(p)
		lines := []string{}
		current := ""
		for _, w := range words {
			if len([]rune(strings.TrimSpace(current+" "+w))) > 28 && current != "" {
				lines = append(lines, current)
				current = w
			} else {
				current = strings.TrimSpace(current + " " + w)
			}
		}
		if current != "" {
			lines = append(lines, current)
		}
		for j, line := range lines {
			if j >= 4 {
				break
			}
			r := []rune(line)
			if len(r) > 34 {
				line = string(r[:34])
			}
			fmt.Fprintf(&b, `<text x="%.1f" y="%d" font-family="Arial,sans-serif" font-size="16" fill="#22283a">%s</text>`, x+18, y+76+j*24, html.EscapeString(line))
		}
	}
	b.WriteString("</svg>")
	return b.String(), nil
}
