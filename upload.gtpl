<!doctype html>
<html>
<head>
<title>Proper Title</title>
</head>
    
<body>
    
    <form id="myForm" method="post" enctype="multipart/form-data">

        Files: <input type="file" id="multiplefiles" name="files" multiple webkitdirectory=""><br/>
		<input type="hidden" name="token" value="{{.}}"/>
		<input type="submit" value="upload">
		
        <div id="selectedFiles"></div>

    </form>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-beta1/jquery.js"></script>
    <script>
    var selDiv = "";
	var files = null;
	var filenames = [];
	var fid = null;
        
    document.addEventListener("DOMContentLoaded", init, false);
    
    function init() {
        document.querySelector('#multiplefiles').addEventListener('change', handleFileSelect, false);
        selDiv = document.querySelector("#selectedFiles");
    }
        
    function handleFileSelect(e) {
        
        if(!e.target.files) return;
        
        selDiv.innerHTML = "";
        
        files = e.target.files;
        for(var i=0; i<files.length; i++) {
            var f = files[i];
            filenames.push(files[i].name)
            selDiv.innerHTML += "<span id="+i+">"+f.name + "<br/>";

        }
        
    }
	
	$('form').submit(function (e) {
		e.preventDefault();
	
		for(var i=0; i<files.length; i++) {
		    var data;
			
		    data = new FormData();
			 
		    data.append('multiplefiles', $('#multiplefiles')[0].files[i]);
		
		    $.ajax({
		        url: '/upload',
		        data: data,
		        processData: false,
				contentType: false,
		        type: 'POST',
		
		
		        success: function (data) {
		            
					var str = data.split(":").pop().trim()
					if(str.indexOf("/") >= 0){
						str = data.split("/").pop()
					}
					
					console.log(data)
					id = filenames.indexOf(str.trim())
					$("#"+id).html(data+"<br>");


		        }
				

				
		    });
		}
	});
	
    </script>

</body>
</html>