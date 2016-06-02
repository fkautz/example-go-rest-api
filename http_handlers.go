package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m := make(map[string]string)
	m["hello"] = vars["person"]
	js, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(js)
}

func structuredHandler(w http.ResponseWriter, r *http.Request) {
	sampleResponse := SampleResponse{
		Hello: "world",
		Count: 10,
		Address: FullAddress{
			StreetNumber: 1234,
			Street:       "Easy St.",
			City:         "Sun City",
			State:        "CA",
			Zip:          54321,
			Country:      "United States",
		},
	}
	js, err := json.Marshal(sampleResponse)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(js)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connection upgraded")
	for {
		p := []byte(time.Now().UTC().String())
		err = conn.WriteMessage(websocket.TextMessage, p)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// Modified from github.com/websocket/examples/echo
func echoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, "ws://"+r.Host+"/ws/time")
}

var tmpl = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {

var output = document.getElementById("output");
var input = document.getElementById("input");
var ws;

var print = function(message) {
var d = document.createElement("div");
d.innerHTML = message;
output.appendChild(d);
};

document.getElementById("open").onclick = function(evt) {
if (ws) {
return false;
}
ws = new WebSocket("{{.}}");
ws.onopen = function(evt) {
print("OPEN");
}
ws.onclose = function(evt) {
print("CLOSE");
ws = null;
}
ws.onmessage = function(evt) {
print("RESPONSE: " + evt.data);
}
ws.onerror = function(evt) {
print("ERROR: " + evt.data);
}
return false;
};

document.getElementById("send").onclick = function(evt) {
if (!ws) {
return false;
}
print("SEND: " + input.value);
ws.send(input.value);
return false;
};

document.getElementById("close").onclick = function(evt) {
if (!ws) {
return false;
}
ws.close();
return false;
};

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<form>
<button id="open">Open</button>
<button id="close">Close</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))

