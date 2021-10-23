/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 4:05
 * @version     v1.0
 * @filename    e_blog.go
 * @description
 ***************************************************************************/
package entity

import (
	"cms/src/module.epau/schema"
	"time"
)

type Blog struct {
	BlogId         uint      `json:"column:blog_id;primary_key;size:12;"` // yyyymmdd0000
	BlogTitle      string    `json:"column:blog_title;size:1024;not null;"`
	BlogAuthor     string    `json:"column:blog_author;size:1024;not null;"`
	BlogReference  string    `json:"column:blog_reference;size:4096;"`
	BlogLink       string    `json:"column:blog_link;size:1024;"`
	BlogPreview    string    `json:"column:blog_preview;size:1024;not null;"`
	BlogPreviewImg string    `json:"column:blog_preview_img;size:1024;"`
	BlogView       uint      `json:"column:blog_view;size:16;not null;default:0;"`
	BlogCreateTime time.Time `json:"column:blog_create_time;not null;"`
	BlogUpdateTime time.Time `json:"column:blog_update_time;not null;"`
}

func (eBlog Blog) EntityToSchemaBlog() *schema.Blog {
	sBlog := new(schema.Blog)
	_ = Transform(eBlog, sBlog)
	return sBlog
}

type BlogContent struct {
	BlogId      uint   `json:"column:blog_id;primary_key;size:12;"` // yyyymmdd0000
	BlogContent string `json:"column:blog_content;size:16000;not null;"`
}

func (eBlogContent BlogContent) EntityToSchemaBlogContent() *schema.BlogContent {
	sBlogContent := new(schema.BlogContent)
	_ = Transform(eBlogContent, sBlogContent)
	return sBlogContent
}

type BlogWithContent struct {
	BlogId         uint      `json:"column:blog_id;primary_key;size:12;"` // yyyymmdd0000
	BlogTitle      string    `json:"column:blog_title;size:1024;not null;"`
	BlogAuthor     string    `json:"column:blog_author;size:1024;not null;"`
	BlogReference  string    `json:"column:blog_reference;size:4096;"`
	BlogLink       string    `json:"column:blog_link;size:1024;"`
	BlogPreview    string    `json:"column:blog_preview;size:1024;not null;"`
	BlogPreviewImg string    `json:"column:blog_preview_img;size:1024;"`
	BlogView       uint      `json:"column:blog_view;size:16;not null;default:0;"`
	BlogCreateTime time.Time `json:"column:blog_create_time;not null;"`
	BlogUpdateTime time.Time `json:"column:blog_update_time;not null;"`
	BlogContent    string    `json:"column:blog_content;size:16000;not null;"`
}
