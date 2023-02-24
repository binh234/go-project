package controller

import (
	"context"
	"graphql/database"
	"graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg *database.MongoInstance = database.GetMongoInstance()
var collection = mg.DB.Collection("jobs")

func GetAllJobs() []*model.JobListing {
	query := bson.D{{}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	jobs := make([]*model.JobListing, 0)
	err = cursor.All(context.TODO(), &jobs)
	if err != nil {
		log.Fatal(err)
	}
	return jobs
}

func GetJobListing(id string) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.M{"_id": jobId}

	var jobListing model.JobListing
	err = collection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inserted, err := collection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"company":     jobInfo.Company,
		"url":         jobInfo.URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := inserted.InsertedID.(primitive.ObjectID).Hex()
	jobListing := model.JobListing{
		ID:          insertedID,
		Title:       jobInfo.Title,
		Description: jobInfo.Description,
		Company:     jobInfo.Company,
		URL:         jobInfo.URL,
	}
	return &jobListing
}

func UpdateJobListing(id string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	updateJobInfo := bson.M{}
	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}
	filter := bson.M{"_id": jobId}
	update := bson.M{"$set": updateJobInfo}
	result := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing
	if err = result.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}
	return &jobListing

}

func DeleteJobListing(id string) *model.DeleteJobResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": jobId}
	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteJobResponse{DeleteJobID: id}
}
