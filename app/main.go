package appwrite

import (
    "context"
    "log"
    "os"

	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
    // Load environment variables
    endpoint := os.Getenv("APPWRITE_ENDPOINT")
    apiKey := os.Getenv("APPWRITE_API_KEY")
    projectID := os.Getenv("APPWRITE_PROJECT_ID")
    databaseID := os.Getenv("APPWRITE_DATABASE_ID")
    productsCollectionID := os.Getenv("PRODUCTS_COLLECTION_ID")


    // Check if environment variables are set
    if endpoint == "" || apiKey == "" || projectID == "" || databaseID == "" || productsCollectionID == "" {
        log.Fatal("Environment variables are not set")
    }

    // Create a new Appwrite client
    appwriteClient, err := appwrite.NewClient(endpoint, projectID, apiKey)
    if err != nil {
        log.Fatal(err)
    }

    // Create a new document
    doc, err := appwriteClient.CreateProduct(context.Background(), databaseID, productsCollectionID, map[string]interface{}{
        "title":       "Product 1",
        "description": "Description of Product 1",
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Document created: %+v\n", doc)

    // Get a list of documents
    docs, err := appwriteClient.ListProducts(context.Background(), databaseID, productsCollectionID)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Documents: %+v\n", docs)

    // Update a document
    updateDoc, err := appwriteClient.UpdateProduct(context.Background(), databaseID, productsCollectionID, doc.ID, map[string]interface{}{
        "title": "Product 1 Updated",
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Updated document: %+v\n", updateDoc)

    // Delete a document
    err = appwriteClient.DeleteProduct(context.Background(), databaseID, productsCollectionID, doc.ID)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Document deleted successfully")
}
