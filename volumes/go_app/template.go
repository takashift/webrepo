package main

import (
    "io"
    "github.com/labstack/echo"
    "github.com/yosssi/ace"
)

// e.Renderer に代入するために必須っぽい
type AceTemplate struct {
}

// ここでレシーバ変数を定義したことでAceTemplateに以下の関数がメソッドとして関連付けられる
func (at *AceTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    // オプションを渡してディレクトリを指定
    tpl, err := ace.Load(name, "", &ace.Options{
      BaseDir:       "views",
    })
    
    if err != nil {
      return err 
    }
   
    return tpl.Execute(w, data)
}

// サイトで共通情報
type ServiceInfo struct {
  Title string
}

var serviceInfo = ServiceInfo {
  "タイトルrrrrrrrrrrr",
}

func init(){

    // ここで入れるべき構造体はinterfaceによって必須のメソッドが定義され、持つべき引数が決まっている。GoDoc参照。
    e.Renderer = &AceTemplate{}
}

