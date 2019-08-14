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
  manifestPath := apppath + "/static/manifest.json"
  manifestBytes, err := ioutil.ReadFile(manifestPath)
  if err != nil {
    logs.Error(err)
  }

  var manifest struct {
    App           map[string]interface{} `json:"app"`
    BabelPolyfill map[string]interface{} `json:"babel-polyfill"`
  }

  if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
    logs.Error(err)
  }

  conn.Data["AppCSS"] = manifest.App["css"]
  conn.Data["AppJS"] = manifest.App["js"]
  conn.Data["BabelPolyfill"] = manifest.BabelPolyfill["js"]

  conn.TplName = "index.html"
}

func init() {
  _, file, _, _ := runtime.Caller(0)
  apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
}
