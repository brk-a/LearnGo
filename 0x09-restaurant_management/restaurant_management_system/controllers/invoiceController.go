package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceViewFormat struct {
    Invoice_id string
    Order_id string
    Table_id interface{}
    Order_details interface{}
    Payment_status *string
    Payment_method string
    Payment_due interface{}
    Payment_due_date time.Time
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

        result, err := invoiceCollection.Find(ctx, bson.M{})
        defer cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching invoices"})
            return
        }

        var allInvoices []bson.M
        if err = result.All(ctx, &allInvoices); err!= nil {
            log.Fatal(err)
        }

        defer cancel()
        c.JSON(http.StatusOK, allInvoices)
    }
}

func GetInvoice() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        invoiceId := c.Param("invoice_id")
        var invoice models.Invoice

        err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "error fetching invoice"})
            return
        }

        var invoiceView InvoiceViewFormat
        allOrderItems, err := ItemsByOrder(invoice.Order_id)
        defer cancel()
        if err!= nil {
            log.Fatal(err)
        }
        invoiceView.Order_id = invoice.Order_id
        invoiceView.Payment_due_date = invoice.Payment_due_date
        invoiceView.Payment_method = "null"
        if invoice.Payment_method!=nil {
            invoiceView.Payment_method = *invoice.Payment_method
        }
        invoiceView.Payment_status = invoice.Payment_status
        invoiceView.Invoice_id = *&invoice.Invoice_id
        invoiceView.Payment_due = allOrderItems[0]["payment_due"]
        invoiceView.Table_id = allOrderItems[0]["table_id"]
        invoiceView.Order_details = allOrderItems[0]["order_details"]

        defer cancel()
        c.JSON(http.StatusOK, invoiceView)
    }
}

func CreateInvoice() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var invoice models.Invoice

        defer cancel()
        if err := c.BindJSON(&invoice); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        validationError := validate.Struct(invoice)
        defer cancel()
        if validationError!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
            return
        }

        var order models.Order
        err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
        defer cancel()
        if err!=nil {
            msg := fmt.Sprintf("error finding order associated with this invoice")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        status := "PENDING"
        if invoice.Payment_status==nil {
            invoice.Payment_status = &status
        }
        invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0 ,0 ,1).Format(time.RFC3339))
        invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        invoice.ID = primitive.NewObjectID()
        invoice.Invoice_id = invoice.ID.Hex()

        result, err := invoiceCollection.InsertOne(ctx, invoice)
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error creating invoice")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"invoice": result})

    }
}

func UpdateInvoice() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var invoice models.Invoice

        defer cancel()
        if err := c.BindJSON(&invoice); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        invoiceId := c.Param("invoice_id")
        filter := bson.M{"invoice_id": invoiceId}
        var updateObj primitive.D
        if invoice.Payment_method!=nil {
            updateObj = append(updateObj, bson.E{"payment_method", invoice.Payment_method})
        }
        if invoice.Payment_status!=nil {
            updateObj = append(updateObj, bson.E{"payment_status", invoice.Payment_status})
        }
        status := "PENDING"
        if invoice.Payment_status==nil {
            invoice.Payment_status = &status
        }

        invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", invoice.Updated_at})

        upsert := true
        opt := options.UpdateOptions{
            Upsert: &upsert,
        }
        result, err := orderCollection.UpdateOne(
            ctx,
            filter,
            bson.M{"$set": invoice},
            &opt,
        )
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error updating invoice")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
    }
}

func DeleteInvoice() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        invoiceId := c.Param("invoice_id")
        filter := bson.M{"invoice_id": invoiceId}

        result, err := invoiceCollection.DeleteOne(ctx, filter)
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error deleting invoice")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        if result.DeletedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"message": "invoice deleted successfully"})
    }
}