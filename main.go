package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/zserge/lorca"
)

type Config struct {
	Docassets bool `json:Docassets`
}

func main() {
	/*
		if _, err := os.Stat("/plugins"); os.IsNotExist(err) {
			var err = os.Mkdir("plugins", 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	*/
	cwd, _ := os.Getwd()
	docpath := path.Join(cwd, "plugins/docassets")
	// swappath := path.Join(cwd, "swapper")
	flags := fmt.Sprintf("--load-extension=%v", docpath)
	var config Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		config = Config{false}
		data, _ := json.Marshal(config)
		os.WriteFile("config.json", data, 0644)
	} else {
		json.Unmarshal(data, &config)
	}
	cfcpy := config

	ui, _ := lorca.New("data:text/html,"+url.PathEscape(`
		<title>Deeeep.io Desktop App</title>
		<style>*{padding: 0; margin: 0; overflow: hidden}
		#loading{text-align: center; font-size: 40px; font-weight: bold; position: absolute; z-index: 2; color: white;}</style>
		<!-- <img id="img" src="https://tinyurl.com/m5a4d3vw"></img> -->
		<img id="img" src="https://sralcodeproj.netlify.app/assets/myhailot.png"></img>
		<h3 id="loading">Loading...</h3>
	`)+query(config), "", 1120, 740, flags, "--disable-dinosaur-easter-egg") // previous: 887x586
	time.Sleep(5 * time.Second)

	err = client.Login("817817065862725682")
	if err != nil {
		fmt.Println(err)
		// panic(err) //this apparently crashes the entire client if it doesn't find disc
	}
	now := time.Now()
	err = client.SetActivity(client.Activity{
		Details:    "Playing Unknown gamemode",
		LargeImage: "deeplarge_2",
		LargeText:  "Playing Deeeep.io Desktop Client",
		SmallImage: "ffa",
		SmallText:  "Playing Unknown Gamemode",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	ui.Bind("setdocassets", func() {
		config.Docassets = !config.Docassets
		ui.Load(`https://beta.deeeep.io` + query(config))
		menu(ui)
	})

	ui.Bind("reload", func() {
		ui.Load(`https://beta.deeeep.io` + query(config))
		menu(ui)
	})

	ui.Load(`https://beta.deeeep.io` + query(config))
	ui.SetBounds(lorca.Bounds{0, 0, 1200, 1000, "maximized"})
	menu(ui)

	defer func() {
		ui.Close()
		if cfcpy != config {
			data, _ := json.Marshal(config)
			os.WriteFile("config.json", data, 0644)
		}
	}()
	<-ui.Done()
}

func query(config Config) string {
	return fmt.Sprintf("?docassets=%v", config.Docassets)
}

/* default
sirreads' pc:
GOARCH=amd64
GOOS=windows
*/
func menu(ui lorca.UI) {
	ui.Eval(`
	window.onload = () => {
		const app = document.getElementById("app")
		//const div = document.createElement('div')
		//div.style.position = "absolute"
		//div.style.zIndex = "100"
		//div.style.width = "100%"
		// div.innerHTML = ''
		document.getElementByClassName("el-row justify-center").innerHTML += "<div class="el-col el-col-12 el-col-sm-8 is-guttered text-center"><button class="el-button el-button--default btn nice-button green has-icon square"></button></div>"
		// document.body.insertBefore(div, app)

		const url = new URL(window.location.href);
		const enabled = url.searchParams.get("docassets")
		const btn = document.getElementById("doc-btn")
		if (enabled == "true") {
			btn.className = "position-absolute bg-success text-white px-1 rounded"
			btn.innerText = "ON"
		} else {
			btn.className = "position-absolute bg-danger text-white px-1 rounded"
			btn.innerText = "OFF"
		}

		function docassets() {
			const btn = document.getElementById("doc-btn")
			if (enabled == "true") {
				localStorage.setItem("docassets", "false")
			} else {
				localStorage.setItem("docassets", "true")
			}
			setdocassets()
		}

		const doc = document.getElementById("docassets")
		doc.onclick = docassets
	}
	
	function dark() {
		const menu = document.getElementById("menu")
		if (menu.classList.contains("navbar-light")) {
			menu.className = "navbar navbar-expand navbar-dark bg-dark"
		} else {
			menu.className = "navbar navbar-expand navbar-light bg-light"
	}
	
	function open_navbar() {
		console.log("DDDDDDDDDDDDDDDD")
	}
	}`)
}
