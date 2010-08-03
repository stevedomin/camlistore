package main

import (
	"crypto/sha1"
	"fmt"
	"http"
	"io"
)

func handleCamliForm(conn *http.Conn, req *http.Request) {
	fmt.Fprintf(conn, `
<html>
<body>
<form method='POST' enctype="multipart/form-data" action="/camli/testform">
<input type="hidden" name="имя" value="брэд" />
Text unix: <input type="file" name="file-unix"><br>
Text win: <input type="file" name="file-win"><br>
Text mac: <input type="file" name="file-mac"><br>
Image png: <input type="file" name="image-png"><br>
<input type=submit>
</form>
</body>
</html>
`)
}


func handleTestForm(conn *http.Conn, req *http.Request) {
	if !(req.Method == "POST" && req.URL.Path == "/camli/testform") {
		badRequestError(conn, "Inconfigured handler.")
		return
	}

	multipart, err := req.MultipartReader()
	if multipart == nil {
		badRequestError(conn, fmt.Sprintf("Expected multipart/form-data POST request; %v", err))
		return
	}

	for {
		part, err := multipart.NextPart()
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}
		if part == nil {
			break
		}
		formName := part.FormName()
		fmt.Printf("New value [%s], part=%v\n", formName, part)

		sha1 := sha1.New()
		io.Copy(sha1, part)
		fmt.Printf("Got part digest: %x\n", sha1.Sum())

	}
	fmt.Println("Done reading multipart body.")

}