{{define "revisions"}}
<div class="row col">
	<div class="list-group">
		{{range $log := .Log}} {{if $log.Link}}
		<a href="?revision={{$log.Hash}}&revisions=1" class="list-group-item">
		{{else}}
		<a href="?revision={{$log.Hash}}&revisions=1" class="list-group-item active">
		{{end}}
		<kbd class="hash">{{$log.Hash}}</kbd> {{$log.Message}} ({{$log.Time}})
		</a>
		{{end}}
	</div>
	<hr />
</div>
{{end}}
