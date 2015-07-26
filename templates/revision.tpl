{{ define "revision" }}
 <!-- Actions for a specific revision (revert, diff etc) -->
<div class="row col">
 <form method="POST">
  <div class="form-group">
   <button type="submit" class="btn btn-danger btn-xs">
    <span class="glyphicon glyphicon-step-backward"></span> Revert to this version
   </button>
   <input type="hidden" name="revert" value="{{ .Revision }}" />
  </div>
 </form>
</div>
{{end}}
