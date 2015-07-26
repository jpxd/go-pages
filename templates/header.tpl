{{define "header"}}
<!doctype html>
<head>
 <meta charset="UTF-8">
 <title>{{.Title}}</title>
 <link href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css" rel="stylesheet">
 <meta name="viewport" content="width=device-width, initial-scale=1">
 <style>
   .container{
     margin-top: 15px;
     max-width: 900px;
   }
   .row{
     margin-left: 0;
     margin-right: 0;
   }
 </style>
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
   </ol>
   {{ if .Revision}}<p class="text-muted">Revision: {{.Revision}}</p>{{end}}
 </div>
{{end}}
