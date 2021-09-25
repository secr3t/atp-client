package client

import "testing"

func TestSearchClient_SearchTilLimit(t *testing.T) {
	c := NewSearchClient("")

	items := c.SearchTilLimit("https://s.taobao.com/search?spm=a230r.1.0.0.e77c9bfciiL3RY&q=%E6%A1%8C%E9%9D%A2%E6%91%86%E4%BB%B6&rs=up&rsclick=7&preq=%E6%91%86%E4%BB%B6&cps=yes&cat=50035867", 200)

	t.Log(items)
}
