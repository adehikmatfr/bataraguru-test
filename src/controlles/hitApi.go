package controlles

import (
	"batara/src/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//type Hit struct {
//	dateTime string `bson:"dateTime"`
//	status   string `bson:"status"`
//}

func HitApi(c *gin.Context) {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	tm:=time.Now();
	dateT:=tm.Format("2006-01-02 15:04:05")
	ctx:= context.Background()
	_, err = db.Collection("hit").InsertOne(ctx, gin.H{"date_time":dateT,"status":"limit"})
	log.Println("limit Request Inserted")
	c.String(http.StatusOK, fmt.Sprintf("Hallo Word limit Request"))
}
