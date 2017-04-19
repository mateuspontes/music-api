package api

import (
	"net/http"
	"testing"

	"net/http/httptest"

	check "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type ListMusicHandlerSuite struct {
	handler http.Handler
}

type GetMusicHandlerSuite struct {
	handler http.Handler
}

var _ = check.Suite(&ListMusicHandlerSuite{})
var _ = check.Suite(&GetMusicHandlerSuite{})

func (s *ListMusicHandlerSuite) SetUpSuite(c *check.C) {
	s.handler = &ListMusicHandler{}
}

func (s *GetMusicHandlerSuite) SetUpSuite(c *check.C) {
	s.handler = NewRouter()
}

func (s *ListMusicHandlerSuite) TestOk(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/api/musics", nil)
	s.handler.ServeHTTP(w, r)

	c.Assert(w.Code, check.Equals, http.StatusOK)
	c.Assert(w.Body.String(), check.Equals, "Listing musics...")
}

func (s *GetMusicHandlerSuite) TestOk(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/api/musics/1", nil)

	s.handler.ServeHTTP(w, r)

	c.Assert(w.Code, check.Equals, http.StatusOK)
	c.Assert(w.Body.String(), check.Equals, "Show music: 1")
}
