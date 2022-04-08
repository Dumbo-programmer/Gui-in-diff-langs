package main

import (
    "github.com/mattn/go-gtk/gdkpixbuf"
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
    "os"
    "os/exec"
    "path"
    "regexp"
    "sort"
    "strings"
)


func main() {
    var menuitem *gtk.MenuItem
    gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetPosition(gtk.WIN_POS_CENTER)
    window.SetTitle("GTK window in GO!")
    window.SetIconName("Icon name here")
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        println("got destroy!", ctx.Data().(string))
        gtk.MainQuit()
    }, "foo")

   
    //--------------------------------------------------------
-----
    // Event
    //--------------------------------------------------------
    window.SetSizeRequest(600, 600)
    window.ShowAll()
    gtk.Main()
}
