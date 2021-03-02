# Workflow Action syntax

> opcode param1 param2 ...

# Grammar

> /internal/app/workflow/action/grammar

# Type

- Integer
- Float
- String
- Boolean
- Message

# Opcode

- get (string|message)
- count
- send (string|message)

## Example

```action
get "https://httpbin.org/get"
json
count
pdf
send "success"
```

```action
get #1
dosomething "param1" "param2" "param3" 
send "success"
```
