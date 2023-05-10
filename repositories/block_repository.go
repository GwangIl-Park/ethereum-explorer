package repositories

import (
	"context"
	"ethereum-explorer/models"

	"ethereum-explorer/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blockRepository struct {
	db *db.DB
}

func NewBlockRepository(db *db.DB) models.BlockRepository {
	return &blockRepository {
		db,
	}
}

func (br *blockRepository) GetBlocks(c context.Context, page int64, show int64) ([]models.Block, error) {
	opts := options.Find().SetSort(bson.D{{Key: "blockheight",Value: -1}}).SetSkip((page-1) * show).SetLimit(show)
	cursor, err := br.db.Collections["blocks"].Find(c, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var blocks []models.Block
	if err := cursor.All(c, &blocks); err != nil {
		return nil, err
	}

	return blocks, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height string) (models.Block, error) {
	cursor := br.db.Collections["blocks"].FindOne(c, bson.M{"blockheight":height})
	
	var block models.Block
	if err := cursor.Decode(&block); err != nil {
		return models.Block{}, err
	}
	
	return block, nil
}
