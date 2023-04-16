package usecase

import (
	"context"
	"fmt"
	"log"
	"order-service/service/model/entity"
	"order-service/service/model/request"
	"order-service/service/model/response"
	"order-service/service/repository"
)

type IcartUsercase interface {
	GetCart(ctx context.Context, cartID uint) (cart response.ListProductResponse, err error)
	UpdateCart(ctx context.Context, cartID uint, req request.UpdateCartRequest) (cerr error)
}

type cartUsercase struct {
	cartRepo    repository.IcartRepo
	productRepo repository.IproductRepo
}

func NewCartUsercase(
	cartRepo repository.IcartRepo,
	productRepo repository.IproductRepo,
) IcartUsercase {
	return &cartUsercase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (c *cartUsercase) GetCart(
	ctx context.Context,
	cartID uint,
) (cart response.ListProductResponse, err error) {
	cartInfo, err := c.cartRepo.GetUserCart(ctx, cartID)
	if err != nil {
		log.Printf("GetUserCart: %v", err)
		return
	}
	listProductCart := []uint{}
	mapProductCartDB := make(map[uint]int)
	for _, product := range cartInfo.ProductCart {
		listProductCart = append(listProductCart, product.ProductID)
		mapProductCartDB[product.ProductID] = product.Quantity
	}

	listProductByIDs, err := c.productRepo.GetByIDs(listProductCart)
	for _, product := range listProductByIDs {
		product.Quantity = mapProductCartDB[product.ID]
		cart.Products = append(cart.Products, product)
	}
	if err != nil {
		log.Printf("GetByIDs: %v", err)
		return
	}
	return
}

func (c *cartUsercase) UpdateCart(
	ctx context.Context,
	cartID uint,
	req request.UpdateCartRequest,
) (cerr error) {
	log.Print(cartID)
	cartInfo, err := c.cartRepo.GetUserCart(ctx, cartID)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return
	}

	mapProductInput := make(map[uint]bool)

	for _, product := range req.Products {
		mapProductInput[product.ProductID] = true
	}
	// TODO: shit logic here
	mapProductCartDB := make(map[uint]uint) // hash map [product id] product cart id
	mapProductCartDBFullModel := make(map[uint]entity.ProductCart)
	var listProductDelete []entity.ProductCart
	var updateCartReq []entity.ProductCart
	for _, item := range cartInfo.ProductCart {
		mapProductCartDB[item.ID] = item.ProductID
		mapProductCartDBFullModel[item.ProductID] = item
		// get list product cart delete
		if !mapProductInput[item.ProductID] {
			listProductDelete = append(listProductDelete, entity.ProductCart{
				ID: item.ID,
			})
		}
	}

	for _, product := range req.Products {
		fmt.Println("REQ", req)
		if _, exist := mapProductCartDBFullModel[product.ProductID]; exist {
			updateCartReq = append(updateCartReq, entity.ProductCart{
				ID:        mapProductCartDBFullModel[product.ProductID].ID,
				CartID:    mapProductCartDBFullModel[product.ProductID].CartID,
				Quantity:  mapProductCartDBFullModel[product.ProductID].Quantity,
				ProductID: mapProductCartDBFullModel[product.ProductID].ProductID,
			})
			continue
		}
		updateCartReq = append(updateCartReq, entity.ProductCart{
			CartID:    cartID,
			Quantity:  product.Quantity,
			ProductID: product.ProductID,
		})

	}
	if len(listProductDelete) != 0 {
		err = c.cartRepo.DeleteCartUser(ctx, listProductDelete)
		if err != nil {
			log.Printf("DeleteCartUser: %v", err)
			return
		}
	}

	err = c.cartRepo.UpdateCartUser(ctx, updateCartReq)
	if err != nil {
		log.Printf("UpdateCartUser: %v", err)
		return
	}
	return
}
