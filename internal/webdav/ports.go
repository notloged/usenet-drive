package webdav

import "github.com/chrisfarms/nntp"

type UsenetClient interface {
	GetArticle(id string) (*nntp.Article, error)
}
