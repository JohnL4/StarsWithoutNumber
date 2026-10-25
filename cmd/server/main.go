package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const distDir = "web/dist"

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", distHandler(distDir)))
}

func distHandler(root string) http.Handler {
	files := http.Dir(root)
	fileServer := http.FileServer(files)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAPI(r.URL.Path) {
			http.NotFound(w, r)
			return
		}
		indexPath := filepath.Join(root, "index.html")
		if _, err := os.Stat(indexPath); err != nil {
			http.Error(w, "web/dist/index.html is missing. Run npm run build in web.", http.StatusServiceUnavailable)
			return
		}
		if isFile(files, r.URL.Path) {
			fileServer.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, indexPath)
	})
}

func isAPI(urlPath string) bool {
	clean := path.Clean(urlPath)
	return clean == "/api" || strings.HasPrefix(clean, "/api/") ||
		clean == "/ws" || strings.HasPrefix(clean, "/ws/")
}

func isFile(dir http.Dir, urlPath string) bool {
	rel := strings.TrimPrefix(path.Clean(urlPath), "/")
	if rel == "" || rel == "." {
		return false
	}
	f, err := dir.Open(rel)
	if err != nil {
		return false
	}
	defer f.Close()
	info, err := f.Stat()
	return err == nil && !info.IsDir()
}
