package main

import (
        //"fmt"
        "log"

        "github.com/go-vgo/robotgo"
        "github.com/gotk3/gotk3/gtk"
)

func main() {
        colour := robotgo.GetMouseColor()

        gtk.Init(nil)

        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
               log.Fatal("Unable to create window:", err)
        }
        win.SetTitle("colourshitter")
        win.Connect("destroy", func() {
                gtk.MainQuit()
        })

        l, err := gtk.LabelNew(colour)
        if err != nil {
                log.Fatal("Unable to create label:", err)
        }

        win.Add(l)
        win.SetDefaultSize(200, 100)
        win.ShowAll()

        gtk.Main()
}

