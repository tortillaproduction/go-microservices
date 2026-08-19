package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/models"
)

type categoryRepositorySQLBoiler struct{}

func NewCategoryRepositorySQLBoiler() categories.CategoryRepository {
	models.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, CategoryAfterDeleteHook)
	return &categoryRepositorySQLBoiler{}
}

// methods
func (rep *categoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	condition := models.CategoryWhere.Name.EQ(category.Name().Value())

	if exists, err := models.Categories(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%s already exists", category.Name().Value()))
	}
}

func (rep *categoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	new_category := models.Category{
		InternalID: 0,
		ID:         category.Id().Value(),
		Name:       category.Name().Value(),
	}

	if err := new_category.Insert(ctx, tran, boil.Whitelist("id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (rep *categoryRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	up_model, err := models.Categories(qm.Where("id = ?", category.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Category ID: %s does not exist and could not be updated", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}

	up_model.ID = category.Id().Value()
	up_model.Name = category.Name().Value()

	if _, err = up_model.Update(ctx, tran, boil.Whitelist("id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (rep *categoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	del_model, err := models.Categories(qm.Where("id = ?", category.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Category ID: %s does not exist and could not be deleted", category.Id().Value()))
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
func CategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("category ID: %s category Name: %s created.\n", category.ID, category.Name)
	return nil
}

func CategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("category ID: %s category Name: %s updated.\n", category.ID, category.Name)
	return nil
}

func CategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("category ID: %s category Name: %s deleted.\n", category.ID, category.Name)
	return nil
}
