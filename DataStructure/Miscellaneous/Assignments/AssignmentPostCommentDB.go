package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=192.168.56.115 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id =  $1", id).Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id =$1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func (comment *Comment) Create() (err error) {

	if comment.Post == nil {
		err = errors.New("Post is not found")
		return
	}

	err = Db.QueryRow("insert into comments (content,author, post_id ) values  ($1, $2, $3) returning id ",
		comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)

	return

}

func main() {
	post := Post{Author: "Purushottam Jetha", Content: "Programming in GO"}

	fmt.Println(post)
	post.Create()

	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPosts, _ := Posts(100)

	for i, row := range readPosts {
		fmt.Println("Print Line %d", i)
		fmt.Println(row)
	}

	readPost.Author = "Anirudh Jetha"
	readPost.Content = "Abacus For Genuise"
	readPost.Update()

	fmt.Println("After update")
	fmt.Println(readPost)

	readPost, _ = GetPost(readPost.Id)
	fmt.Println("Get updated Row")
	fmt.Println(readPost)

}
