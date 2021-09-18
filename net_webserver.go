package main

import (
	"html/template"
	"net/http"
)

const HTML = `
<!DOCTYPE html>
<html lang="en">
	 <head>
	    <meta charset="utf-8">
	    <title>selected attribute</title>
	    
	</head>
	<body>
		<form method="GET">
			<div>
				<label>Places:</label>
				<select id="places" name="places">
					<option value=""></option>
					<option value="test" id="test">test</option>
					<option value="test1" id="test1">test1</option>
					<option value="test2" id="test2">test2</option>
				</select>
			</div>
			<input type="submit" value="submit">
		</form>
	</body>
</html>
`

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("t").Parse(HTML)

	tmpl.Execute(w, selected)
}
