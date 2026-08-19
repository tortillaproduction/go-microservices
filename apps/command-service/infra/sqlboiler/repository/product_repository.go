package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/models"
)

type productRepositorySQLBoiler struct{}

func NewProductRepositorySQLBoiler() products.ProductRepository {
	models.AddProductHook(boil.AfterInsertHook, ProductAfterInsertHook)
	models.AddProductHook(boil.AfterUpdateHook, ProductAfterUpdateHook)
	models.AddProductHook(boil.AfterDeleteHook, ProductAfterDeleteHook)
	return &productRepositorySQLBoiler{}
}

// methods
func (rep *productRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	condition := models.ProductWhere.Name.EQ(product.Name().Value())
	exists, err := models.Products(condition).Exists(ctx, tran)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%s already exists", product.Name().Value()))
	}
}

func (rep *productRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	new_product := models.Product{
		InternalID: 0,
		ID:         product.Id().Value(),
		Name:       product.Name().Value(),
		Price:      int(product.Price().Value()),
		CategoryID: product.Category().Id().Value(),
	}

	if err := new_product.Insert(ctx, tran, boil.Whitelist("id", "name", "price", "category_id")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (rep *productRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	up_model, err := models.Products(qm.Where("id = ?", product.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Product ID: %s does not exist and could not be updated", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}

	up_model.Name = product.Name().Value()
	up_model.Price = int(product.Price().Value())

	if _, err = up_model.Update(ctx, tran, boil.Whitelist("id", "name", "price")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (rep *productRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	del_model, err := models.Products(qm.Where("id = ?", product.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Product ID: %s does not exist and could not be deleted", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}

	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

// hooks
func ProductAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("product ID: %s product Name: %s price: %d category number: %s created.\n", product.ID, product.Name, product.Price, product.CategoryID)
	return nil
}

func ProductAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("product ID: %s product Name: %s price: %d category number: %s updated.\n", product.ID, product.Name, product.Price, product.CategoryID)
	return nil
}

func ProductAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("product ID: %s product Name: %s price: %d category number: %s deleted.\n", product.ID, product.Name, product.Price, product.CategoryID)
	return nil
}
