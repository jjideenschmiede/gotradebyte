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
	Items     []OrdersBodyItem    `xml:"ITEMS"`
}

type OrdersBodyOrderData struct {
	OrderDate       string `xml:"ORDER_DATE"`
	ChannelSign     string `xml:"CHANNEL_SIGN"`
	ChannelId       string `xml:"CHANNEL_ID"`
	ChannelNo       string `xml:"CHANNEL_NO"`
	Approved        int    `xml:"APPROVED"`
	ItemCount       int    `xml:"ITEM_COUNT"`
	TotalItemAmount string `xml:"TOTAL_ITEM_AMOUNT"`
}

type OrdersBodySellTo struct {
	ChannelNo    int    `xml:"CHANNEL_NO"`
	FirstName    string `xml:"FIRSTNAME"`
	LastName     string `xml:"LASTNAME"`
	StreetNo     string `xml:"STREET_NO"`
	Zip          string `xml:"ZIP"`
	City         string `xml:"CITY"`
	Country      string `xml:"COUNTRY"`
	PhonePrivate string `xml:"PHONE_PRIVATE"`
	Email        string `xml:"EMAIL"`
}

type OrdersBodyShipTo struct {
	ChannelNo int    `xml:"CHANNEL_NO"`
	FirstName string `xml:"FIRSTNAME"`
	LastName  string `xml:"LASTNAME"`
	StreetNo  string `xml:"STREET_NO"`
	Zip       string `xml:"ZIP"`
	City      string `xml:"CITY"`
	Country   string `xml:"COUNTRY"`
}

type OrdersBodyItem struct {
	ChannelId   int     `xml:"CHANNEL_ID"`
	Sku         string  `xml:"SKU"`
	ChannelSku  string  `xml:"CHANNEL_SKU"`
	Ean         string  `xml:"EAN"`
	Quantity    int     `xml:"QUANTITY"`
	BillingText string  `xml:"BILLING_TEXT"`
	ItemPrice   float64 `xml:"ITEM_PRICE"`
}
