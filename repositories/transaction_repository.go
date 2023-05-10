package repositories

import (
	"context"
	"ethereum-explorer/db"
	"ethereum-explorer/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type transactionRepository struct {
	db *db.DB
}

func NewTransactionRepository(db *db.DB) models.TransactionRepository {
	return &transactionRepository {
		db,
	}
}

func (tr *transactionRepository) GetTransactions(c context.Context, page int64, show int64) ([]models.Transaction, error) {
	opts := options.Find().SetSort(bson.D{{Key: "blockheight",Value: -1}}).SetSkip((page-1) * show).SetLimit(show)
	cursor, err := tr.db.Collections["transactions"].Find(c, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	if err := cursor.All(c, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tr *transactionRepository) GetTransactionByHash(c context.Context, hashParam string) (models.Transaction, error) {
	cursor := tr.db.Collections["transactions"].FindOne(c, bson.M{"hash":hashParam})
	
	var transaction models.Transaction
	if err := cursor.Decode(&transaction); err != nil {
		return models.Transaction{}, err
	}
	
	return transaction, nil
}

func (tr *transactionRepository) GetTransactionsByAccount(c context.Context, account string) ([]models.Transaction, error) {
	cursor, err := tr.db.Collections["transactions"].Find(c, bson.M{"account":account})
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	if err := cursor.All(c, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}