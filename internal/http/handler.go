package http

import (
	"html/template"
	"net/http"

	"github.com/oxtoacart/bpool"

	"github.com/krtffl/poop"
	"github.com/krtffl/poop/internal/logger"
)

const K = 42

type Content struct {
	HX bool
}

type Handler struct {
	template *template.Template
	bpool    *bpool.BufferPool
}

func NewHandler(
	bpool *bpool.BufferPool,
) *Handler {
	tmpls, err := template.New("").ParseFS(poop.Public, "public/templates/*.html")
	if err != nil {
		logger.Fatal("[Handler] - Failed to parse templates. %v", err)
	}

	return &Handler{
		template: tmpls,
		bpool:    bpool,
	}
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	logger.Info("[Handler - Login] Incoming request")

	buf := h.bpool.Get()
	defer h.bpool.Put(buf)

	if err := h.template.ExecuteTemplate(buf, "login.html", Content{}); err != nil {
		logger.Error("[Handler - Index] Couldn't execute template. %v", err)
		h.template.ExecuteTemplate(w, "error.html", Content{})
		return
	}

	buf.WriteTo(w)
	return
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	logger.Info("[Handler - Index] Incoming request")

	buf := h.bpool.Get()
	defer h.bpool.Put(buf)

	if err := h.template.ExecuteTemplate(buf, "index.html", Content{}); err != nil {
		logger.Error("[Handler - Index] Couldn't execute template. %v", err)
		h.template.ExecuteTemplate(w, "error.html", Content{})
		return
	}

	buf.WriteTo(w)
	return
}

func isHX(r *http.Request) bool {
	if r.Header.Get("HX-Request") == "true" {
		return true
	}
	return false
}
