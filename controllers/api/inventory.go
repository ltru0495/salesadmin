package api

import (
	"admin/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Code     string  `json:"code" bson:"code"`
	Brand    string  `json:"brand" json:"brand"`
	PFC      string  `json:"pfc" bson:"pfc"` //Factory Product Code
	Size     int     `json:"size" bson:"size"`
	Model    string  `json:"model" bson:"model"`
	Price    float64 `json:"price" bson:"price"`
	Location string  `json:"location" bson:"location"`
	Note     string  `json:"note" bson:"note"`
}

type Response struct {
	PageCount int64     `json:"pageCount"`
	Products  []Product `json:"products"`
}

func pagOpts(limit, page int64) *options.FindOptions {
	var skip int64
	skip = limit * (page - 1)
	return &options.FindOptions{Limit: &limit, Skip: &skip}
}

func GetDB() *mongo.Database {
	uri := "mongodb://127.0.0.1:27017/"
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client.Database("salesadmin")
}

func GetCollection(db *mongo.Database, c string) *mongo.Collection {
	return db.Collection(c)
}

func PrintProducts(prods []Product) {
	for _, p := range prods {
		log.Println(p.Code)
	}
}

func getFilter(r *http.Request) bson.M {
	filter := bson.M{}
	// fields := []string{"code", "brand", "pfc", "size", "model", "price"}
	fields := []string{"code", "brand", "pfc", "size", "model", "price", "location", "note"}
	for _, f := range fields {
		regex := r.URL.Query().Get(f)
		if f == "size" || f == "price" {
			s, _ := strconv.Atoi(regex)
			if s != 0 {
				filter[f] = s
			}
			continue
		}
		if regex != "" {
			filter[f] = bson.M{
				"$regex": regex,
			}
		}
	}
	return filter
}

func pages(count, limit int64) int64 {
	p := count / limit
	if count%limit == 0 {
		return p
	}
	return p + 1
}

func Inventory(w http.ResponseWriter, r *http.Request) {
	var DB *mongo.Database = GetDB()

	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	collection := GetCollection(DB, "products")

	filter := getFilter(r)
	// log.Println(filter)

	limit := int64(10)
	opts := pagOpts(limit, int64(page))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	// f := bson.D{{"size", 25}}

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	result := []Product{}
	for cur.Next(ctx) {
		var p Product
		if err := cur.Decode(&p); err != nil {
			log.Fatal(err)
		}
		if utils.GetUser(r).Role != "admin" {
			p.Price = 0.0
			p.Code = ""
		}
		result = append(result, p)
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(&Response{
		PageCount: pages(count, limit),
		Products:  result,
	})
	if err != nil {
		log.Fatal(err)
	}

	w.Write(j)
}
