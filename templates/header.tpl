{{define "header"}}
<!doctype html>

<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<link href="/static/css/hljs/default.css" rel="stylesheet">
	<link href="/static/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/css/main.css" rel="stylesheet">
</head>

<body>
	<div class="container">
		<div class="row col">
			<ol class="breadcrumb">
				{{range $dir := .Dirs }}
				{{if $dir.Active }}
				<li class="active">{{$dir.Name}}</li>
				{{ else }}
				<li><a href="{{ $dir.Path }}">{{$dir.Name}}</a></li>
				{{ end }}
				{{ end }}
				<li class="no-before">{{ if .Revision}}<a href="?revision={{.Revision}}&revisions=1" class="text-muted">{{.Revision}}</a>{{end}}</li>
				{{ if .Edit | or .Revisions }}
				<li class="edit-right">
					<a href="?" class="text-muted">
						<span class="glyphicon glyphicon-remove"></span> Close</a>
				</li>
				{{ else }}
				<li class="edit-right">
					<a href="?edit=1" class="text-muted">
						<span class="glyphicon glyphicon-edit"></span> Edit</a>
				</li>
				{{end}}
			</ol>

		</div>
{{end}}
