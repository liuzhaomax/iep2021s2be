/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/16 0:26
 * @version     v1.0
 * @filename    b_blog.go
 * @description
 ***************************************************************************/
package service

import (
	"cms/src/core"
	"cms/src/module.epau/model"
	"cms/src/module.epau/schema"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

var BlogSet = wire.NewSet(wire.Struct(new(BBlog), "*"))

type BBlog struct {
	MBlog *model.Blog
}

func (bBlog *BBlog) ReadBlogList(c *gin.Context) ([]*schema.Blog, *core.Error) {
	blogList, _ := bBlog.MBlog.QueryBlogList()
	return blogList, nil
}

func (bBlog *BBlog) ReadFEBlogList(c *gin.Context) ([]*schema.Blog, *core.Error) {
	limitStr := c.Param("limit")
	limit := StrToUInt(limitStr)
	blogList, _ := bBlog.MBlog.QueryBlogListWithLimit(limit)
	return blogList, nil
}

func (bBlog *BBlog) ReadBlog(c *gin.Context) (*schema.BlogWithContent, *core.Error) {
	blogIdStr := c.Param("blogid")
	blogId := StrToUInt(blogIdStr)
	blog, _ := bBlog.MBlog.QueryOneBlog(blogId)
	blogContent, _ := bBlog.MBlog.QueryOneBlogContent(blogId)
	blogWithContent := new(schema.BlogWithContent)
	_ = copier.Copy(&blogWithContent, &blog)
	_ = copier.Copy(&blogWithContent, &blogContent)
	return blogWithContent, nil
}

func (bBlog *BBlog) CreateBlog(c *gin.Context, sBlogWithContent *schema.BlogWithContent) error {
	var sBlog schema.Blog
	var sBlogContent schema.BlogContent
	_ = copier.Copy(&sBlog, &sBlogWithContent)
	_ = copier.Copy(&sBlogContent, &sBlogWithContent)
	now := time.Now()
	sBlog.BlogCreateTime = now
	sBlog.BlogUpdateTime = now
	err := core.ExecTrans(c, bBlog.MBlog.Tx, func(ctx context.Context) error {
		err := bBlog.MBlog.CreateOneBlog(&sBlog, &sBlogContent)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (bBlog *BBlog) UpdateBlog(c *gin.Context, sBlogWithContent *schema.BlogWithContent) error {
	var sBlog schema.Blog
	var sBlogContent schema.BlogContent
	_ = copier.Copy(&sBlog, &sBlogWithContent)
	_ = copier.Copy(&sBlogContent, &sBlogWithContent)
	now := time.Now()
	sBlog.BlogUpdateTime = now
	err := core.ExecTrans(c, bBlog.MBlog.Tx, func(ctx context.Context) error {
		blog, _ := bBlog.MBlog.QueryOneBlog(sBlog.BlogId)
		sBlog.BlogView = blog.BlogView
		sBlog.BlogCreateTime = blog.BlogCreateTime
		err := bBlog.MBlog.UpdateOneBlog(&sBlog, &sBlogContent)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (bBlog *BBlog) DeleteBlog(c *gin.Context) error {
	blogIdStr := c.Param("blogid")
	blogId := StrToUInt(blogIdStr)
	err := core.ExecTrans(c, bBlog.MBlog.Tx, func(ctx context.Context) error {
		err := bBlog.MBlog.DeleteOneBlog(blogId)
		if err != nil {
			return err
		}
		err = bBlog.MBlog.DeleteOneBlogContent(blogId)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (bBlog *BBlog) IncreaseViews(c *gin.Context) error {
	blogIdStr := c.Param("blogid")
	blogId := StrToUInt(blogIdStr)
	err := bBlog.MBlog.UpdateBlogViews(blogId)
	if err != nil {
		return err
	}
	return err
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
