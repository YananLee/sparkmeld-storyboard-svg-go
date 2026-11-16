package sparkmeld

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	svg, err := RenderStoryboard([]string{`<Scene & one>`}, Options{})
	if err != nil || !strings.Contains(svg, `&lt;Scene &amp; one&gt;`) {
		t.Fatalf("escape/render failed: %v %s", err, svg)
	}
	if _, err := RenderStoryboard(nil, Options{}); err == nil {
		t.Fatal("accepted no prompts")
	}
}
