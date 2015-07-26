{{define "header"}}
<!doctype html>

<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<link href="/static/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/css/wiki.css" rel="stylesheet">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
	<div class="container">
		<div class="row col">
			<ol class="breadcrumb">
				{{range $dir := .Dirs }} {{if $dir.Active }}
				<li class="active">{{$dir.Name}}</li>
				{{ else }}
				<li><a href="{{ $dir.Path }}">{{$dir.Name}}</a></li>
				{{ end }} {{ end }}
				<li class="no-before">{{ if .Revision}}<a href="?revision={{.Revision}}&revisions=1" class="text-muted">{{.Revision}}</a>{{end}}</li>
				{{ if not .Edit}}
				<li class="edit-right">
					<a href="?edit=1" class="text-muted">
						<span class="glyphicon glyphicon-edit"></span> Edit</a>
				</li>{{end}}
			</ol>

		</div>
{{end}}
