### Rule

```
	title: $(".title a.storylink").text
	site: $(".title span.sitestr").text
	link: $(".title a.storylink").href
```

#### Fun

Fun represents the data processing process.

There are all supported funs:

| Name      | Parameters                       | Description                                    |
| --------- | -------------------------------- | ---------------------------------------------- |
| $         | (selector: string)               | Relative CSS selector (select from parent node)|
| html      |                                  | inner HTML                                     |
| text      |                                  | inner text                                     |
| outerHTML |                                  | outer HTML                                     |
| attr      | (attr: string)                   | attribute value                                |
| style     |                                  | style attribute value                          |
| href      |                                  | href attribute value                           |
| src       |                                  | src attribute value                            |
| class     |                                  | class attribute value                          |
| id        |                                  | id attribute value                             |
| calc      | (prec: int)                      | calculate arithmetic expression                |
| match     | (regexp: string)                 | match first sub-string via regular expression  |
| expand    | (regexp: string, target: string) | expand matched strings to target string        |
