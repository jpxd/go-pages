{{define "header"}}
<!doctype html>

<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<link href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css" rel="stylesheet">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		.container {
			padding-top: 20px;
			max-width: 900px;
		}

		.row {
			margin-left: 0;
			margin-right: 0;
		}

		h1:first-child {
			margin-top: 0;
		}

		.edit-right {
			float: right;
		}

		.edit-right:before {
			content: none!important;
			padding: 0!important;
		}

		.hash {
			background-color: rgba(1, 1, 1, 0.14);
			padding: 4px;
			border-radius: 4px;
			margin-right: 8px;
		}
	</style>
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
