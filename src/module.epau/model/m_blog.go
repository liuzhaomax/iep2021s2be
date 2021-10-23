/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 6:06
 * @version     v1.0
 * @filename    m_blog.go
 * @description
 ***************************************************************************/
package model

import (
	"cms/src/core"
	"cms/src/module.epau/entity"
	"cms/src/module.epau/schema"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

var BlogSet = wire.NewSet(wire.Struct(new(Blog), "*"))

type Blog struct {
	Tx *gorm.DB
}

func (mBlog *Blog) QueryBlogList() ([]*schema.Blog, error) {
	var eBlogList []entity.Blog
	mBlog.Tx.Order("blog_create_time desc").Find(&eBlogList)
	var sBlogList []*schema.Blog
	for i := 0; i < len(eBlogList); i++ {
		sBlogList = append(sBlogList, eBlogList[i].EntityToSchemaBlog())
	}
	return sBlogList, nil
}

func (mBlog *Blog) QueryBlogListWithLimit(limit uint) ([]*schema.Blog, error) {
	var eBlogList []entity.Blog
	mBlog.Tx.Order("blog_create_time desc").Limit(limit).Find(&eBlogList)
	var sBlogList []*schema.Blog
	for i := 0; i < len(eBlogList); i++ {
		sBlogList = append(sBlogList, eBlogList[i].EntityToSchemaBlog())
	}
	return sBlogList, nil
}

func (mBlog *Blog) QueryOneBlog(blogId uint) (*schema.Blog, error) {
	eBlog := &entity.Blog{}
	mBlog.Tx.Where("blog_id=?", blogId).Find(&eBlog)
	if uint(eBlog.BlogId) == 0 {
		return nil, core.NewError(104, nil)
	}
	sBlog := eBlog.EntityToSchemaBlog()
	return sBlog, nil
}

func (mBlog *Blog) QueryOneBlogContent(blogId uint) (*schema.BlogContent, error) {
	eBlogContent := &entity.BlogContent{}
	mBlog.Tx.Where("blog_id=?", blogId).Find(&eBlogContent)
	if uint(eBlogContent.BlogId) == 0 {
		return nil, core.NewError(104, nil)
	}
	sBlogContent := eBlogContent.EntityToSchemaBlogContent()
	return sBlogContent, nil
}

func (mBlog *Blog) CreateOneBlog(sBlog *schema.Blog, sBlogContent *schema.BlogContent) error {
	eBlog := &entity.Blog{}
	_ = copier.Copy(&eBlog, &sBlog)
	result := addEL(mBlog.Tx).Create(&eBlog)
	if result.Error != nil {
		return result.Error
	}
	eBlogContent := &entity.BlogContent{}
	_ = copier.Copy(&eBlogContent, &sBlogContent)
	result = addEL(mBlog.Tx).Create(&eBlogContent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mBlog *Blog) UpdateOneBlog(sBlog *schema.Blog, sBlogContent *schema.BlogContent) error {
	eBlog := &entity.Blog{}
	_ = copier.Copy(&eBlog, &sBlog)
	result := addEL(mBlog.Tx).Table("blog").Where("blog_id=?", eBlog.BlogId).Update(&eBlog)
	if result.Error != nil {
		return result.Error
	}
	eBlogContent := &entity.BlogContent{}
	_ = copier.Copy(&eBlogContent, &sBlogContent)
	result = addEL(mBlog.Tx).Table("blog_content").Where("blog_id=?", eBlog.BlogId).Update(&eBlogContent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mBlog *Blog) DeleteOneBlog(blogId uint) error {
	result := mBlog.Tx.Table("blog").Where("blog_id=?", blogId).Unscoped().Delete(entity.Blog{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mBlog *Blog) DeleteOneBlogContent(blogId uint) error {
	result := mBlog.Tx.Table("blog_content").Where("blog_id=?", blogId).Unscoped().Delete(entity.BlogContent{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mBlog *Blog) UpdateBlogViews(blogId uint) error {
	eBlog := &entity.Blog{}
	mBlog.Tx.Where("blog_id=?", blogId).Find(&eBlog)
	views := eBlog.BlogView + 1
	result := addEL(mBlog.Tx).Table("blog").Where("blog_id=?", blogId).Update("blog_view", views)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
