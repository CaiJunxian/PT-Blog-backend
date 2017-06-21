package models

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Article struct {
	Id         bson.ObjectId
	Email      string
	CDate      time.Time
	Title      string
	Subject    string
	CommentCnt int
	ReadCnt    int
	Year       int
}

// validate the submit article
func (article *Article) Validate(v *revel.Validation) {
	v.Check(article.Title, revel.Required{}, revel.MinSize{1}, revel.MaxSize{200})
	v.Check(article.Email, revel.Required{}, revel.MaxSize{50})
	v.Email(article.Email)
	v.Check(article.Subject, revel.Required{}, revel.MinSize{1})
}

// find a list of article
func (dao *Dao) FindArticles() []Article {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	articleList := []Article{}
	query := blogCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&articleList)
	return articleList
}

// find article by id
func (dao *Dao) FindArticleById(id string) Article {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	article := Article{}
	fmt.Print(id)
	_id := bson.ObjectIdHex(id)
	query := blogCollection.Find(bson.M{"_id": _id})
	fmt.Print(query)
	query.One(&article)
	return article
}

// get short title of a blog
func (article *Article) GetShortTitle() string {
	if len(article.Title) > 35 {
		return article.Title[:35]
	}
	return article.Title
}

// get short content of a blog
func (article *Article) GetShortContent() string {
	if len(article.Subject) > 200 {
		return article.Subject[:200]
	}
	return article.Subject
}

// create a new article
func (dao *Dao) CreateArticle(article *Article) error {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	article.Id = bson.NewObjectId()
	article.CDate = time.Now()
	article.Year = article.CDate.Year()
	_, err := blogCollection.Upsert(bson.M{"_id": article.Id}, article)
	if err != nil {
		revel.WARN.Printf("Unable to save blog: %v error %v", article, err)
	}
	return err
}
