package examples

import (
	"strings"
)

func GetExamples() interface{} {
	return []struct {
		Name     string `json:"name"`
		Document string `json:"document"`
		Expr     string `json:"expr"`
	}{
		{
			Name: "array return",
			Document: `<html><body>
<div class="item">
    <div class="title">title A</div>
    <div class="author">author A</div>
</div>
<div class="item">
    <div class="title">title B</div>
    <div class="author">author B</div>
</div>
<div class="item">
    <div class="title">title C</div>
    <div class="author">author C</div>
</div>
</body></html>`,
			Expr: strings.Join([]string{
				"item `css(\".item\")` [",
				"    {",
				"        title `css(\".title\")`",
				"        author `css(\".author\")`",
				"    }",
				"]",
			}, "\r\n"),
		},
		{
			Name: "string return",
			Document: `<html><body>
    <div class="title">Article</div>
</body></html>
                `,
			Expr: strings.Join([]string{
				"title `css(\".title\")`",
			}, "\r\n"),
		},
		{
			Name: "mixed format",
			Document: `<html><body>
    <div class="title">Article</div>
    <div class="tags">
        <div class="tag">tag0</div>
        <div class="tag">tag1</div>
        <div class="tag">tag2</div>
    </div>
</body></html>`,
			Expr: strings.Join([]string{
				"{",
				"    title `css(\".title\")`",
				"    tags `css(\".tag\")` [",
				"        tag `text()`",
				"    ]",
				"}",
			}, "\r\n"),
		},
		{
			Name: "2-dimensional array",
			Document: `<html><body>
    <div class="items">
        <div class="item">
            <div class="pos">1.0</div>
            <div class="pos">1.1</div>
            <div class="pos">1.2</div>
        </div>
        <div class="item">
            <div class="pos">2.0</div>
            <div class="pos">2.1</div>
            <div class="pos">2.2</div>
        </div>
        <div class="item">
            <div class="pos">3.0</div>
            <div class="pos">3.1</div>
            <div class="pos">3.2</div>
        </div>
    </div>
</body></html>`,
			Expr: strings.Join([]string{
				"items `css(\".item\")` [",
				"    item `css(\".pos\")` [",
				"        pos `text()`",
				"    ]",
				"]",
			}, "\r\n"),
		},
		{
			Name: "1-dimensional array",
			Document: `<html><body>
    <div class="items">
        <div class="item">
            <div class="pos">1.0</div>
            <div class="pos">1.1</div>
            <div class="pos">1.2</div>
        </div>
        <div class="item">
            <div class="pos">2.0</div>
            <div class="pos">2.1</div>
            <div class="pos">2.2</div>
        </div>
        <div class="item">
            <div class="pos">3.0</div>
            <div class="pos">3.1</div>
            <div class="pos">3.2</div>
        </div>
    </div>
</body></html>`,
			Expr: strings.Join([]string{
				"item `css(\".pos\")` [",
				"    pos `text()`",
				"]",
			}, "\r\n"),
		},
		{
			Name: "obejct array",
			Document: `<html>
    <body>
        <a href="01.html">Page 1</a>
        <a href="02.html">Page 2</a>
        <a href="03.html">Page 3</a>
    </body>
</html>`,
			Expr: strings.Join([]string{
				"{",
				"    anchor `css(\"a\")` [",
				"        content `text()`",
				"    ]",
				"}",
			}, "\r\n"),
		},
		{
			Name: "object array2",
			Document: `<html>
    <body>
        <a href="01.html">Page 1</a>
        <a href="02.html">Page 2</a>
        <a href="03.html">Page 3</a>
    </body>
</html>`,
			Expr: strings.Join([]string{
				"{",
				"    anchor `css(\"a\")` [",
				"        {",
				"            title `text()`",
				"        }",
				"    ]",
				"}",
			}, "\r\n"),
		},
		{
			Name: "single line",
			Document: `<html>
    <body>
        <a href="01.html">Page 1</a>
        <a href="02.html">Page 2</a>
        <a href="03.html">Page 3</a>
    </body>
</html>`,
			Expr: strings.Join([]string{
				"anchor `css(\"a\")` [{title `text()`;  url `attr(\"href\")`;}]",
			}, "\r\n"),
		},
		{
			Name:     "anchor",
			Document: `Please try to enter any URL address in the input box above and click Download button.`,
			Expr: strings.Join([]string{
				"a `css(\"a\")` [{",
				"    title `text();trim()`",
				"    url  `attr(\"href\")`",
				"}]",
			}, "\r\n"),
		},
		{
			Name: "stunning!",
			Document: `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Test Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script>
        var data = {
            "name": {
                "first": "Tom",
                "last": "Anderson"
            },
            "age": 37,
            "children": ["Sara", "Alex", "Jack"],
            "fav.movie": "Deer Hunter",
            "friends": [{
                    "first": "Dale",
                    "last": "Murphy",
                    "age": 44
                },
                {
                    "first": "Roger",
                    "last": "Craig",
                    "age": 68
                },
                {
                    "first": "Jane",
                    "last": "Murphy",
                    "age": 47
                }
            ]
        }
    </script>
</head>
<body>
    <library>
        <!-- Great book. -->
        <book id="b0836217462" available="true">
            <isbn>0836217462</isbn>
            <title lang="en">Being a Dog Is a Full-Time Job</title>
            <quote>I'd dog paddle the deepest ocean.</quote>
            <author id="CMS">
                <?echo "go rocks"?>
                    <name>Charles M Schulz</name>
                    <born>1922-11-26</born>
                    <dead>2000-02-12</dead>
            </author>
            <character id="PP">
                <name>Peppermint Patty</name>
                <born>1966-08-22</born>
                <qualification>bold, brash and tomboyish</qualification>
            </character>
            <character id="Snoopy">
                <name>Snoopy</name>
                <born>1950-10-04</born>
                <qualification>extroverted beagle</qualification>
            </character>
        </book>
    </library>
</body>
</html>`,
			Expr: strings.Join([]string{
				"{",
				"    __json__ `regex(\"var data = ([\\s\\S]*?)\\s+</script>\")`",
				"    __firstname__ `link(\"__json__\");json(\"name.first\")`",
				"    __lastname__ `link(\"__json__\");json(\"name.last\")`",
				"    author `template(\"{$__firstname__} {$__lastname__}\")`",
				"    bookname `css(\"library title\")`",
				"    bookID `css(\"book\");attr(\"id\")`",
				"    character `xpath(\"//character\")` [{",
				"        born `css(\"born\")`",
				"        name  `xpath(\"name\")`",
				"        qualification `regex(\"qualification>(.*?)<\")`",
				"    }]",
				"    friends `link(\"__json__\");json(\"friends\")` [{",
				"        first `json(\"first\")`",
				"        last `json(\"last\")`",
				"        age `json(\"age\")`",
				"    }]",
				"}",
			}, "\r\n"),
		},
	}
}
