# GoPlayground

##Mesauring a go file in Linux/Window

`Linux time go run main.go`

This works only in powershell

`Windows Measure-Command {go run main.go}`

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

`Windows Measure-Command {go run main.go} | select @{n="time";e={$_.Minutes,"Minutes",$_.Seconds,"Seconds",$_.Milliseconds,"Milliseconds" -join " "}}`

0 Minutes 4 Seconds 908 Milliseconds
