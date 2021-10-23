/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/5 1:42
 * @version     v1.0
 * @filename    s_blog.go
 * @description
 ***************************************************************************/
package schema

import "time"

type Blog struct {
	BlogId         uint      `json:"blogId"`
	BlogTitle      string    `json:"blogTitle"`
	BlogAuthor     string    `json:"blogAuthor"`
	BlogReference  string    `json:"blogReference"`
	BlogLink       string    `json:"blogLink"`
	BlogPreview    string    `json:"blogPreview"`
	BlogPreviewImg string    `json:"blogPreviewImg"`
	BlogView       uint      `json:"blogView"`
	BlogCreateTime time.Time `json:"blogCreateTime"`
	BlogUpdateTime time.Time `json:"blogUpdateTime"`
}

type BlogContent struct {
	BlogId      uint   `json:"blogId"`
	BlogContent string `json:"blogContent"`
}

type BlogWithContent struct {
	BlogId         uint      `json:"blogId"`
	BlogTitle      string    `json:"blogTitle"`
	BlogAuthor     string    `json:"blogAuthor"`
	BlogReference  string    `json:"blogReference"`
	BlogLink       string    `json:"blogLink"`
	BlogPreview    string    `json:"blogPreview"`
	BlogPreviewImg string    `json:"blogPreviewImg"`
	BlogView       uint      `json:"blogView"`
	BlogCreateTime time.Time `json:"blogCreateTime"`
	BlogUpdateTime time.Time `json:"blogUpdateTime"`
	BlogContent    string    `json:"blogContent"`
}
