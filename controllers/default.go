package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// MainController doc false
type MainController struct {
	beego.Controller
}

var apppath string

// Get http method
func (conn *MainController) Get() {
	env := beego.AppConfig.String
	if env("runmode") == "prod" {
		manifestPath := apppath + "/static/manifest.json"
		manifestBytes, err := ioutil.ReadFile(manifestPath)
		if err != nil {
			logs.Error(err)
		}

		var manifest struct {
			App    map[string]interface{} `json:"app"`
			El     map[string]interface{} `json:"element"`
			Vue    map[string]interface{} `json:"vue"`
			Babel  map[string]interface{} `json:"babel-polyfill"`
			BabelV map[string]interface{} `json:"vendors~babel-polyfill"`
		}

		if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
			logs.Error(err)
		}

		conn.Data["AppCSS"] = manifest.App["css"]
		conn.Data["ElCSS"] = manifest.El["css"]
		conn.Data["AppJS"] = manifest.App["js"]
		conn.Data["ElJS"] = manifest.El["js"]
		conn.Data["VueJS"] = manifest.Vue["js"]
		conn.Data["BabelPolyfillJS"] = manifest.Babel["js"]
		conn.Data["VBabelPolyfillJS"] = manifest.BabelV["js"]
	} else {
		conn.Data["AppCSS"] = "/css/app.css"
		conn.Data["ElCSS"] = "/css/element.css"
		conn.Data["AppJS"] = "/js/app.js"
		conn.Data["ElJS"] = "/js/element.js"
		conn.Data["VueJS"] = "/js/vue.js"
		conn.Data["BabelPolyfillJS"] = "/js/babel-polyfill.js"
		conn.Data["VBabelPolyfillJS"] = "/js/vendors~babel-polyfill.js"
	}

	conn.TplName = "index.html"
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
}
