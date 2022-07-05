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

// MESSAGES_LIST is to encode xml data
type MESSAGES_LIST struct {
	Message []DeliveryMessagesBodyMessage `xml:"MESSAGE"`
}

type DeliveryMessagesBodyMessage struct {
	MessageType       string `xml:"MESSAGE_TYPE"`
	TbOrderId         int    `xml:"TB_ORDER_ID"`
	TbOrderItemId     int    `xml:"TB_ORDER_ITEM_ID"`
	Sku               string `xml:"SKU"`
	ChannelSign       string `xml:"CHANNEL_SIGN"`
	ChannelOrderId    string `xml:"CHANNEL_ORDER_ID"`
	Quantity          int    `xml:"QUANTITY"`
	CarrierParcelType string `xml:"CARRIER_PARCEL_TYPE"`
	IdCode            string `xml:"IDCODE"`
}
