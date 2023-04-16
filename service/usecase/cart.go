package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"order-service/service/model/entity"
	"order-service/service/model/request"
	"order-service/service/model/response"
	"order-service/service/repository"
)

type IcartUsercase interface {
	GetCart(ctx context.Context, cartID uint) (cart response.ListProductResponse, err error)
	UpdateCart(ctx context.Context, cartID uint, req request.UpdateCartRequest) (cerr error)
	CreateSalesOrder(ctx context.Context, cartID uint, userID uint) error
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

func (c *cartUsercase) CreateSalesOrder(ctx context.Context, cartID uint, userID uint) error {
	cartInfo, err := c.GetCart(ctx, cartID)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return err
	}

	var productOrder []entity.ProductOrder
	var totalPrice int
	for _, product := range cartInfo.Products {
		productOrder = append(productOrder, entity.ProductOrder{
			ID:        cartID,
			ProductID: product.ID,
			Quantity:  product.Quantity,
			Price:     product.Price,
		})
		totalPrice += int(product.Price) * product.Quantity
	}

	var request = entity.Order{
		ProductOrder: productOrder,
		UserID:       userID,
		TotalPrice:   totalPrice,
	}

	err = c.cartRepo.CreateSaleOrder(ctx, request)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return err
	}

	out, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	c.CallGoogleChat(string(out))

	return nil
}

func (c *cartUsercase) CallGoogleChat(text string) {
	var ThreadName = "spaces/AAAAtTHk3b8/threads/I1ghIyG8pDo"
	var URL = "https://chat.googleapis.com/v1/spaces/AAAAtTHk3b8/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=p82eX5PmSpaRfjmd_rajXI2ZBBIVonlV8qX9X5hPeqE%3D"

	type ggResponse struct {
		Name string `json:"name"`
		Text string `json:"text"`
	}
	type thread struct {
		Name string `json:"name"`
	}
	messageData := struct {
		Text   string `json:"text"`
		Thread thread `json:"thread"`
	}{
		Text: text,
		Thread: thread{
			Name: ThreadName,
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(messageData)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.Post(URL, "application/json", &buf)
	if err != nil {
		panic(err)
	}
}
