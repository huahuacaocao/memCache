/*
@Time : 26/10/2020
@Author : GC
@Desc : 
*/

package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"memcache/cache"
	"net/http"
	"strings"
)

type Server struct {
	cache.Cache
}

// 处理 cache
func (s *Server) cacheHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 操作的键
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := r.Method
	switch m {
	case http.MethodGet:
		b, err := s.Get(key)
		if err != nil {
			log.Println("err when get", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO 意思是不能存 "" ?
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
	case http.MethodPut:
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("err when get from body", err)
			return
		}
		// TODO 改造支持 存储 ""
		if len(b) != 0 {
			e := s.Set(key, b)
			if e != nil {
				log.Println("err when put", err)
				return
			}
		}
		// TODO 无返回值, 易用性差
		return
	case http.MethodDelete:
		err := s.Del(key)
		if err != nil {
			log.Println("err when delete", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		// TODO 删除成功无返回值, 易用性差
		return
	}
}

// 获取 cache 状态
func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, err := json.Marshal(s.GetStat())
	if err != nil {
		log.Println("err when parse to json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	return
}

func (s *Server) Listen() {
	http.HandleFunc("/cache/", s.cacheHandler)
	http.HandleFunc("/status/", s.statusHandler)

	http.ListenAndServe(":9999", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
