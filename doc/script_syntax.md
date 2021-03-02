# Workflow Flowscript (syntax)[./internal/app/workflow/script/grammar]

# Type

- Integer
- Float
- String
- Boolean
- List
- Dict
- Message
- Node
- Workflow

# Statements

- if
- while
- print
- assignment
- flow

# Operators

- and
- or
- +, -, *, /
- >, >=, <, <=
- !=, ==
- #
- @

## Example
```Flowscript
#!/usr/bin/env flowscript

node abc (cron):
	with: {
			"mode": "custom",
			"cron_expression": "* * * * *"
		}
	secret: github_key
end

node xkcd (http):
	with: { 
			"method": "GET",
	 		"url": "https://xkcd.com",
	 		"response_format": "html",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar"},
			"extract": {
				"url": {
				  "css": "#comic img",
				  "value": "@src"
				},
				"title": {
				  "css": "#comic img",
				  "value": "@alt"
				},
				"hovertext": {
				  "css": "#comic img",
				  "value": "@title"
				}
			}
	 	}
end

node httpbin (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/json",
	 		"response_format": "json",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar" },
			"extract": {
			  "title": { "path": "slideshow.slides.#.title", "value": "string(.)" },
			  "type": { "path": "slideshow.slides.#.type", "value": "string(.)" }
			}
	 	}
end

node hi (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/uuid",
	 		"response_format": "text",
			"extract": {
			  "uuid": { "regexp": "(\w{8}\-\w{4}\-\w{4}\-\w{4}\-\w{12})", "index": 1, "value": "string(.)" }
			}
	 	}
end

node notice (pushover):
	with: {
		"title": "title - {{0.title}}",
		"message": "message - {{0.title}}",
		"url": "{{0.url}}"
	}
	secret: pushover
end

workflow demo:
    @xkcd -> @httpbin
end

workflow main:
    @abc -> @xkcd -> @httpbin -> @hi -> @notice;
	@xkcd -> @notice
end
```