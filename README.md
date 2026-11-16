# sparkmeld-storyboard-svg-go

A small Go package for generating SVG graphics associated with [SparkMeld](https://sparkmeld.com/). It runs locally and does not call the site's AI or export service.

```sh
go get github.com/YananLee/sparkmeld-storyboard-svg-go@v0.1.0
```

```go
import "github.com/YananLee/sparkmeld-storyboard-svg-go"

svg, err := sparkmeld.RenderStoryboard([]string{"Opening scene", "Closing scene"}, sparkmeld.Options{})
if err != nil { panic(err) }
_ = svg // write to an .svg file or embed in a page
```

Licensed under MIT.
