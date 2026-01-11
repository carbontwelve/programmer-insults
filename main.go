package main

import (
	"bufio"
	"embed"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

//go:embed insults.txt
var content embed.FS

var (
	insults []string
	tmpl    *template.Template
)

func init() {
	// Load insults
	data, err := content.ReadFile("insults.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			insults = append(insults, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(insults) == 0 {
		log.Fatal("No insults found in insults.txt")
	}

	// Initialize template
	tmpl = template.Must(template.New("index").Parse(indexHTML))
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	insult := insults[rand.Intn(len(insults))]

	data := struct {
		Insult string
	}{
		Insult: insult,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

const indexHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Programmer Insults!</title>
  <meta name="description" content="Insults for programmers">
  <meta name="author" content="Simon Dann">
  <link href="//fonts.googleapis.com/css?family=Fira+Mono:400,700" rel="stylesheet" type="text/css">
  <style>
	html,body{
		font-family: 'Fira Mono' , monospace;
		padding:0;
		margin:0;
		width: 100vw;
		height: 100vh;
	
		display: flex;
		flex-direction: column; 
		justify-content: center;
		align-items: center;
	}
	
	main {
		max-width: 640px;
		text-align: center;
		flex-grow:1;
	
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	
	h1{
		font-weight:bold;
		font-size: 25px;
	}
	
	a {
		color: #333;
		text-decoration: none; 
	}
	
	main a{
		border: 1px solid #888;
		padding: 8px;
	}
	
	main a:hover{
		background: #e4e4e4;
	}
	
	footer{
		text-align:center;
		padding: 10px 0;
	}
	
	footer a:hover {
		text-decoration: underline;
	}
  </style>
</head>
<body>
    <main>
        <h1>{{.Insult}}</h1>
        <br><br>
        <small><a href="/">More?</a></small>
    </main>
    <footer>Insults sourced from various places. <a href="https://github.com/carbontwelve/programmer-insults">src</a></footer>
</body>
</html>
`
