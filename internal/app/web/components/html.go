package components

import (
	"fmt"
	"html/template"
)

type Html struct {
	Name    string
	Title   string
	UseIcon bool
	Page    Component
	css     template.CSS
	js      template.JS
}

func (c *Html) SetCss(css template.CSS) {
	c.css = css
}

func (c *Html) SetJs(js template.JS) {
	c.js = js
}

func (c *Html) GetContent() template.HTML {
	iconLink := ""
	if c.UseIcon {
		iconLink = `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css">`
	}
	return template.HTML(fmt.Sprintf(`<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge">
    <meta name="viewport" content="width=device-width,initial-scale=1,shrink-to-fit=no,user-scalable=no">
    <meta name="theme-color" content="#000000">
    <meta name="version" content="0.1">
    <link rel="shortcut icon" href="favicon.ico">
	<title>%s</title>
    <meta name="description" content="">
    <meta name="keywords" content="">
	%s
    <style>
		a {
			text-decoration: none;
			color: #003C97;
		}
        #app-container {
            color: #333;
            padding: 15px 25px;
        }
        .title {
            font-size: 21px;
            border-bottom: 1px solid #e8ecf1;
            padding-bottom: 15px;
        }
		.content {
	    	display: flex;
			flex-direction: row;
			flex-wrap: wrap;
		}
        .content h2 {
            font-size: 18px;
            margin-top: 25px;
            margin-bottom: 15px;
        }
        .link .link-block {
            background-color: #f7f7f7;
			padding: 10px;
        }
        .link a {
            width: 305px;
            height: 60px;
            border: 1px solid #ececed;
            background: #f7f7f7;
            display: flex;
            margin: 0 auto;
            align-items: center;
            justify-content: center;
            text-decoration: none;
            color: #4a4a5a;
        }
        .link-icon {
            width: 30px;
            height: 30px;
        }
		.app {
			display: flex;
			flex-direction: column;
			align-items: center;
			width: 50%%;
			height: 210px;
			justify-content: center;
		}
		.app a {
			display: flex;
			flex-direction: column;
			align-items: center;
		}
		.app span {
			margin-top: 20px;
		}
		%s
    </style>
</head>
<body>

<div id="app-container">
    %s
</div>

%s
</body>
</html>`, c.Title, iconLink, c.css, c.Page.GetContent(), c.js))
}
