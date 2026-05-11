package main
import ("fmt";"mime";"os";"path";"strings")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: mime-type <filename>"); os.Exit(1) }
	filename := os.Args[1]
	ext := path.Ext(filename)
	mimeType := mime.TypeByExtension(strings.ToLower(ext))
	if mimeType == "" { mimeType = "application/octet-stream" }
	fmt.Println(mimeType)
}
