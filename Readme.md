# GoPlayground

## Mesauring the run time of a go file in Linux/Windows

### Linux

`time go run main.go`

### Windows(Powershell Only)

`Measure-Command {go run main.go}`

### Windows - output

Days : 0
Hours : 0
Minutes : 0
Seconds : 4
Milliseconds : 889
Ticks : 48897706
TotalDays : 5.65945671296296E-05
TotalHours : 0.00135826961111111
TotalMinutes : 0.0814961766666667
TotalSeconds : 4.8897706
TotalMilliseconds : 4889.7706

### Windows(Powershell Only)

`Measure-Command {go run main.go} | select @{n="time";e={$_.Minutes,"Minutes",$_.Seconds,"Seconds",$_.Milliseconds,"Milliseconds" -join " "}}`

### Windows - output

0 Minutes 4 Seconds 908 Milliseconds

## Build go files

`go build .`

this will result in a command file in mac os and exe file in windows

`go build main.go`

you can also do this for building a speific file
