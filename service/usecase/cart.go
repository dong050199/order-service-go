package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"order-service/pkg/infra"
	"order-service/service/model/entity"
	"order-service/service/model/request"
	"order-service/service/model/response"
	"order-service/service/repository"

	"gorm.io/gorm"
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
		cart.Products = append(cart.Products, response.ProductResponse{
			Product:    product,
			TotalPrice: product.Quantity * int(product.Price),
		})
	}
	if err != nil {
		log.Printf("GetByIDs: %v", err)
		return
	}

	for _, product := range cart.Products {
		cart.TotalPrice += int(product.Price) * product.Quantity
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

	for i, product := range req.Products {
		fmt.Println("REQ", req)
		if _, exist := mapProductCartDBFullModel[product.ProductID]; exist {
			updateCartReq = append(updateCartReq, entity.ProductCart{
				ID:        mapProductCartDBFullModel[product.ProductID].ID,
				CartID:    mapProductCartDBFullModel[product.ProductID].CartID,
				Quantity:  req.Products[i].Quantity,
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

type CreateSaleOrderRequest struct {
	CartID         uint `json:"cart_id"`
	UserID         uint `json:"user_id"`
	PaymentDetails struct {
		TypeOfPayment  string `json:"type_of_payment"` // online or COD
		ShippingMethod string `json:"shipping_method"` // method like ... shipping
	} `json:"payment_details"`
}

func (c *cartUsercase) CreateSalesOrder(ctx context.Context, cartID uint, userID uint) error {
	cartInfo, err := c.GetCart(ctx, cartID)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return err
	}
	tx, err := infra.BeginTransaction()
	if err != nil {
		log.Printf("BeginTransaction: %v", err)
		return err
	}

	var productOrder []entity.ProductOrder
	var totalPrice int
	mapProductQuantity := make(map[uint]int)
	var productIDs []uint
	for _, product := range cartInfo.Products {
		productOrder = append(productOrder, entity.ProductOrder{
			ProductID: product.ID,
			Quantity:  product.Quantity,
			Price:     product.Price,
		})
		totalPrice += int(product.Price) * product.Quantity
		productIDs = append(productIDs, product.ID)
		mapProductQuantity[product.ID] = product.Quantity
	}

	products, err := c.productRepo.GetByIDs(productIDs)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return err
	}

	var productsUpdate []entity.Product

	for _, product := range products {
		if product.Quantity < mapProductQuantity[product.ID] {
			return errors.New("Number of products is out of range.")
		}

		product.Quantity = product.Quantity - mapProductQuantity[product.ID]
		productsUpdate = append(productsUpdate, product)
	}

	var saleOrder = entity.Order{
		ProductOrder: productOrder,
		UserID:       userID,
		TotalPrice:   totalPrice,
	}

	err = c.productRepo.Update(productsUpdate, tx)
	if err != nil {
		log.Printf("Update: %v", err)
		return err
	}
	infra.ReleaseTransaction(tx, err)

	err = c.cartRepo.CreateSaleOrder(ctx, saleOrder)
	if err != nil {
		log.Printf("GetCart: %v", err)
		return err
	}
	// remove all from cart
	err = c.UpdateCart(ctx, cartID, request.UpdateCartRequest{
		Products: []request.ProductCart{},
	})
	if err != nil {
		panic(err)
	}

	out, _ := json.Marshal(saleOrder)
	// send message to google chat
	go c.CallGoogleChat(string(out))

	return nil
}

func (c *cartUsercase) UpdateStock(tx *gorm.DB) error {

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
