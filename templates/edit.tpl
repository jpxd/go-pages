{{ template "header" . }}
<div class="row col">
	<form method="POST" action="?">
		<div class="form-group col">
			<textarea type="text" class="form-control editbox" rows="15" placeholder="Insert markdown here" name="content">{{ .Content }}</textarea>
		</div>
		<div class="form-inline col">
			<div class="form-group col-md-8">
				<input type="text" class="form-control" name="msg" placeholder="Changelog" />
			</div>
			<div class="form-group col-md-2">
				<input type="text" class="form-control" name="author" placeholder="Author" />
			</div>
			<div class="form-group col-md-2">
				<button type="submit" class="btn btn-default">
					<span class="glyphicon glyphicon-floppy-disk"></span> Save
				</button>
			</div>
		</div>
	</form>
</div>
{{ template "footer" . }}
