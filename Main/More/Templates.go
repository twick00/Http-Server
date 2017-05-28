package Templates

//Tpl is a thing
const Tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		<h2>{{.Yesman}}</h2>
	</body>
</html>`

//Tyler is cool
func Tyler(athing string) bool {
	if athing == "jump" {
		return true
	}
	return false
}
