package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

//func main33() {
//	w := watcher.New()
//	w.IgnoreHiddenFiles(true)
//	w.FilterOps(watcher.Create,watcher.Remove)
//	pathDir := "/users/yuhang/data/google"
//
//	// Uncomment to use SetMaxEvents set to 1 to allow at most 1 event to be received
//	// on the Event channel per watching cycle.
//	//
//	// If SetMaxEvents is not set, the default is to send all events.
//	// w.SetMaxEvents(1)
//
//	// Uncomment to only notify rename and move events.
//	// w.FilterOps(watcher.Rename, watcher.Move)
//
//	// Uncomment to filter files based on a regular expression.
//	//
//	// Only files that match the regular expression during file listing
//	// will be watched.
//	r := regexp.MustCompile(`^*\.mbtiles$`)
//	//fmt.Println(r.MatchString(".mbtileserver"))
//	w.AddFilterHook(watcher.RegexFilterHook(r, false))
//	//wchan := make(chan watcher.Event,1)
//
//
//
//	go func() {
//		for {
//			select {
//			case event := <-w.Event:
//				fmt.Println(event) // Print the event's info.
//			case err := <-w.Error:
//				log.Fatalln(err)
//			case <-w.Closed:
//				return
//			}
//		}
//	}()
//
//	// Watch test_folder recursively for changes.
//	if err := w.AddRecursive(pathDir); err != nil {
//		log.Fatalln(err)
//	}
//
//	// Print a list of all of the files and folders currently
//	// being watched and their paths.
//	//for path, f := range w.WatchedFiles() {
//	//	fmt.Printf("%s: %s\n", path, f.Name())
//	//}
//	//
//	//fmt.Println()
//
//	// Start the watching process - it'll check for changes every 100ms.
//	if err := w.Start(time.Millisecond * 100); err != nil {
//		log.Fatalln(err)
//	}
//
//}

var fs *fsnotify.Watcher

func main3() {
	fs, _ = fsnotify.NewWatcher()
	defer fs.Close()

	if err := filepath.Walk("/users/yuhang/data/google", watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-fs.Events:
				fmt.Printf("EVENT! %#v\n", event)

			case err := <-fs.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return fs.Add(path)
	}

	return nil
}
