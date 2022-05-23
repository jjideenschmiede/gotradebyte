//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gotradebyte.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gotradebyte

// OrdersBody is to decode xml data
type OrdersBody struct {
	Order []OrdersBodyOrder `xml:"ORDER"`
}

type OrdersBodyOrder struct {
	OrderData OrdersBodyOrderData `xml:"ORDER_DATA"`
	SellTo    OrdersBodySellTo    `xml:"SELL_TO"`
	ShipTo    OrdersBodyShipTo    `xml:"SHIP_TO"`
	Shipment  OrdersBodyShipment  `xml:"SHIPMENT"`
	Payment   OrdersBodyPayment   `xml:"PAYMENT"`
	Items     OrdersBodyItems     `xml:"ITEMS"`
}

type OrdersBodyOrderData struct {
	OrderDate       string `xml:"ORDER_DATE"`
	TbId            int    `xml:"TB_ID"`
	ChannelSign     string `xml:"CHANNEL_SIGN"`
	ChannelId       string `xml:"CHANNEL_ID"`
	ChannelNo       string `xml:"CHANNEL_NO"`
	Approved        int    `xml:"APPROVED"`
	ItemCount       int    `xml:"ITEM_COUNT"`
	TotalItemAmount string `xml:"TOTAL_ITEM_AMOUNT"`
}

type OrdersBodySellTo struct {
	TbId            int    `xml:"TB_ID"`
	ChannelNo       int    `xml:"CHANNEL_NO"`
	FirstName       string `xml:"FIRSTNAME"`
	LastName        string `xml:"LASTNAME"`
	StreetNo        string `xml:"STREET_NO"`
	StreetExtension string `xml:"STREET_EXTENSION"`
	Zip             string `xml:"ZIP"`
	City            string `xml:"CITY"`
	Country         string `xml:"COUNTRY"`
	PhonePrivate    string `xml:"PHONE_PRIVATE"`
	Email           string `xml:"EMAIL"`
}

type OrdersBodyShipTo struct {
	TbId            int    `xml:"TB_ID"`
	ChannelNo       int    `xml:"CHANNEL_NO"`
	FirstName       string `xml:"FIRSTNAME"`
	LastName        string `xml:"LASTNAME"`
	StreetNo        string `xml:"STREET_NO"`
	StreetExtension string `xml:"STREET_EXTENSION"`
	Zip             string `xml:"ZIP"`
	City            string `xml:"CITY"`
	Country         string `xml:"COUNTRY"`
}

type OrdersBodyShipment struct {
	Price float64 `xml:"PRICE"`
}

type OrdersBodyPayment struct {
	Price       float64     `xml:"PRICE"`
	DirectDebit interface{} `xml:"DIRECTDEBIT"`
}

type OrdersBodyItems struct {
	Item []OrdersBodyItem `xml:"ITEM"`
}

type OrdersBodyItem struct {
	TbId          int     `xml:"TB_ID"`
	ChannelId     string  `xml:"CHANNEL_ID"`
	Sku           int     `xml:"SKU"`
	ChannelSku    string  `xml:"CHANNEL_SKU"`
	Quantity      int     `xml:"QUANTITY"`
	BillingText   string  `xml:"BILLING_TEXT"`
	TransferPrice float64 `xml:"TRANSFER_PRICE"`
	ItemPrice     float64 `xml:"ITEM_PRICE"`
	DateCreated   string  `xml:"DATE_CREATED"`
}
