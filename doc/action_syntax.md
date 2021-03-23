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

```
count : (any -> integer)
cron [string]
debug [bool]? : (nil -> bool)
dedupe [string]? : (any -> any)
echo [any] : (nil -> bool)
else
get [any] : (nil -> any)
if
json : (string -> any)
message : (any -> bool)
query [string:(css|json|regex)] [string] [string]? : (any -> any)
set [any]... : (nil -> any)
status [string:(http|tcp|dns|tls)] [string] : (nil -> bool)
task [integer] : (nil -> bool)
webhook [string] [string]?
```

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
