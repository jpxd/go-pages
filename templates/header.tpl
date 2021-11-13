{{define "header"}}
<!doctype html>

<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<link href="{{ .Basepath }}/static/css/hljs/zenburn.css" rel="stylesheet">
	<link href="{{ .Basepath }}/static/css/bootstrap.min.css" rel="stylesheet">
	<link href="{{ .Basepath }}/static/css/main.css" rel="stylesheet">
</head>

<body>
	<div class="container">
		<div class="row col">
			<ol class="breadcrumb">
				{{range $dir := .Dirs }}
				{{if $dir.Active }}
				<li class="active">{{$dir.Name}}</li>
				{{ else }}
				<li><a href="{{ $.Basepath }}{{ $dir.Path }}">{{$dir.Name}}</a></li>
				{{ end }}
				{{ end }}
				<li class="no-before">{{ if .Revision}}<a href="?revision={{.Revision}}&revisions=1" class="text-muted">{{.Revision}}</a>{{end}}</li>
				<li class="edit-right">
				
				{{ if .Edit | and .AskDelete }}
				<a href="?delete=1" class="text-muted"><span class="glyphicon glyphicon-trash"></span> Are you sure?</a>&nbsp;
				{{ else if .Edit }}
				<a href="?edit=1&askdelete=1" class="text-muted"><span class="glyphicon glyphicon-trash"></span> Delete</a>&nbsp;
				{{ end }}
				{{ if .Edit | or .Revisions }}
					<a href="?" class="text-muted"><span class="glyphicon glyphicon-remove"></span> Close</a>
				{{ else }}
					<a href="?edit=1" class="text-muted"><span class="glyphicon glyphicon-edit"></span> Edit</a>
				{{ end }}
				</li>
			</ol>

		</div>
{{end}}
