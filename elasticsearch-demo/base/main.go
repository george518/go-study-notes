/************************************************************
** @Description: base
** @Author: george hao
** @Date:   2019-01-22 10:05
** @Last Modified by:  george hao
** @Last Modified time: 2019-01-22 10:05
*************************************************************/
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"fmt"

	"reflect"

	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

type Student struct {
	Name    string `json:"name"`
	Class   string `json:"class"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func main() {
	var (
		client *elastic.Client
		err    error
		std1   Student
		ctx    context.Context
	)
	//创建连接
	client, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200/"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	ctx = context.Background()

	//#################存储

	//第一种方式存储 bodyjson 对象
	std1 = Student{
		Name:    "george",
		Message: "good boy2",
		Age:     21,
		Class:   "三年一班1",
	}

	put, err := client.Index().
		Index("school").
		Type("class").
		Id("1").
		BodyJson(std1).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(put)

	//第二种方式，使用bodyString方式 json

	std2 := `{"name":"george","age":3,"message":"good child","class":"幼儿园"}`

	put2, err := client.Index().
		Index("school").
		Type("class").
		Id("3").
		BodyString(std2).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(put2)

	////#################获取

	//根据ID获取
	get1, err := client.Get().Index("school").Type("class").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(get1.Source))

	//转成student结构体
	b, err := get1.Source.MarshalJSON()
	std3 := &Student{}
	err = json.Unmarshal(b, std3)

	if err != nil {
		panic(err)
	}
	fmt.Println(get1.Found, std3.Age, std3.Message, std3.Name)

	//查询
	termQuery := elastic.NewTermQuery("name", "george")
	searchResult, err := client.Search().Index("school").Query(termQuery).From(0).Size(10).Pretty(true).Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var stdtype Student
	for _, item := range searchResult.Each(reflect.TypeOf(stdtype)) {
		t := item.(Student)
		fmt.Printf("students name %s: message:%s\n", t.Name, t.Message)
	}
	fmt.Printf("Found a total of %d students\n", searchResult.TotalHits())

	////#################更新

	res, err := client.Update().Index("school").Type("class").Id("2").Doc(map[string]interface{}{"name": "shaun"}).Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)

	//###################删除
	resdel, err := client.Delete().Index("school").Type("class").Id("2").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("del data", resdel)
}
