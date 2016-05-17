package main

import (
    "log"
    "net/http"
    "os"
    "fmt"
    "time"
    "io"
    "crypto/md5"
    "strconv"
    "html/template"
	"path/filepath"
)



func upload(folder string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      fmt.Println("method:", r.Method)
      if r.Method == "GET" {
          crutime := time.Now().Unix()
          h := md5.New()
          io.WriteString(h, strconv.FormatInt(crutime, 10))
          token := fmt.Sprintf("%x", h.Sum(nil))

          t, _ := template.ParseFiles("upload.gtpl")
          t.Execute(w, token)
      } else {
          r.ParseMultipartForm(200000)
    //get the *fileheaders
	formdata := r.MultipartForm // ok, no problem so far, read the Form data
    files := formdata.File["multiplefiles"] // grab the filenames
   	for i, _ := range files { // loop through the files one by one
		thepath, err := filepath.Abs(folder+"/"+filepath.Dir(files[i].Filename))

        if err != nil {
     	   fmt.Fprintf(w, err.Error())
        }else {
			if _, err := os.Stat(thepath); os.IsNotExist(err) {
				// path/to/whatever does not exist
				os.MkdirAll(thepath,0777);

			}
		}
		
        file, err := files[i].Open()
    	defer file.Close()
    	if err != nil {
     	    fmt.Fprintln(w, err)
     	    return
        }

    	out, err := os.Create(folder+"/" + files[i].Filename)
  
    	defer out.Close()
    	if err != nil {
     	    fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
     	    return
       }
  
        _, err = io.Copy(out, file) // file not files[i] !
  
    	if err != nil {
     	fmt.Fprintln(w, err)
     	return
     }
  
    	fmt.Fprintf(w, "Files uploaded successfully : ")
    	fmt.Fprintf(w, files[i].Filename+"\n")
  
    }
      }
    }
}




func main() {
    linkaddress := os.Args[1]
    folder := os.Args[2]
    fpath := folder+"/"
    http.HandleFunc("/upload", upload(folder))
    fmt.Printf("%s\n",fpath)
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(fpath))))
    if err := http.ListenAndServe(linkaddress, nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
