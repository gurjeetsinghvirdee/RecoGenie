package main

import (
	"context"
	"log"
	"os"


	"github.com/appwrite/sdk-for-go/client"
)

func main() {
	// Load environment variables
	endpoint := os.Getenv("APPWRITE_ENDPOINT")
	apiKey := os.Getenv("APPWRITE_API_KEY")
	projectID := os.Getenv("APPWRITE_PROJECT_ID")
	databaseID := os.Getenv("APPWRITE_DATABASE_ID")
	usersCollectionID := os.Getenv("USERS_COLLECTION_ID")
	productsCollectionID := os.Getenv("PRODUCTS_COLLECTION_ID")
	ratingsCollectionID := os.Getenv("RATINGS_COLLECTION_ID")
	recommendationsCollectionID := os.Getenv("RECOMMENDATIONS_COLLECTION_ID")

	// Check if environment variables are set
	if endpoint == "" || apiKey == "" || projectID == "" || databaseID == "" || usersCollectionID == "" || productsCollectionID == "" || ratingsCollectionID == "" || recommendationsCollectionID == "" {
		log.Fatal("Environment variables are not set")
	}

	// Create a new document
	doc, err := client.CreateDocument(context.Background(), databaseID, productsCollectionID, map[string]interface{}{
		"title": "Product 1",
		"description": "Description of Product 1",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Document created: %+v\n", doc)

	// Get a list of documents
	docs, err := client.ListDocuments(context.Background(), databaseId, productsCollectionID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Documents: %+v\n", docs)

	// Update a document
	updateDoc, err := client.UpdateDocument(context.Background(), databaseID, productsCollectionID, doc.ID, map[string]interface{}{
		"title": "Product 1 Updated",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated document: %+v\n", updateDoc)

	// Delete a document
	err = client.DeleteDocument(context.Background(), databaseID, productsCollectionID, doc.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Document deleted successfully")
}