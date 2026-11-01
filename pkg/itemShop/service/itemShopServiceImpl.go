package service

import (
	_itemShopRepository "github.com/Inwza/Inwza-shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	ItemShopRepostitory _itemShopRepository.ItemShopRepostitory
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepostitory,
) ItemShopService {
	return &itemShopServiceImpl{ItemShopRepostitory: itemShopRepository}
}
