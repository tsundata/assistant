package components

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/version"
	"html/template"
)

type Html struct {
	Name          string
	Title         string
	UseIcon       bool
	UseCodeEditor bool
	Page          Component
	css           template.CSS
	js            template.JS
}

func (c *Html) SetCss(css template.CSS) {
	c.css = css
}

func (c *Html) SetJs(js template.JS) {
	if js != "" {
		c.js = `<script>` + js + `</script>`
	}
}

func (c *Html) GetContent() template.HTML {
	title := fmt.Sprintf("%s (v%s)", c.Title, version.Version)
	iconLink := ""
	if c.UseIcon {
		iconLink = `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css">`
	}
	codeEditorLink := ""
	codeEditorScript := ""
	if c.UseCodeEditor {
		codeEditorLink = `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/codemirror@5.59.4/lib/codemirror.css">`
		codeEditorScript = `<script src="https://cdn.jsdelivr.net/npm/codemirror@5.59.4/lib/codemirror.js"></script>
<script>
    var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
        lineNumbers: true,
        tabSize: 4,
        indentUnit: 4,
        indentWithTabs: true,
        mode: "text/x-go"
    });
	editor.setSize(null, 800);
</script>
`
	}
	return template.HTML(fmt.Sprintf(`
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge">
    <meta name="viewport" content="width=device-width,initial-scale=1,shrink-to-fit=no,user-scalable=no">
	<title>%s</title>
	%s
	%s
    <style>
		a {
			text-decoration: none;
			color: #003C97;
			word-break: break-all;
		}
		.form {
			width: 100%%;
			margin-top: 50px;
		}
		.form button {
			width: 100%%;
			height: 50px;
			border: none;
			margin: 10px;
			border-radius: 5px;
		}
		.form .button {
			width: 100%%;
			display: flex;
			flex-direction: row;
		}
		.form .input {
			display: flex;
			flex-direction: column;
			margin-bottom: 20px;
		}
		.form .input label {
			margin-bottom: 10px;
		}
		.form .input input {
			height: 50px;
			padding-left: 20px;
		}
		.form .select {
			display: flex;
			flex-direction: column;
			margin-bottom: 20px;
		}
		.form .select label {
			margin-bottom: 10px;
		}
		.form .select select {
			height: 50px;
			padding-left: 20px;
		}
        #app-container {
            color: #333;
            padding: 15px 25px;
        }
        .title {
            font-size: 21px;
            border-bottom: 1px solid #e8ecf1;
            padding-bottom: 15px;
			display: flex;
			flex-direction: row;
			justify-content: space-between;
        }
		.title span {
			font-size: 13px;
		}
		.content {
	    	display: flex;
			flex-direction: row;
			flex-wrap: wrap;
		}
        .content .text {
            font-size: 18px;
            margin-top: 25px;
            margin-bottom: 15px;
			width: 100%%;
        }
		.link-button {
			width: 100%%;
		}
        .link-button .link-content {
            background-color: #f7f7f7;
			padding: 10px;
        }
        .link-button a {
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
		.memo {
			width: 100%%;
			padding: 20px;
			border: 1px solid #f9f9f9;
			border-radius: 15px;
			margin-bottom: 15px;
			margin-top: 10px;
			background-color: #fafafa;
		}
		.memo .time {
			font-size: 13px
		}
		.memo .tags {
			display: flex;
			margin-top: 10px;
			font-size: 13px;
		}
		.memo .tags span {
			display: flex;
			background-color: #A9EA79;
			border-radius: 10px;
			padding: 2px 7px;
			margin-right: 5px;
			align-items: center;
			justify-content: center;
		}
		.memo .content {
			font-size: 15px;
		}
		.memo .content .text {
			font-weight: normal;
		}
		.script {
            width: 100%%;
        }
        .script .title {
            font-size: 20px;
            display: flex;
            flex-direction: row;
            align-items: center;
            justify-content: space-between;
        }
        .script .id {
            width: 20%%;
        }
        .script .run-btn {
            width: 20%%;
            height: 40px;
			text-align: center;
			line-height: 40px;
			background-color: #efefef;
			border-radius: 5px;
        }
        .script div {
            font-size: 20px;
            padding: 10px;
        }
        .script pre {
            background-color: #FAFAFAFF;
            padding: 25px;
            overflow: scroll;
        }
		%s
    </style>
</head>
<body>

<div id="app-container">
    %s
</div>

%s
%s
</body>
</html>`, title, iconLink, codeEditorLink, c.css, c.Page.GetContent(), codeEditorScript, c.js))
}
