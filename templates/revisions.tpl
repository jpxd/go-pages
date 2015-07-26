{{define "revisions"}}
<div class="row col">
 <hr />
 <div class="list-group">
  {{range $log := .Log}}
   {{if $log.Link}}
    <a href="?revision={{$log.Hash}}&revisions=1" class="list-group-item">
   {{else}}
    <a href="?revision={{$log.Hash}}&revisions=1" class="list-group-item active">
   {{end}}
    <kbd class="hash">{{$log.Hash}}</kbd> {{$log.Message}} ({{$log.Time}})
   </a>
   </li>
  {{end}}
 </div>
</div>
{{end}}
