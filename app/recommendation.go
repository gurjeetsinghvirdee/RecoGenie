package appwrite

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/appwrite/sdk-for-go/appwrite"
)

// RecommendationENgine represents a recommendation engine
type RecommendationEngine struct {
	appwriteClient *appwrite.Client
	databaseID	 	string
	usersCollectionID string
	productCollectionID string
	ratingCollectionID string
	recommendationCollectionID string
}

// NewRecommendationEngine creates a new recommendation engine
func NewRecommendationEngine(appwriteClient *appwrite.Client, databaseID string, usersCollectionID string, productsCollectionID string, ratingsCollectionID string, recommendationsCollectionID string) *RecommendationEngine {
	return &RecommendationEngine{
		appwriteClient: appwriteClient,
		databaseID: databaseID,
		usersCollectionID: usersCollectionID,
		productCollectionID: productsCollectionID,
		ratingCollectionID: ratingsCollectionID,
		recommendationCollectionID: recommendationsCollectionID,
	}
}

// GetRecommendations returns a list of recommended products
func (re *RecommendationEngine) GetRecommendations(ctx context.Context, userID string) ([]map[string]interface{}, error) {
	// Get the user's ratings
	ratings, err := re.getUserRatings(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Get the similarity matrix
	similarityMatrix, err := re.getSimilarityMatrix(ctx, ratings)
	if err != nil {
		return nil, err
	}

	// Get the recommended products
	recommendedProducts, err := re.getRecommendedProducts(ctx, similarityMatrix, userID)
	if err != nil {
		return nil, err
	}

	return recommendedProducts, nil
}

// getUserRatings retrieves the ratings data for a given user
func (re *RecommendationEngine) getUserRatings(ctx context.Context, userID string) (map[string]float64, error) {
	// Retrieve the user's ratings data from the database
	database := appwrite.NewDatabases(re.appwriteClient)
	document, err := database.GetDocument(ctx, re.ratingCollectionID, userID)
	if err != nil {
		return nil, err
	}

	// Extract the ratings data from the document
	ratings := make(map[string]float64)
	for _, attribute := range document.Attributes {
		ratings[attribute.Key] = attribute.Value.(float64)
	}

	return ratings, nil
}

// getSimilarityMatrix Calculates the similarity matrix for a given user's ratings
func (re *RecommendationEngine) getSimilarityMatrix(ctx context.Context, ratings map[string]float64) (map[string]float64, error) {
	// Retrieve the ratings data for all users from the database
	database := appwrite.NewDatabases(re.appwriteClient)
	users, err := database.ListDocuments(ctx, re.ratingCollectionID)
	if err != nil {
		return nil, err
	}

	// Calculate the similarity b/w the given user's rating and each other user's rating
	similarityMatrix := make(map[string]float64)
	for _, user := range users.Documents {
		similarity := calculateSimilarity(ratings, user.Attributes)
		similarityMatrix[user.ID] = similarity
	}

	return similarityMatrix, nil
}

func calculateSimilarity(ratings1, ratings2 []appwrite.DocumentAttribute) float64 {
	// Calculate the similarity b/w the two sets of ratings
	similarity := 0.0
	for _, rating1 := range ratings1 {
		for _, rating2 := range ratings2 {
			if rating1.Key == rating2.Key {
				similarity++
			}
		}
	}
	similarity /= float64(len(ratings1))

	return similarity
}

// getRecommendedProducts generates product recommendations for a given user based on the similarity matrix
func (re *RecommendationEngine) getRecommendedProducts(ctx context.Context, similarityMatrix map[string]float64, userID string) ([]map[string]interface{}, error) {
	// Retrieve the product data from the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	products, err := database.ListDocuments(ctx, re.productCollectionID)
	if err != nil {
		return nil, err
	}

	// Generate recommendations for the user bases on the similarity matrix
	recommendedProducts := make([]map[string]interface{}, 0)
	for _, product := range products.Documents {
		score := 0.0
		for user, similarity := range similarityMatrix {
			if user != userID {
				userRatings, err := re.getUserRatings(ctx, user)
				if err != nil {
					return nil, err
				}
				if rating, ok := userRatings[product.ID]; ok {
					score += rating * similarity
				}
			}
		}
		recommendedProducts := map[string]interface{}{
			"productID": product.ID,
			"score": score,
		}
		recommendedProducts = append(recommendedProducts, recommendedProduct)
	}

	// Sort the recommended products by score
	sort.Slice(recommendedProducts, func(i, j int) bool {
		score1, ok1 := recommendedProducts[i]["score"].(float64)
		score2, ok2 := recommendedProducts[j]["score"].(float64)
		if ok1 && ok2 {
			return score1 > score2
		}
		return false
	})

	return recommendedProducts, nil
}

// GetSimilarUsers return a list of similar users for a given user
func (re *RecommendationEngine) GetSimilarUsers(ctx context.Context, userID string) ([]map[string]interface{}, error) {
	// Retrieve the user's ratings data
	ratings, err := re.getUserRatings(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Retrieve the ratings data for all users
	database := appwrite.NewDatabases(re.appwriteClient)
	users, err := database.ListDocuments(ctx, re.ratingCollectionID)
	if err != nil {
		return nil, err
	}

	// Calculate the similarity b/w the given user and each other user
	similarUsers := make([]map[string]interface{}, 0)
	for _, user := range users.Documents {
		if user.ID != userID {
			similarity := calculateSimilarity(ratings, user.Attributes)
			similarUser := map[string]interface{}{
				"userID": user.ID,
				"similarity": similarity,
			}
			similarUsers = append(similarUsers, similarUser)
		}
	}

	// Sort the similar users by similarity
	sort.Slice(similarUsers, func(i, j int) bool {
		similarity1, ok1 := similarUsers[i]["similarity"].(float64)
		similarity2, ok2 := similarUsers[j]["similarity"].(float64)
		if ok1 && ok2 {
			return similarity1 > similarity2
		}
		return false
	})

	return similarUsers, nil
}

// UpdateUserRatings updates the ratings data for a given user
func (re *RecommendationEngine) UpdateUserRatings(ctx context.Context, userID string, ratings map[string]float64) error {
	// Update the user's ratings data in the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	_, err := database.UpdateDocument(ctx, re.ratingCollectionID, userID, ratings)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserRatings deletes the ratings data for a given user
func (re *RecommendationEngine) DeleteUserRatings(ctx context.Context, userID string) error {
	// Delete the user's ratings data from the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	_, err := database.DeleteDocument(ctx, re.ratingCollectionID, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetProductRatings returns the ratings data for a given product
func (re *RecommendationEngine) GetProductRatings(ctx context.Context, productID string) ([]map[string]interface{}, error) {
	// Retrieve the ratings data for the product from the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	ratings, err := database.ListDocuments(ctx, re.ratingCollectionID, appwrite.QueryEqual("productID", productID))
	if err != nil {
		return nil, err
	}

	// Extract the ratings data from the documents
	productRatings := make([]map[string]interface{}, 0)
	for _, rating := range ratings.Documents {
		productRating := map[string]interface{}{
			"userID": rating.ID,
			"rating": rating.Attributes["rating"].(float64),
			"productID": productID,
		}
		productRatings = append(productRatings, productRating)
	}

	return productRatings, nil
}

// UpdateProductRatings updates the ratings data for a given product
func (re *RecommendationEngine) UpdateProductRatings(ctx context.Context, productID string, ratings []map[string]interface{}) error {
	// Update the ratings data for the product in the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	for _, rating := range ratings {
		_, err := database.UpdateDocument(ctx, re.ratingCollectionID, rating["userID"].(string), map[string]interface{}{
			"rating": rating["rating"].(float64),
			"productID": productID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteProductRatings deletes the ratings data for a given product
func (re *RecommendationEngine) DeleteProductRatings(ctx context.Context, productID string) error {
	// Delete the ratings data for the product from the DB
	database := appwrite.NewDatabases(re.appwriteClient)
	_, err := database.DeleteDocuments(ctx, re.ratingCollectionID, appwrite.QueryEqual("productID", productID))
	if err != nil {
		return err
	}

	return nil
}

// Train trains the recommendation engine using the provided ratings data
func (re *RecommendationEngine) Train(ctx context.Context, ratings []map[string]interface{}) error {
	// This can be implemented using a machine learning algorithm like collaborative filtering
	// For now, just return nil
	return nil
}

func recommend() {
	// Load the environment variables
	endpoint := os.Getenv("APPWRITE_ENDPOINT")
	databaseID := os.Getenv("APPWRITE_DATABASE_ID")
	usersCollectionID := os.Getenv("USERS_COLLECTION_ID")
	productsCollectionID := os.Getenv("PRODUCTS_COLLECTION_ID")
	ratingsCollectionID := os.Getenv("RATINGS_COLLECTION_ID")
	recommendationsCollectionID := os.Getenv("RECOMMENDATIONS_COLLECTION_ID")

	// Initialize the Appwrite client
	appwriteClient := appwrite.NewClient(context.Background(), endpoint)

	// Initialize the recommendation engine
	RecommendationEngine := NewRecommendationEngine(appwriteClient, databaseID, usersCollectionID, productsCollectionID, ratingsCollectionID, recommendationsCollectionID)

	// Get recommendations for a user
	userID := "userID"
	recommendations, err := recommendationEngine.GetRecommendations(context.Background(), userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recommendations)
}