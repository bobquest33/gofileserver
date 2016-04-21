<html>
<head>
    <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" method="post">
	  <label for="file">Filenames:</label>
      <input type="file" name="multiplefiles" id="multiplefiles" multiple webkitdirectory=""/>
      <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload" />
</form>
</body>
</html>
