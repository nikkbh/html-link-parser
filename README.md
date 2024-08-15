### HTML Link Parser Package in Go
A simple package in Go that parses an HTML file and extract all the links inside of <a href="..."> tag's href and the text inside the <a>...</a>

### Options
- A CLI flag `-path` to pass in the HTML file you want to parse. Default stays the `ex1.html` in the `htmls` folder

The output is returned in the terminal of struct type `Link{ Href string, Text string}`
