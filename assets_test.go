package main

import (
	"github.com/go-chi/chi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("assets.go", func() {

	Describe("#Dir", func() {

		It("serves static assets", func() {
			r := chi.NewMux()
			r.Handle("/*", http.FileServer(Dir(false, "/assets")))

			ts := httptest.NewServer(r)
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				Fail(err.Error())
			}
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			res, err = http.Get(ts.URL + "/js/index.js")
			if err != nil {
				Fail(err.Error())
			}
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			res, err = http.Get(ts.URL + "/index.html")
			if err != nil {
				Fail(err.Error())
			}
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

	})
})
