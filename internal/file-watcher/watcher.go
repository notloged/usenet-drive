package filewatcher

import (
	"context"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	uploadqueue "github.com/javi11/usenet-drive/internal/upload-queue"
)

type Watcher struct {
	watcher       *fsnotify.Watcher
	queue         uploadqueue.UploadQueue
	log           *log.Logger
	fileWhitelist []string
}

func NewWatcher(queue uploadqueue.UploadQueue, log *log.Logger, fileWhitelist []string) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		watcher:       watcher,
		queue:         queue,
		log:           log,
		fileWhitelist: fileWhitelist,
	}, nil
}

func (w *Watcher) Start(ctx context.Context) {
	w.log.Printf("Starting file watcher...")

	go func() {
		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					// Check if file extension is on fileWhitelist
					for _, ext := range w.fileWhitelist {
						if strings.HasSuffix(event.Name, ext) {
							w.log.Printf("File %s created, adding to upload queue", event.Name)
							w.queue.AddJob(ctx, event.Name)
							break
						}
					}
				}
			case err, ok := <-w.watcher.Errors:
				if err != nil {
					w.log.Printf("file watcher error: %v", err)
					return
				}
				if !ok {
					return
				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *Watcher) Add(path string) error {
	w.log.Printf("Adding %s to file watcher", path)
	return w.watcher.Add(path)
}

func (w *Watcher) Close() error {
	return w.watcher.Close()
}
