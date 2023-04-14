package mock

import (
	"encoding/json"
	"log"
	"order-service/service/model/entity"
	"sync"
	"time"

	"gorm.io/gorm"
)

func MakeMockData(db *gorm.DB) (err error) {

	var cates []entity.Category
	err = json.Unmarshal([]byte(DataCate), &cates)
	if err != nil {
		log.Fatal("Can't unmarshal data mock error", err)
		return
	}
	var count int64
	err = db.Model(&entity.Product{}).Count(&count).Error
	if err != nil {
		log.Fatal("Can't count data from db error", err)
	}
	if count == 0 {
		InsertCateMock(cates, db)
	}
	var dataOut []entity.Product
	err = json.Unmarshal([]byte(DataMock), &dataOut)
	if err != nil {
		log.Fatal("Can't unmarshal data mock error", err)
		return

	}

	err = db.Model(&entity.Product{}).Count(&count).Error
	if err != nil {
		log.Fatal("Can't count data from db error", err)
	}
	if count == 0 {
		for i := 0; i < 10; i++ {
			InsertDatMock(dataOut, db)
			time.Sleep(5 * time.Second)
		}
	}

	return
}

func InsertCateMock(cates []entity.Category, db *gorm.DB) {
	var count int64
	err := db.Model(&entity.Category{}).Count(&count).Error
	if err != nil {
		log.Fatal("Can't count data from db error", err)
	}
	if count != 0 {
		return
	}

	db.Session(&gorm.Session{CreateBatchSize: 100000})

	err = db.Create(&cates).Error
	if err != nil {
		log.Printf("Error creating %v", err)
		return
	}
}

func InsertDatMock(products []entity.Product, db *gorm.DB) error {

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			newListProduct := []entity.Product{}

			for j, product := range products {
				product.UniqueOffset = int(time.Now().Unix()) + i*10000 + j
				newListProduct = append(newListProduct, product)
			}
			defer wg.Done()
			err := db.Table("products").Create(&newListProduct).Error
			if err != nil {
				log.Printf("Error creating %v", err)
				return
			}
		}()
	}
	return nil
}
